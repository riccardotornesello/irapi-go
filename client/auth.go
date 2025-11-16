package client

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const tokenUrl = "https://oauth.iracing.com/oauth2/token"

type iRacingTokenResponse struct {
	AccessToken           string `json:"access_token"`
	TokenType             string `json:"token_type"`
	ExpiresIn             int    `json:"expires_in"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresIn int    `json:"refresh_token_expires_in"`
}

func mask(secret string, id string) string {
	normalizedID := strings.ToLower(strings.TrimSpace(id))
	sum := sha256.Sum256([]byte(secret + normalizedID))
	return base64.StdEncoding.EncodeToString(sum[:])
}

func getPasswordLimitedAccessToken(clientId, clientSecret, username, password string) (*iRacingTokenResponse, error) {
	res, err := http.PostForm(tokenUrl, map[string][]string{
		"grant_type":    {"password_limited"},
		"client_id":     {clientId},
		"client_secret": {mask(clientSecret, clientId)},
		"username":      {username},
		"password":      {mask(password, username)},
		"scope":         {"iracing.auth"},
	})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Non-OK HTTP status: %d - %s", res.StatusCode, string(body))
	}

	tokenResponse := &iRacingTokenResponse{}
	err = json.Unmarshal(body, tokenResponse)
	if err != nil {
		return nil, err
	}

	return tokenResponse, nil
}

func refreshToken(clientId, clientSecret, refreshToken string) (*iRacingTokenResponse, error) {
	res, err := http.PostForm(tokenUrl, map[string][]string{
		"grant_type":    {"refresh_token"},
		"client_id":     {clientId},
		"client_secret": {mask(clientSecret, clientId)},
		"refresh_token": {refreshToken},
	})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Non-OK HTTP status: %d - %s", res.StatusCode, string(body))
	}

	tokenResponse := &iRacingTokenResponse{}
	err = json.Unmarshal(body, tokenResponse)
	if err != nil {
		return nil, err
	}

	return tokenResponse, nil
}
