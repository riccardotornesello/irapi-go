package track

import (
	"encoding/json"
)

type TrackGetResponse []struct {
	AiEnabled              bool    `json:"ai_enabled"`
	AllowPitlaneCollisions bool    `json:"allow_pitlane_collisions"`
	AllowRollingStart      bool    `json:"allow_rolling_start"`
	AllowStandingStart     bool    `json:"allow_standing_start"`
	AwardExempt            bool    `json:"award_exempt"`
	Category               string  `json:"category"`
	CategoryId             int     `json:"category_id"`
	Closes                 string  `json:"closes"`
	ConfigName             string  `json:"config_name"`
	CornersPerLap          int     `json:"corners_per_lap"`
	Created                string  `json:"created"`
	FirstSale              string  `json:"first_sale"`
	Folder                 string  `json:"folder"`
	FreeWithSubscription   bool    `json:"free_with_subscription"`
	FullyLit               bool    `json:"fully_lit"`
	GridStalls             int     `json:"grid_stalls"`
	HasOptPath             bool    `json:"has_opt_path"`
	HasShortParadeLap      bool    `json:"has_short_parade_lap"`
	HasStartZone           bool    `json:"has_start_zone"`
	HasSvgMap              bool    `json:"has_svg_map"`
	IsDirt                 bool    `json:"is_dirt"`
	IsOval                 bool    `json:"is_oval"`
	IsPsPurchasable        bool    `json:"is_ps_purchasable"`
	LapScoring             int     `json:"lap_scoring"`
	Latitude               float64 `json:"latitude"`
	Location               string  `json:"location"`
	Longitude              float64 `json:"longitude"`
	MaxCars                int     `json:"max_cars"`
	NightLighting          bool    `json:"night_lighting"`
	NominalLapTime         float64 `json:"nominal_lap_time"`
	NumberPitstalls        int     `json:"number_pitstalls"`
	Opens                  string  `json:"opens"`
	PackageId              int     `json:"package_id"`
	PitRoadSpeedLimit      int     `json:"pit_road_speed_limit"`
	Price                  float64 `json:"price"`
	PriceDisplay           string  `json:"price_display"`
	Priority               int     `json:"priority"`
	Purchasable            bool    `json:"purchasable"`
	QualifyLaps            int     `json:"qualify_laps"`
	RainEnabled            bool    `json:"rain_enabled"`
	RestartOnLeft          bool    `json:"restart_on_left"`
	Retired                bool    `json:"retired"`
	SearchFilters          string  `json:"search_filters"`
	SiteUrl                string  `json:"site_url"`
	Sku                    int     `json:"sku"`
	SmallImage             string  `json:"small_image"`
	SoloLaps               int     `json:"solo_laps"`
	StartOnLeft            bool    `json:"start_on_left"`
	SupportsGripCompound   bool    `json:"supports_grip_compound"`
	TechTrack              bool    `json:"tech_track"`
	TimeZone               string  `json:"time_zone"`
	TrackConfigLength      float64 `json:"track_config_length"`
	TrackDirpath           string  `json:"track_dirpath"`
	TrackId                int     `json:"track_id"`
	TrackName              string  `json:"track_name"`
	TrackTypes             []struct {
		TrackType string `json:"track_type"`
	} `json:"track_types"`
	Banking string `json:"banking"`
}

func (api *TrackApi) GetTrack() (*TrackGetResponse, error) {
	url := "/data/track/get"

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &TrackGetResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
