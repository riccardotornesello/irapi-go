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

	// Client credentials for token refresh
	clientId     string
	clientSecret string

	client     *http.Client
	retryAfter time.Time

	concurrency int
	semaphore   chan struct{}
	mutex       sync.Mutex
	// tokenMutex is used to synchronize token refresh operations
	tokenMutex sync.Mutex
}

func createApiClientWithToken(accessToken, refreshToken string, accessTokenExpiry, refreshTokenExpiry time.Time, clientId, clientSecret string, concurrency int) *ApiClient {
	var semaphore chan struct{}
	if concurrency > 0 {
		semaphore = make(chan struct{}, concurrency)
	}

	return &ApiClient{
		accessToken:        accessToken,
		refreshToken:       refreshToken,
		accessTokenExpiry:  accessTokenExpiry,
		refreshTokenExpiry: refreshTokenExpiry,

		clientId:     clientId,
		clientSecret: clientSecret,

		client: &http.Client{
			Timeout: 60 * time.Second, // TODO: allow customization
		},

		concurrency: concurrency,
		semaphore:   semaphore,
		mutex:       sync.Mutex{},
		tokenMutex:  sync.Mutex{},
	}
}

func NewPasswordLimitedApiClient(clientId, clientSecret, username, password string) (*ApiClient, error) {
	tokenResponse, err := getPasswordLimitedAccessToken(clientId, clientSecret, username, password)
	if err != nil {
		return nil, err
	}

	return createApiClientWithToken(
		tokenResponse.AccessToken,
		tokenResponse.RefreshToken,
		time.Now().Add(time.Duration(tokenResponse.ExpiresIn)*time.Second),
		time.Now().Add(time.Duration(tokenResponse.RefreshTokenExpiresIn)*time.Second),
		clientId,
		clientSecret,
		10, // TODO: allow customization
	), nil
}

func NewApiClient(accessToken, refreshToken string) *ApiClient {
	// Extract expiry times from JTW tokens.
	expAccess, err := getExpiryFromJwt(accessToken)
	if err != nil {
		expAccess = time.Time{}
	}

	expRefresh, err := getExpiryFromJwt(refreshToken)
	if err != nil {
		expRefresh = time.Time{}
	}

	return createApiClientWithToken(
		accessToken,
		refreshToken,
		expAccess,
		expRefresh,
		"", // No client credentials for manual token client
		"", // No client credentials for manual token client
		10, // TODO: allow customization
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
