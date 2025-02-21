package series

import (
	"encoding/json"
)

type SeriesGetResponse []struct {
	AllowedLicenses []struct {
		GroupName       string `json:"group_name"`
		LicenseGroup    int    `json:"license_group"`
		MaxLicenseLevel int    `json:"max_license_level"`
		MinLicenseLevel int    `json:"min_license_level"`
	} `json:"allowed_licenses"`
	Category    string `json:"category"`
	CategoryId  int    `json:"category_id"`
	Eligible    bool   `json:"eligible"`
	FirstSeason struct {
		SeasonYear    int `json:"season_year"`
		SeasonQuarter int `json:"season_quarter"`
	} `json:"first_season"`
	ForumUrl        string `json:"forum_url"`
	MaxStarters     int    `json:"max_starters"`
	MinStarters     int    `json:"min_starters"`
	OvalCautionType int    `json:"oval_caution_type"`
	RoadCautionType int    `json:"road_caution_type"`
	SeriesId        int    `json:"series_id"`
	SeriesName      string `json:"series_name"`
	SeriesShortName string `json:"series_short_name"`
	SearchFilters   string `json:"search_filters"`
}

func (api *SeriesApi) GetSeries() (*SeriesGetResponse, error) {
	url := "https://members-ng.iracing.com/data/series/get"

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &SeriesGetResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
