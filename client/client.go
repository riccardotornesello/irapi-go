package client

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
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
	// tokenMutex is used to synchronize token refresh operations
	tokenMutex sync.Mutex
}

type ApiClientOptions struct {
	Concurrency  int
	ClientId     string
	ClientSecret string
}

func createApiClientWithToken(accessToken, refreshToken string, options *ApiClientOptions) *ApiClient {
	// Use default options if none provided
	if options == nil {
		options = &ApiClientOptions{}
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
		tokenMutex:  sync.Mutex{},
	}
}

func NewPasswordLimitedApiClient(username, password string, options *ApiClientOptions) (*ApiClient, error) {
	if options == nil {
		options = &ApiClientOptions{}
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

func NewApiClient(accessToken, refreshToken string, options *ApiClientOptions) *ApiClient {
	return createApiClientWithToken(
		accessToken,
		refreshToken,
		options,
	)
}

func getExpiryFromJwt(token string) (time.Time, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return time.Time{}, fmt.Errorf("invalid JWT token")
	}

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return time.Time{}, err
	}

	var claims struct {
		Exp int64 `json:"exp"`
	}
	err = json.Unmarshal(payload, &claims)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(claims.Exp, 0), nil
}

// ensureValidToken checks if the access token is expired and refreshes it if necessary.
// It returns an error if the refresh token is also expired or if the refresh fails.
func (c *ApiClient) ensureValidToken() error {
	// Check if access token is still valid (with 30 second buffer to avoid edge cases)
	if time.Now().Add(30 * time.Second).Before(c.accessTokenExpiry) {
		return nil
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
	fmt.Printf("[Token Refresh] Refreshing access token...\n")
	tokenResponse, err := refreshToken(c.clientId, c.clientSecret, c.refreshToken)
	if err != nil {
		return fmt.Errorf("failed to refresh token: %w", err)
	}

	// Update tokens and expiry times
	c.accessToken = tokenResponse.AccessToken
	c.refreshToken = tokenResponse.RefreshToken
	c.accessTokenExpiry = time.Now().Add(time.Duration(tokenResponse.ExpiresIn) * time.Second)
	c.refreshTokenExpiry = time.Now().Add(time.Duration(tokenResponse.RefreshTokenExpiresIn) * time.Second)

	fmt.Printf("[Token Refresh] Token refreshed successfully. New expiry: %v\n", c.accessTokenExpiry)
	return nil
}
