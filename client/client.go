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

	client     *http.Client
	retryAfter time.Time

	concurrency int
	semaphore   chan struct{}
	mutex       sync.Mutex
}

func createApiClientWithToken(accessToken, refreshToken string, accessTokenExpiry, refreshTokenExpiry time.Time, concurrency int) *ApiClient {
	var semaphore chan struct{}
	if concurrency > 0 {
		semaphore = make(chan struct{}, concurrency)
	}

	return &ApiClient{
		accessToken:        accessToken,
		refreshToken:       refreshToken,
		accessTokenExpiry:  accessTokenExpiry,
		refreshTokenExpiry: refreshTokenExpiry,

		client: &http.Client{
			Timeout: 60 * time.Second, // TODO: allow customization
		},

		concurrency: concurrency,
		semaphore:   semaphore,
		mutex:       sync.Mutex{},
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
