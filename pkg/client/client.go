package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/riccardotornesello/irapi-go/pkg/logger"
)

type ApiClient struct {
	accessToken        string
	refreshToken       string
	accessTokenExpiry  time.Time
	refreshTokenExpiry time.Time

	// clientId is required for token refresh
	clientId string
	// clientSecret is required for token refresh
	clientSecret string

	client     *http.Client
	retryAfter time.Time

	concurrency int
	semaphore   chan struct{}
	mutex       sync.Mutex

	// autoRefreshDisabled indicates whether automatic token refresh is disabled
	autoRefreshDisabled bool
	// tokenMutex is used to synchronize token refresh operations
	tokenMutex sync.Mutex
}

type Options struct {
	Concurrency int

	ClientId     string
	ClientSecret string

	AutoRefreshDisabled bool
}

type iRacingResponse struct {
	Link string `json:"link"`
}

func createApiClientWithToken(accessToken, refreshToken string, options *Options) *ApiClient {
	// Use default options if none provided
	if options == nil {
		options = &Options{}
	}

	// Create the semaphore channel if concurrency limit is set
	var semaphore chan struct{}
	if options.Concurrency > 0 {
		semaphore = make(chan struct{}, options.Concurrency)
	}

	// Extract expiry times from JWT tokens.
	accessTokenExpiry, _ := getExpiryFromJwt(accessToken)
	refreshTokenExpiry, _ := getExpiryFromJwt(refreshToken)

	return &ApiClient{
		accessToken:        accessToken,
		refreshToken:       refreshToken,
		accessTokenExpiry:  accessTokenExpiry,
		refreshTokenExpiry: refreshTokenExpiry,

		clientId:     options.ClientId,
		clientSecret: options.ClientSecret,

		client: &http.Client{
			Timeout: 60 * time.Second, // TODO: allow customization
		},

		concurrency: options.Concurrency,
		semaphore:   semaphore,
		mutex:       sync.Mutex{},

		autoRefreshDisabled: options.AutoRefreshDisabled,
		tokenMutex:          sync.Mutex{},
	}
}

func NewPasswordLimitedApiClient(username, password string, options *Options) (*ApiClient, error) {
	if options == nil {
		options = &Options{}
	}
	if options.ClientId == "" || options.ClientSecret == "" {
		return nil, fmt.Errorf("clientId and clientSecret must be provided in options")
	}

	tokenResponse, err := getPasswordLimitedAccessToken(options.ClientId, options.ClientSecret, username, password)
	if err != nil {
		return nil, err
	}

	return createApiClientWithToken(
		tokenResponse.AccessToken,
		tokenResponse.RefreshToken,
		options,
	), nil
}

func NewApiClient(accessToken, refreshToken string, options *Options) *ApiClient {
	return createApiClientWithToken(
		accessToken,
		refreshToken,
		options,
	)
}

// ensureValidToken checks if the access token is expired and refreshes it if necessary.
// It returns an error if the refresh token is also expired or if the refresh fails.
func (c *ApiClient) ensureValidToken() error {
	// Check if access token is still valid (with 30 second buffer to avoid edge cases)
	if time.Now().Add(30 * time.Second).Before(c.accessTokenExpiry) {
		return nil
	}

	// Check if auto-refresh is disabled
	if c.autoRefreshDisabled {
		return fmt.Errorf("automatic token refresh is disabled")
	}

	// Access token is expired or about to expire, need to refresh
	c.tokenMutex.Lock()
	defer c.tokenMutex.Unlock()

	// Double-check after acquiring lock (another goroutine might have already refreshed)
	if time.Now().Add(30 * time.Second).Before(c.accessTokenExpiry) {
		return nil
	}

	// Check if refresh token is still valid
	if time.Now().After(c.refreshTokenExpiry) {
		return fmt.Errorf("refresh token has expired, please re-authenticate")
	}

	// Check if we have client credentials to refresh
	if c.clientId == "" || c.clientSecret == "" {
		return fmt.Errorf("cannot refresh token: client credentials not available")
	}

	// Perform token refresh
	logger.Info("refreshing access token")
	tokenResponse, err := refreshToken(c.clientId, c.clientSecret, c.refreshToken)
	if err != nil {
		return fmt.Errorf("failed to refresh token: %w", err)
	}

	// Update tokens and expiry times
	c.accessToken = tokenResponse.AccessToken
	c.refreshToken = tokenResponse.RefreshToken
	c.accessTokenExpiry = time.Now().Add(time.Duration(tokenResponse.ExpiresIn) * time.Second)
	c.refreshTokenExpiry = time.Now().Add(time.Duration(tokenResponse.RefreshTokenExpiresIn) * time.Second)

	logger.Info("token refreshed successfully", "expiry", c.accessTokenExpiry)
	return nil
}

func (c *ApiClient) GetResourceUrl(path string, parametersData string) (string, error) {
	url := "https://members-ng.iracing.com" + path + "?" + parametersData

	// If concurrency is enabled, acquire the semaphore
	if c.concurrency > 0 {
		c.semaphore <- struct{}{}

		// Release the semaphore when done
		defer func() { <-c.semaphore }()
	}

	// TODO: max retries
	for {
		// Ensure we have a valid access token before making the request
		err := c.ensureValidToken()
		if err != nil {
			return "", fmt.Errorf("error ensuring valid token for %s: %w", path, err)
		}

		// Check if we are currently in a "Rate Limit Pause" state.
		c.mutex.Lock()
		waitTime := time.Until(c.retryAfter)
		c.mutex.Unlock()

		// If a global pause is active, sleep for the remaining duration.
		if waitTime > 0 {
			logger.Debug("pausing request due to rate limit", "path", path, "wait_time", waitTime)
			time.Sleep(waitTime)
		}

		// Perform the HTTP request
		req, err := http.NewRequest(http.MethodGet, url, nil)
		req.Header.Set("Authorization", "Bearer "+c.accessToken)
		if err != nil {
			return "", fmt.Errorf("error creating request for %s: %w", path, err)
		}

		resp, err := c.client.Do(req)
		if err != nil {
			return "", fmt.Errorf("error getting %s: %w", path, err)
		}
		defer resp.Body.Close()

		// Check for Rate Limiting (HTTP 429).
		if resp.StatusCode == 429 {
			// Read the X-RateLimit-Reset header.
			// Standard is usually a Unix timestamp (seconds).
			resetHeader := resp.Header.Get("X-RateLimit-Reset")
			resetTimeSec, _ := strconv.ParseInt(resetHeader, 10, 64)

			// Calculate when to wake up. Default to 5 seconds if header is missing/bad.
			wakeUpTime := time.Now().Add(5 * time.Second)
			if resetTimeSec > 0 {
				wakeUpTime = time.Unix(resetTimeSec, 0)
			}

			// Update the global lock so ALL other threads pause.
			c.mutex.Lock()
			if wakeUpTime.After(c.retryAfter) {
				c.retryAfter = wakeUpTime
				logger.Warn("rate limit exceeded", "path", path, "retry_after", wakeUpTime)
			}
			c.mutex.Unlock()

			// Continue the loop to RETRY this specific request after the sleep (in next iteration).
			continue
		}

		// Handle response
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("error reading response body for %s: %w", path, err)
		}

		if resp.StatusCode != 200 {
			return "", fmt.Errorf("error getting %s: %d - %s", path, resp.StatusCode, string(bodyBytes))
		}

		response := &iRacingResponse{}
		err = json.Unmarshal(bodyBytes, response)
		if err != nil {
			return "", fmt.Errorf("error unmarshaling response for %s: %w", path, err)
		}

		return response.Link, nil
	}
}

func (c *ApiClient) Get(path string, parametersData string) (*http.Response, error) {
	responseLink, err := c.GetResourceUrl(path, parametersData)
	if err != nil {
		return nil, err
	}

	payloadRes, err := http.DefaultClient.Get(responseLink)
	if err != nil {
		return nil, err
	}

	return payloadRes, nil
}
