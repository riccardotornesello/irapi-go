package series

import (
	"encoding/json"
)

type SeriesAssetsResponse map[string]SeriesAssetsResponseValue

type SeriesAssetsResponseValue struct {
	LargeImage interface{} `json:"large_image"`
	Logo       string      `json:"logo"`
	SeriesCopy string      `json:"series_copy"`
	SeriesID   int64       `json:"series_id"`
	SmallImage interface{} `json:"small_image"`
}

func (api *SeriesApi) Assets() (*SeriesAssetsResponse, error) {
	return api.GetJson[SeriesAssetsResponse]("/data/series/assets")
}
