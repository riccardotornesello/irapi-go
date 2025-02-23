package client

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const tokenCookieName = "authtoken_members"

type ApiClient struct {
	Client     *http.Client
	RetryAfter time.Time
}

type IRacingAuthResponse struct {
	Authcode string `json:"authcode"`
}

type IRacingResponse struct {
	Link string `json:"link"`
}

type IRacingChunkInfo struct {
	ChunkSize       int      `json:"chunk_size"`
	NumChunks       int      `json:"num_chunks"`
	Rows            int      `json:"rows"`
	BaseDownloadUrl string   `json:"base_download_url"`
	ChunkFileNames  []string `json:"chunk_file_names"`
}

func NewApiClient(email string, password string) (*ApiClient, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Jar: jar,
	}

	tokenIn := []byte(password + strings.ToLower(email))
	hasher := sha256.New()
	hasher.Write(tokenIn)
	tokenHash := hasher.Sum(nil)
	tokenB64 := base64.StdEncoding.EncodeToString(tokenHash)

	resp, err := client.Post("https://members-ng.iracing.com/auth", "application/json", strings.NewReader(`{"email":"`+email+`","password":"`+tokenB64+`"}`))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error authenticating: %s", resp.Status)
	}

	defer resp.Body.Close()

	authResponse := &IRacingAuthResponse{}
	err = json.NewDecoder(resp.Body).Decode(authResponse)
	if err != nil {
		return nil, err
	}

	return &ApiClient{
		Client: client,
	}, nil
}

func NewApiClientWithToken(token string) *ApiClient {
	jar, _ := cookiejar.New(nil)

	// Set a cookie in the jar
	u, _ := url.Parse("https://members-ng.iracing.com")
	jar.SetCookies(u, []*http.Cookie{
		{
			Name:  tokenCookieName,
			Value: token,
		},
	})

	client := &http.Client{
		Jar: jar,
	}

	return &ApiClient{
		Client: client,
	}
}

func (c *ApiClient) Get(path string) (io.ReadCloser, error) {
	var resp *http.Response
	var err error

	for {
		if c.RetryAfter.After(time.Now()) {
			slog.Info(fmt.Sprintf("Rate limit exceeded, waiting until %v", c.RetryAfter.Format(time.RFC3339)))
			time.Sleep(time.Until(c.RetryAfter))
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second) // TODO: allow timeout customization
		defer cancel()

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://members-ng.iracing.com"+path, nil)
		if err != nil {
			return nil, fmt.Errorf("error creating request for %s: %w", path, err)
		}

		resp, err = c.Client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("error getting %s: %w", path, err)
		}

		if resp.StatusCode != 429 {
			break
		}

		slog.Info(fmt.Sprintf("Rate limit exceeded for %s, retrying in a bit", path))

		// TODO: allow to skip retrying
		// TODO: allow max retry count
		rateLimitReset := resp.Header.Get("X-RateLimit-Reset")
		if rateLimitReset == "" {
			break
		}

		rateLimitResetInt, err := strconv.ParseInt(rateLimitReset, 10, 64)
		if err != nil {
			break
		}

		// Not atomic, but we don't care
		c.RetryAfter = time.Unix(rateLimitResetInt, 0).Add(2 * time.Second)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error getting %s: %d", path, resp.StatusCode)
		}
		return nil, fmt.Errorf("error getting %s: %d - %s", path, resp.StatusCode, string(body))
	}

	response := &IRacingResponse{}
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return nil, err
	}

	payloadResp, err := c.Client.Get(response.Link)
	if err != nil {
		return nil, err
	}

	return payloadResp.Body, nil
}

func (c *ApiClient) GetChunks(chunkInfo *IRacingChunkInfo) ([]io.ReadCloser, error) {
	out := make([]io.ReadCloser, len(chunkInfo.ChunkFileNames))

	for i, chunkFileName := range chunkInfo.ChunkFileNames {
		resp, err := c.Client.Get(chunkInfo.BaseDownloadUrl + chunkFileName)
		if err != nil {
			return nil, err
		}

		out[i] = resp.Body
	}

	return out, nil
}

func (c *ApiClient) GetAuthToken() string {
	url, _ := url.Parse("https://members-ng.iracing.com")
	for _, cookie := range c.Client.Jar.Cookies(url) {
		if cookie.Name == tokenCookieName {
			return cookie.Value
		}
	}

	return ""
}
