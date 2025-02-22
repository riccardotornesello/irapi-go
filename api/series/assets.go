package series

import (
	"encoding/json"
)

type SeriesAssetsResponse map[string]struct {
	LargeImage interface{} `json:"large_image"`
	Logo       string      `json:"logo"`
	SeriesCopy string      `json:"series_copy"`
	SeriesId   int         `json:"series_id"`
	SmallImage interface{} `json:"small_image"`
}

// image paths are relative to https://images-static.iracing.com/
func (api *SeriesApi) GetSeriesAssets() (*SeriesAssetsResponse, error) {
	url := "/data/series/assets"

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &SeriesAssetsResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
