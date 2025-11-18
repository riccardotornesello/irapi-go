package get

type SeriesGetResponse []SeriesGetResponseElement

type SeriesGetResponseElement struct {
	AllowedLicenses []AllowedLicense `json:"allowed_licenses"`
	Category        string           `json:"category"`
	CategoryID      int64            `json:"category_id"`
	Eligible        bool             `json:"eligible"`
	FirstSeason     FirstSeason      `json:"first_season"`
	ForumURL        *string          `json:"forum_url,omitempty"`
	MaxStarters     int64            `json:"max_starters"`
	MinStarters     int64            `json:"min_starters"`
	OvalCautionType int64            `json:"oval_caution_type"`
	RoadCautionType int64            `json:"road_caution_type"`
	SeriesID        int64            `json:"series_id"`
	SeriesName      string           `json:"series_name"`
	SeriesShortName string           `json:"series_short_name"`
	SearchFilters   *string          `json:"search_filters,omitempty"`
}

type AllowedLicense struct {
	GroupName       string `json:"group_name"`
	LicenseGroup    int64  `json:"license_group"`
	MaxLicenseLevel int64  `json:"max_license_level"`
	MinLicenseLevel int64  `json:"min_license_level"`
}

type FirstSeason struct {
	SeasonYear    int64 `json:"season_year"`
	SeasonQuarter int64 `json:"season_quarter"`
}
