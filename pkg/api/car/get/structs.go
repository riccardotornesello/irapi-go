package get

import (
	"github.com/riccardotornesello/irapi-go/pkg/types"
)

type CarGetResponse []CarGetResponseElement

type CarGetResponseElement struct {
	AIEnabled               bool           `json:"ai_enabled"`
	AllowNumberColors       bool           `json:"allow_number_colors"`
	AllowNumberFont         bool           `json:"allow_number_font"`
	AllowSponsor1           bool           `json:"allow_sponsor1"`
	AllowSponsor2           bool           `json:"allow_sponsor2"`
	AllowWheelColor         bool           `json:"allow_wheel_color"`
	AwardExempt             bool           `json:"award_exempt"`
	CarConfigDefs           []CarConfigDef `json:"car_config_defs"`
	CarConfigs              []CarConfig    `json:"car_configs"`
	CarDirpath              string         `json:"car_dirpath"`
	CarID                   int64          `json:"car_id"`
	CarName                 string         `json:"car_name"`
	CarNameAbbreviated      string         `json:"car_name_abbreviated"`
	CarTypes                []CarType      `json:"car_types"`
	CarWeight               int64          `json:"car_weight"`
	Categories              []string       `json:"categories"`
	Created                 types.DateTime `json:"created"`
	FirstSale               types.DateTime `json:"first_sale"`
	Folder                  string         `json:"folder"`
	ForumURL                *string        `json:"forum_url,omitempty"`
	FreeWithSubscription    bool           `json:"free_with_subscription"`
	HasHeadlights           bool           `json:"has_headlights"`
	HasMultipleDryTireTypes bool           `json:"has_multiple_dry_tire_types"`
	HasRainCapableTireTypes bool           `json:"has_rain_capable_tire_types"`
	HP                      int64          `json:"hp"`
	IsPSPurchasable         bool           `json:"is_ps_purchasable"`
	Logo                    *string        `json:"logo"`
	MaxPowerAdjustPct       int64          `json:"max_power_adjust_pct"`
	MaxWeightPenaltyKg      int64          `json:"max_weight_penalty_kg"`
	MinPowerAdjustPct       int64          `json:"min_power_adjust_pct"`
	PackageID               int64          `json:"package_id"`
	Patterns                int64          `json:"patterns"`
	Price                   float64        `json:"price"`
	PriceDisplay            *string        `json:"price_display,omitempty"`
	RainEnabled             bool           `json:"rain_enabled"`
	Retired                 bool           `json:"retired"`
	SearchFilters           string         `json:"search_filters"`
	Sku                     int64          `json:"sku"`
	SmallImage              string         `json:"small_image"`
	SponsorLogo             *string        `json:"sponsor_logo"`
	CarMake                 *string        `json:"car_make,omitempty"`
	CarModel                *string        `json:"car_model,omitempty"`
	PaintRules              *PaintRules    `json:"paint_rules,omitempty"`
	SiteURL                 *string        `json:"site_url,omitempty"`
}

type CarConfigDef struct {
	Carcfg         int64   `json:"carcfg"`
	CFGSubdir      *string `json:"cfg_subdir"`
	CustomPaintEXT *string `json:"custom_paint_ext"`
	Name           string  `json:"name"`
}

type CarConfig struct {
	Carcfg    int64  `json:"carcfg"`
	TrackID   *int64 `json:"track_id,omitempty"`
	TrackType *int64 `json:"track_type,omitempty"`
}

type CarType struct {
	CarType string `json:"car_type"`
}

type PaintRules struct {
	The1                *The0  `json:"1,omitempty"`
	RestrictCustomPaint *bool  `json:"RestrictCustomPaint,omitempty"`
	The24               *The24 `json:"24,omitempty"`
	The2                *The0  `json:"2,omitempty"`
	The3                *The0  `json:"3,omitempty"`
	The4                *The0  `json:"4,omitempty"`
	The5                *The0  `json:"5,omitempty"`
	The0                *The0  `json:"0,omitempty"`
	The6                *The0  `json:"6,omitempty"`
	The7                *The0  `json:"7,omitempty"`
	The8                *The0  `json:"8,omitempty"`
	The9                *The0  `json:"9,omitempty"`
}

type The0 struct {
	PaintCarAvailable       bool    `json:"PaintCarAvailable"`
	Color1                  string  `json:"Color1"`
	Color2                  string  `json:"Color2"`
	Color3                  string  `json:"Color3"`
	Sponsor1Available       bool    `json:"Sponsor1Available"`
	Sponsor2Available       bool    `json:"Sponsor2Available"`
	Sponsor1                string  `json:"Sponsor1"`
	Sponsor2                string  `json:"Sponsor2"`
	PaintWheelAvailable     *bool   `json:"PaintWheelAvailable,omitempty"`
	WheelColor              *string `json:"WheelColor,omitempty"`
	RimTypeAvailable        *bool   `json:"RimTypeAvailable,omitempty"`
	RimType                 *string `json:"RimType,omitempty"`
	AllowNumberFontChanges  *bool   `json:"AllowNumberFontChanges,omitempty"`
	NumberFont              *string `json:"NumberFont,omitempty"`
	AllowNumberColorChanges *bool   `json:"AllowNumberColorChanges,omitempty"`
	NumberColor1            *string `json:"NumberColor1,omitempty"`
	NumberColor2            *string `json:"NumberColor2,omitempty"`
	NumberColor3            *string `json:"NumberColor3,omitempty"`
	RulesExplanation        string  `json:"RulesExplanation"`
}

type The24 struct {
	PaintCarAvailable bool   `json:"PaintCarAvailable"`
	Color1            string `json:"Color1"`
	Color2            string `json:"Color2"`
	Color3            string `json:"Color3"`
	Sponsor1Available bool   `json:"Sponsor1Available"`
	Sponsor2Available bool   `json:"Sponsor2Available"`
	Sponsor1          string `json:"Sponsor1"`
	Sponsor2          string `json:"Sponsor2"`
	RulesExplanation  string `json:"RulesExplanation"`
}
