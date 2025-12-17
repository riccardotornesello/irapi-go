package client

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/jszwec/csvutil"
)

type IRacingChunkInfo struct {
	ChunkSize       int      `json:"chunk_size"`
	NumChunks       int      `json:"num_chunks"`
	Rows            int      `json:"rows"`
	BaseDownloadUrl string   `json:"base_download_url"`
	ChunkFileNames  []string `json:"chunk_file_names"`
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
	parametersData, err := query.Values(parameters)
	if err != nil {
		return nil, fmt.Errorf("error encoding parameters for %s: %w", path, err)
	}

	payloadRes, err := c.Get(path, parametersData.Encode())
	if err != nil {
		return nil, err
	}
	defer payloadRes.Body.Close()

	// Handle non-200 responses
	if payloadRes.StatusCode != 200 {
		respBody, err := io.ReadAll(payloadRes.Body)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("error getting JSON from %s: %d - %s", path, payloadRes.StatusCode, string(respBody))
	}

	response := new(T)
	err = json.NewDecoder(payloadRes.Body).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func GetCsv[T any](c *ApiClient, path string, parameters interface{}) ([]T, error) {
	parametersData, err := query.Values(parameters)
	if err != nil {
		return nil, fmt.Errorf("error encoding parameters for %s: %w", path, err)
	}

	payloadRes, err := c.Get(path, parametersData.Encode())
	if err != nil {
		return nil, err
	}
	defer payloadRes.Body.Close()

	// Handle non-200 responses
	if payloadRes.StatusCode != 200 {
		respBody, err := io.ReadAll(payloadRes.Body)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("error getting JSON from %s: %d - %s", path, payloadRes.StatusCode, string(respBody))
	}

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
