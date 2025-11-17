package stats_series

import ()

type SeriesStatsSeriesResponse []SeriesStatsSeriesResponseElement

type SeriesStatsSeriesResponseElement struct {
	SeriesID          int64              `json:"series_id"`
	SeriesName        string             `json:"series_name"`
	SeriesShortName   string             `json:"series_short_name"`
	CategoryID        int64              `json:"category_id"`
	Category          string             `json:"category"`
	Active            bool               `json:"active"`
	Official          bool               `json:"official"`
	FixedSetup        bool               `json:"fixed_setup"`
	Logo              *string            `json:"logo"`
	LicenseGroup      int64              `json:"license_group"`
	LicenseGroupTypes []LicenseGroupType `json:"license_group_types"`
	AllowedLicenses   []AllowedLicense   `json:"allowed_licenses"`
	Seasons           []Season           `json:"seasons"`
	SearchFilters     *string            `json:"search_filters,omitempty"`
}

type AllowedLicense struct {
	GroupName       string `json:"group_name"`
	LicenseGroup    int64  `json:"license_group"`
	MaxLicenseLevel int64  `json:"max_license_level"`
	MinLicenseLevel int64  `json:"min_license_level"`
	ParentID        int64  `json:"parent_id"`
}

type LicenseGroupType struct {
	LicenseGroupType int64 `json:"license_group_type"`
}

type Season struct {
	SeasonID          int64              `json:"season_id"`
	SeriesID          int64              `json:"series_id"`
	SeasonName        string             `json:"season_name"`
	SeasonShortName   string             `json:"season_short_name"`
	SeasonYear        int64              `json:"season_year"`
	SeasonQuarter     int64              `json:"season_quarter"`
	Active            bool               `json:"active"`
	Official          bool               `json:"official"`
	DriverChanges     bool               `json:"driver_changes"`
	FixedSetup        bool               `json:"fixed_setup"`
	LicenseGroup      int64              `json:"license_group"`
	HasSupersessions  bool               `json:"has_supersessions"`
	CarSwitching      bool               `json:"car_switching"`
	LicenseGroupTypes []LicenseGroupType `json:"license_group_types"`
	CarClasses        []CarClass         `json:"car_classes"`
	RaceWeeks         []RaceWeek         `json:"race_weeks"`
}

type CarClass struct {
	CarClassID    int64  `json:"car_class_id"`
	ShortName     string `json:"short_name"`
	Name          string `json:"name"`
	RelativeSpeed int64  `json:"relative_speed"`
}

type RaceWeek struct {
	SeasonID    int64 `json:"season_id"`
	RaceWeekNum int64 `json:"race_week_num"`
	Track       Track `json:"track"`
}

type Track struct {
	ConfigName *string `json:"config_name,omitempty"`
	TrackID    int64   `json:"track_id"`
	TrackName  string  `json:"track_name"`
}
