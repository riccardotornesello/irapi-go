package series

import (
	"encoding/json"
)

type SeriesStatsSeriesResponse []struct {
	SeriesId          int     `json:"series_id"`
	SeriesName        string  `json:"series_name"`
	SeriesShortName   string  `json:"series_short_name"`
	CategoryId        int     `json:"category_id"`
	Category          string  `json:"category"`
	Active            bool    `json:"active"`
	Official          bool    `json:"official"`
	FixedSetup        bool    `json:"fixed_setup"`
	SearchFilters     string  `json:"search_filters"`
	Logo              *string `json:"logo"`
	LicenseGroup      int     `json:"license_group"`
	LicenseGroupTypes []struct {
		LicenseGroupType int `json:"license_group_type"`
	} `json:"license_group_types"`
	AllowedLicenses []struct {
		GroupName       string `json:"group_name"`
		LicenseGroup    int    `json:"license_group"`
		MaxLicenseLevel int    `json:"max_license_level"`
		MinLicenseLevel int    `json:"min_license_level"`
		ParentId        int    `json:"parent_id"`
	} `json:"allowed_licenses"`
	Seasons []struct {
		SeasonId          int    `json:"season_id"`
		SeriesId          int    `json:"series_id"`
		SeasonName        string `json:"season_name"`
		SeasonShortName   string `json:"season_short_name"`
		SeasonYear        int    `json:"season_year"`
		SeasonQuarter     int    `json:"season_quarter"`
		Active            bool   `json:"active"`
		Official          bool   `json:"official"`
		DriverChanges     bool   `json:"driver_changes"`
		FixedSetup        bool   `json:"fixed_setup"`
		LicenseGroup      int    `json:"license_group"`
		HasSupersessions  bool   `json:"has_supersessions"`
		LicenseGroupTypes []struct {
			LicenseGroupType int `json:"license_group_type"`
		} `json:"license_group_types"`
		CarClasses []struct {
			CarClassId    int    `json:"car_class_id"`
			ShortName     string `json:"short_name"`
			Name          string `json:"name"`
			RelativeSpeed int    `json:"relative_speed"`
		} `json:"car_classes"`
		RaceWeeks []struct {
			SeasonId    int `json:"season_id"`
			RaceWeekNum int `json:"race_week_num"`
			Track       struct {
				ConfigName string `json:"config_name"`
				TrackId    int    `json:"track_id"`
				TrackName  string `json:"track_name"`
			} `json:"track"`
		} `json:"race_weeks"`
	} `json:"seasons"`
}

// To get series and seasons for which standings should be available, filter the list by official: true.
func (api *SeriesApi) GetSeriesStatsSeries() (*SeriesStatsSeriesResponse, error) {
	url := "/data/series/stats_series"

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &SeriesStatsSeriesResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
