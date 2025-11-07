package get

import (
	"time"
)

type ResultsGetParams struct {
	SubsessionId    int  `url:"subsession_id,omitempty,comma"`
	IncludeLicenses bool `url:"include_licenses,omitempty,comma"`
}

type ResultsGetResponse struct {
	SubsessionID            int64           `json:"subsession_id"`
	AssociatedSubsessionIDS []int64         `json:"associated_subsession_ids"`
	CanProtest              bool            `json:"can_protest"`
	CarClasses              []CarClass      `json:"car_classes"`
	CautionType             int64           `json:"caution_type"`
	CooldownMinutes         int64           `json:"cooldown_minutes"`
	CornersPerLap           int64           `json:"corners_per_lap"`
	DamageModel             int64           `json:"damage_model"`
	DriverChangeParam1      int64           `json:"driver_change_param1"`
	DriverChangeParam2      int64           `json:"driver_change_param2"`
	DriverChangeRule        int64           `json:"driver_change_rule"`
	DriverChanges           bool            `json:"driver_changes"`
	EndTime                 time.Time       `json:"end_time"`
	EventAverageLap         int64           `json:"event_average_lap"`
	EventBestLapTime        int64           `json:"event_best_lap_time"`
	EventLapsComplete       int64           `json:"event_laps_complete"`
	EventStrengthOfField    int64           `json:"event_strength_of_field"`
	EventType               int64           `json:"event_type"`
	EventTypeName           string          `json:"event_type_name"`
	HeatInfoID              int64           `json:"heat_info_id"`
	HostID                  int64           `json:"host_id"`
	LeagueID                int64           `json:"league_id"`
	LeagueName              string          `json:"league_name"`
	LeagueSeasonID          int64           `json:"league_season_id"`
	LicenseCategory         string          `json:"license_category"`
	LicenseCategoryID       int64           `json:"license_category_id"`
	LimitMinutes            int64           `json:"limit_minutes"`
	MaxTeamDrivers          int64           `json:"max_team_drivers"`
	MaxWeeks                int64           `json:"max_weeks"`
	MinTeamDrivers          int64           `json:"min_team_drivers"`
	NumCautionLaps          int64           `json:"num_caution_laps"`
	NumCautions             int64           `json:"num_cautions"`
	NumDrivers              int64           `json:"num_drivers"`
	NumLapsForQualAverage   int64           `json:"num_laps_for_qual_average"`
	NumLapsForSoloAverage   int64           `json:"num_laps_for_solo_average"`
	NumLeadChanges          int64           `json:"num_lead_changes"`
	OfficialSession         bool            `json:"official_session"`
	PointsType              string          `json:"points_type"`
	PrivateSessionID        int64           `json:"private_session_id"`
	RaceWeekNum             int64           `json:"race_week_num"`
	RestrictResults         bool            `json:"restrict_results"`
	ResultsRestricted       bool            `json:"results_restricted"`
	SeasonID                int64           `json:"season_id"`
	SeasonName              string          `json:"season_name"`
	SeasonQuarter           int64           `json:"season_quarter"`
	SeasonShortName         string          `json:"season_short_name"`
	SeasonYear              int64           `json:"season_year"`
	SeriesID                int64           `json:"series_id"`
	SeriesName              string          `json:"series_name"`
	SeriesShortName         string          `json:"series_short_name"`
	SessionID               int64           `json:"session_id"`
	SessionName             string          `json:"session_name"`
	SessionResults          []SessionResult `json:"session_results"`
	SessionSplits           []SessionSplit  `json:"session_splits"`
	SpecialEventType        int64           `json:"special_event_type"`
	StartTime               time.Time       `json:"start_time"`
	Track                   Track           `json:"track"`
	TrackState              TrackState      `json:"track_state"`
	Weather                 Weather         `json:"weather"`
}

type CarClass struct {
	CarClassID      int64         `json:"car_class_id"`
	ShortName       ShortName     `json:"short_name"`
	Name            Name          `json:"name"`
	StrengthOfField int64         `json:"strength_of_field"`
	NumEntries      int64         `json:"num_entries"`
	CarsInClass     []CarsInClass `json:"cars_in_class"`
}

type CarsInClass struct {
	CarID int64 `json:"car_id"`
}

type SessionResult struct {
	SimsessionNumber   int64    `json:"simsession_number"`
	SimsessionName     string   `json:"simsession_name"`
	SimsessionType     int64    `json:"simsession_type"`
	SimsessionTypeName string   `json:"simsession_type_name"`
	SimsessionSubtype  int64    `json:"simsession_subtype"`
	Results            []Result `json:"results"`
}

type Result struct {
	CustID                  int64          `json:"cust_id"`
	DisplayName             string         `json:"display_name"`
	AggregateChampPoints    int64          `json:"aggregate_champ_points"`
	AI                      bool           `json:"ai"`
	AverageLap              int64          `json:"average_lap"`
	BestLapNum              int64          `json:"best_lap_num"`
	BestLapTime             int64          `json:"best_lap_time"`
	BestNlapsNum            int64          `json:"best_nlaps_num"`
	BestNlapsTime           int64          `json:"best_nlaps_time"`
	BestQualLapAt           time.Time      `json:"best_qual_lap_at"`
	BestQualLapNum          int64          `json:"best_qual_lap_num"`
	BestQualLapTime         int64          `json:"best_qual_lap_time"`
	CarClassID              int64          `json:"car_class_id"`
	CarClassName            Name           `json:"car_class_name"`
	CarClassShortName       ShortName      `json:"car_class_short_name"`
	CarID                   int64          `json:"car_id"`
	CarName                 CarName        `json:"car_name"`
	Carcfg                  int64          `json:"carcfg"`
	ChampPoints             int64          `json:"champ_points"`
	ClassInterval           int64          `json:"class_interval"`
	CountryCode             CountryCode    `json:"country_code"`
	Division                int64          `json:"division"`
	DropRace                bool           `json:"drop_race"`
	FinishPosition          int64          `json:"finish_position"`
	FinishPositionInClass   int64          `json:"finish_position_in_class"`
	FlairID                 int64          `json:"flair_id"`
	FlairName               FlairName      `json:"flair_name"`
	FlairShortname          FlairShortname `json:"flair_shortname"`
	Friend                  bool           `json:"friend"`
	Helmet                  Helmet         `json:"helmet"`
	Incidents               int64          `json:"incidents"`
	Interval                int64          `json:"interval"`
	LapsComplete            int64          `json:"laps_complete"`
	LapsLead                int64          `json:"laps_lead"`
	LeagueAggPoints         int64          `json:"league_agg_points"`
	LeaguePoints            int64          `json:"league_points"`
	LicenseChangeOval       int64          `json:"license_change_oval"`
	LicenseChangeRoad       int64          `json:"license_change_road"`
	Livery                  Livery         `json:"livery"`
	MaxPctFuelFill          int64          `json:"max_pct_fuel_fill"`
	NewCpi                  float64        `json:"new_cpi"`
	NewLicenseLevel         int64          `json:"new_license_level"`
	NewSubLevel             int64          `json:"new_sub_level"`
	NewTtrating             int64          `json:"new_ttrating"`
	NewiRating              int64          `json:"newi_rating"`
	OldCpi                  float64        `json:"old_cpi"`
	OldLicenseLevel         int64          `json:"old_license_level"`
	OldSubLevel             int64          `json:"old_sub_level"`
	OldTtrating             int64          `json:"old_ttrating"`
	OldiRating              int64          `json:"oldi_rating"`
	OptLapsComplete         int64          `json:"opt_laps_complete"`
	Position                int64          `json:"position"`
	QualLapTime             int64          `json:"qual_lap_time"`
	ReasonOut               ReasonOut      `json:"reason_out"`
	ReasonOutID             int64          `json:"reason_out_id"`
	StartingPosition        int64          `json:"starting_position"`
	StartingPositionInClass int64          `json:"starting_position_in_class"`
	Suit                    Suit           `json:"suit"`
	Watched                 bool           `json:"watched"`
	WeightPenaltyKg         int64          `json:"weight_penalty_kg"`
}

type Helmet struct {
	Pattern    int64  `json:"pattern"`
	Color1     string `json:"color1"`
	Color2     string `json:"color2"`
	Color3     string `json:"color3"`
	FaceType   int64  `json:"face_type"`
	HelmetType int64  `json:"helmet_type"`
}

type Livery struct {
	CarID        int64       `json:"car_id"`
	Pattern      int64       `json:"pattern"`
	Color1       string      `json:"color1"`
	Color2       string      `json:"color2"`
	Color3       Color3      `json:"color3"`
	NumberFont   int64       `json:"number_font"`
	NumberColor1 Color3      `json:"number_color1"`
	NumberColor2 string      `json:"number_color2"`
	NumberColor3 string      `json:"number_color3"`
	NumberSlant  int64       `json:"number_slant"`
	Sponsor1     int64       `json:"sponsor1"`
	Sponsor2     int64       `json:"sponsor2"`
	CarNumber    string      `json:"car_number"`
	WheelColor   interface{} `json:"wheel_color"`
	RimType      int64       `json:"rim_type"`
}

type Suit struct {
	Pattern int64  `json:"pattern"`
	Color1  string `json:"color1"`
	Color2  string `json:"color2"`
	Color3  string `json:"color3"`
}

type SessionSplit struct {
	SubsessionID         int64 `json:"subsession_id"`
	EventStrengthOfField int64 `json:"event_strength_of_field"`
}

type Track struct {
	Category   string `json:"category"`
	CategoryID int64  `json:"category_id"`
	ConfigName string `json:"config_name"`
	TrackID    int64  `json:"track_id"`
	TrackName  string `json:"track_name"`
}

type TrackState struct {
	LeaveMarbles   bool  `json:"leave_marbles"`
	PracticeRubber int64 `json:"practice_rubber"`
	QualifyRubber  int64 `json:"qualify_rubber"`
	RaceRubber     int64 `json:"race_rubber"`
	WarmupRubber   int64 `json:"warmup_rubber"`
}

type Weather struct {
	AllowFog                      bool      `json:"allow_fog"`
	Fog                           int64     `json:"fog"`
	PrecipMm2HrBeforeFinalSession int64     `json:"precip_mm2hr_before_final_session"`
	PrecipMmFinalSession          int64     `json:"precip_mm_final_session"`
	PrecipOption                  int64     `json:"precip_option"`
	PrecipTimePct                 int64     `json:"precip_time_pct"`
	RelHumidity                   int64     `json:"rel_humidity"`
	SimulatedStartTime            time.Time `json:"simulated_start_time"`
	Skies                         int64     `json:"skies"`
	TempUnits                     int64     `json:"temp_units"`
	TempValue                     int64     `json:"temp_value"`
	TimeOfDay                     int64     `json:"time_of_day"`
	TrackWater                    int64     `json:"track_water"`
	Type                          int64     `json:"type"`
	Version                       int64     `json:"version"`
	WeatherVarInitial             int64     `json:"weather_var_initial"`
	WeatherVarOngoing             int64     `json:"weather_var_ongoing"`
	WindDir                       int64     `json:"wind_dir"`
	WindUnits                     int64     `json:"wind_units"`
	WindValue                     int64     `json:"wind_value"`
}

type Name string

const (
	HostedAllCarsClass Name = "Hosted All Cars Class"
)

type ShortName string

const (
	HostedAllCars ShortName = "Hosted All Cars"
)

type CarName string

const (
	Ferrari296GT3            CarName = "Ferrari 296 GT3"
	LamborghiniHuracánGT3EVO CarName = "Lamborghini Huracán GT3 EVO"
)

type CountryCode string

const (
	It CountryCode = "IT"
)

type FlairName string

const (
	Italy FlairName = "Italy"
)

type FlairShortname string

const (
	Ita FlairShortname = "ITA"
)

type Color3 string

const (
	Ffffff    Color3 = "ffffff"
	The0000Ff Color3 = "0000ff"
	The03Bbbd Color3 = "03bbbd"
)

type ReasonOut string

const (
	Running ReasonOut = "Running"
)
