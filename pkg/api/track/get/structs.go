package get

import (
	"github.com/riccardotornesello/irapi-go/pkg/types"
)

type TrackGetResponse []TrackGetResponseElement

type TrackGetResponseElement struct {
	AIEnabled              bool           `json:"ai_enabled"`
	AllowPitlaneCollisions bool           `json:"allow_pitlane_collisions"`
	AllowRollingStart      bool           `json:"allow_rolling_start"`
	AllowStandingStart     bool           `json:"allow_standing_start"`
	AwardExempt            bool           `json:"award_exempt"`
	Category               string         `json:"category"`
	CategoryID             int64          `json:"category_id"`
	Closes                 string         `json:"closes"`
	ConfigName             *string        `json:"config_name,omitempty"`
	CornersPerLap          int64          `json:"corners_per_lap"`
	Created                types.DateTime `json:"created"`
	FirstSale              types.DateTime `json:"first_sale"`
	Folder                 string         `json:"folder"`
	FreeWithSubscription   bool           `json:"free_with_subscription"`
	FullyLit               bool           `json:"fully_lit"`
	GridStalls             int64          `json:"grid_stalls"`
	HasOptPath             bool           `json:"has_opt_path"`
	HasShortParadeLap      bool           `json:"has_short_parade_lap"`
	HasStartZone           bool           `json:"has_start_zone"`
	HasSVGMap              bool           `json:"has_svg_map"`
	IsDirt                 bool           `json:"is_dirt"`
	IsOval                 bool           `json:"is_oval"`
	IsPSPurchasable        bool           `json:"is_ps_purchasable"`
	LapScoring             int64          `json:"lap_scoring"`
	Latitude               float64        `json:"latitude"`
	Location               string         `json:"location"`
	Logo                   string         `json:"logo"`
	Longitude              float64        `json:"longitude"`
	MaxCars                int64          `json:"max_cars"`
	NightLighting          bool           `json:"night_lighting"`
	NumberPitstalls        int64          `json:"number_pitstalls"`
	Opens                  string         `json:"opens"`
	PackageID              int64          `json:"package_id"`
	PitRoadSpeedLimit      *int64         `json:"pit_road_speed_limit,omitempty"`
	Price                  float64        `json:"price"`
	Priority               int64          `json:"priority"`
	Purchasable            bool           `json:"purchasable"`
	QualifyLaps            int64          `json:"qualify_laps"`
	RainEnabled            bool           `json:"rain_enabled"`
	RestartOnLeft          bool           `json:"restart_on_left"`
	Retired                bool           `json:"retired"`
	SearchFilters          string         `json:"search_filters"`
	SiteURL                *string        `json:"site_url,omitempty"`
	Sku                    int64          `json:"sku"`
	SmallImage             string         `json:"small_image"`
	SoloLaps               int64          `json:"solo_laps"`
	StartOnLeft            bool           `json:"start_on_left"`
	SupportsGripCompound   bool           `json:"supports_grip_compound"`
	TechTrack              bool           `json:"tech_track"`
	TimeZone               string         `json:"time_zone"`
	TrackConfigLength      float64        `json:"track_config_length"`
	TrackDirpath           string         `json:"track_dirpath"`
	TrackID                int64          `json:"track_id"`
	TrackName              string         `json:"track_name"`
	TrackType              int64          `json:"track_type"`
	TrackTypeText          string         `json:"track_type_text"`
	TrackTypes             []TrackType    `json:"track_types"`
	PriceDisplay           *string        `json:"price_display,omitempty"`
	NominalLapTime         *float64       `json:"nominal_lap_time,omitempty"`
	Banking                *string        `json:"banking,omitempty"`
}

type TrackType struct {
	TrackType string `json:"track_type"`
}
