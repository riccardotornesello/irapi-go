package client

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/jszwec/csvutil"
	"github.com/riccardotornesello/irapi-go/pkg/logger"
)

type iRacingResponse struct {
	Link string `json:"link"`
}

type IRacingChunkInfo struct {
	ChunkSize       int      `json:"chunk_size"`
	NumChunks       int      `json:"num_chunks"`
	Rows            int      `json:"rows"`
	BaseDownloadUrl string   `json:"base_download_url"`
	ChunkFileNames  []string `json:"chunk_file_names"`
}

func (c *ApiClient) GetResourceUrl(path string, parameters interface{}) (string, error) {
	parametersData, err := query.Values(parameters)
	if err != nil {
		return "", fmt.Errorf("error encoding parameters for %s: %w", path, err)
	}

	url := "https://members-ng.iracing.com" + path + "?" + parametersData.Encode()

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

func GetChunks[T any](chunkInfo *IRacingChunkInfo) ([]T, error) {
	out := make([]T, 0)

	for i, chunkFileName := range chunkInfo.ChunkFileNames {
		resp, err := http.Get(chunkInfo.BaseDownloadUrl + chunkFileName)
		if err != nil {
			return nil, err
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		resp.Body.Close()

		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("error getting chunk %d: %d - %s", i, resp.StatusCode, string(bodyBytes))
		}

		var chunkData []T
		err = json.Unmarshal(bodyBytes, &chunkData)
		if err != nil {
			return nil, err
		}

		out = append(out, chunkData...)
	}

	return out, nil
}

func GetJson[T any](c *ApiClient, path string, parameters interface{}) (*T, error) {
	responseLink, err := c.GetResourceUrl(path, parameters)
	if err != nil {
		return nil, err
	}

	payloadRes, err := http.DefaultClient.Get(responseLink)
	if err != nil {
		return nil, err
	}
	defer payloadRes.Body.Close()

	respBody, err := io.ReadAll(payloadRes.Body)
	if err != nil {
		return nil, err
	}

	// Handle non-200 responses
	if payloadRes.StatusCode != 200 {
		return nil, fmt.Errorf("error getting JSON from %s: %d - %s", path, payloadRes.StatusCode, string(respBody))
	}

	response := new(T)
	err = json.Unmarshal(respBody, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func GetCsv[T any](c *ApiClient, path string, parameters interface{}) ([]T, error) {
	responseLink, err := c.GetResourceUrl(path, parameters)
	if err != nil {
		return nil, err
	}

	payloadRes, err := http.DefaultClient.Get(responseLink)
	if err != nil {
		return nil, err
	}
	defer payloadRes.Body.Close()

	reader := csv.NewReader(payloadRes.Body)

	var header T
	userHeader, _ := csvutil.Header(header, "csv")
	dec, _ := csvutil.NewDecoder(reader, userHeader...)

	var rows []T
	for {
		var row T
		if err := dec.Decode(&row); err == io.EOF {
			break
		}
		rows = append(rows, row)
	}

	return rows, nil
}
