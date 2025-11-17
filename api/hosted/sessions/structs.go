package sessions

import (
	"time"
)

type HostedSessionsResponse struct {
	Subscribed bool      `json:"subscribed"`
	Sessions   []Session `json:"sessions"`
	Success    bool      `json:"success"`
}

type Session struct {
	AdaptiveAIDifficulty       *int64             `json:"adaptive_ai_difficulty,omitempty"`
	AdaptiveAIEnabled          bool               `json:"adaptive_ai_enabled"`
	Admins                     []Host             `json:"admins"`
	AIAvoidPlayers             bool               `json:"ai_avoid_players"`
	AIRosterName               *string            `json:"ai_roster_name,omitempty"`
	AllowedLeagues             []interface{}      `json:"allowed_leagues"`
	AllowedTeams               []interface{}      `json:"allowed_teams"`
	CarTypes                   []CarType          `json:"car_types"`
	Cars                       []Car              `json:"cars"`
	CarsLeft                   int64              `json:"cars_left"`
	Category                   string             `json:"category"`
	CategoryID                 int64              `json:"category_id"`
	ConnectionBlackFlag        bool               `json:"connection_black_flag"`
	ConsecCautionWithinNlaps   int64              `json:"consec_caution_within_nlaps"`
	ConsecCautionsSingleFile   bool               `json:"consec_cautions_single_file"`
	CountByCarClassID          map[string]int64   `json:"count_by_car_class_id"`
	CountByCarID               map[string]int64   `json:"count_by_car_id"`
	DamageModel                int64              `json:"damage_model"`
	DisallowVirtualMirror      bool               `json:"disallow_virtual_mirror"`
	DoNotCountCautionLaps      bool               `json:"do_not_count_caution_laps"`
	DoNotPaintCars             bool               `json:"do_not_paint_cars"`
	DriverChangeRule           int64              `json:"driver_change_rule"`
	DriverChanges              bool               `json:"driver_changes"`
	Elig                       Elig               `json:"elig"`
	EnablePitlaneCollisions    bool               `json:"enable_pitlane_collisions"`
	EntryCount                 int64              `json:"entry_count"`
	EventTypes                 []EventType        `json:"event_types"`
	Farm                       Farm               `json:"farm"`
	FixedSetup                 bool               `json:"fixed_setup"`
	FullCourseCautions         bool               `json:"full_course_cautions"`
	GreenWhiteCheckeredLimit   int64              `json:"green_white_checkered_limit"`
	HardcoreLevel              int64              `json:"hardcore_level"`
	Host                       Host               `json:"host"`
	IncidentLimit              int64              `json:"incident_limit"`
	IncidentWarnMode           int64              `json:"incident_warn_mode"`
	IncidentWarnParam1         int64              `json:"incident_warn_param1"`
	IncidentWarnParam2         int64              `json:"incident_warn_param2"`
	LaunchAt                   time.Time          `json:"launch_at"`
	LeagueID                   int64              `json:"league_id"`
	LeagueSeasonID             int64              `json:"league_season_id"`
	LicenseGroupTypes          []LicenseGroupType `json:"license_group_types"`
	LoneQualify                bool               `json:"lone_qualify"`
	LuckyDog                   bool               `json:"lucky_dog"`
	MaxAIDrivers               int64              `json:"max_ai_drivers"`
	MaxDrivers                 int64              `json:"max_drivers"`
	MaxIR                      int64              `json:"max_ir"`
	MaxLicenseLevel            int64              `json:"max_license_level"`
	MaxTeamDrivers             int64              `json:"max_team_drivers"`
	MaxVisorTearoffs           int64              `json:"max_visor_tearoffs"`
	MinIR                      int64              `json:"min_ir"`
	MinLicenseLevel            int64              `json:"min_license_level"`
	MinTeamDrivers             int64              `json:"min_team_drivers"`
	MulticlassType             int64              `json:"multiclass_type"`
	MustUseDiffTireTypesInRace bool               `json:"must_use_diff_tire_types_in_race"`
	NoLapperWaveArounds        bool               `json:"no_lapper_wave_arounds"`
	NumFastTows                int64              `json:"num_fast_tows"`
	NumOptLaps                 int64              `json:"num_opt_laps"`
	OpenRegExpires             time.Time          `json:"open_reg_expires"`
	OrderID                    int64              `json:"order_id"`
	PaceCarClassID             *int64             `json:"pace_car_class_id"`
	PaceCarID                  *int64             `json:"pace_car_id"`
	PasswordProtected          bool               `json:"password_protected"`
	PitsInUse                  int64              `json:"pits_in_use"`
	PracticeLength             int64              `json:"practice_length"`
	PrivateSessionID           int64              `json:"private_session_id"`
	QualifierMustStartRace     bool               `json:"qualifier_must_start_race"`
	QualifyLaps                int64              `json:"qualify_laps"`
	QualifyLength              int64              `json:"qualify_length"`
	RaceLaps                   int64              `json:"race_laps"`
	RaceLength                 int64              `json:"race_length"`
	Restarts                   int64              `json:"restarts"`
	RestrictResults            bool               `json:"restrict_results"`
	RestrictViewing            bool               `json:"restrict_viewing"`
	RollingStarts              bool               `json:"rolling_starts"`
	SessionFull                bool               `json:"session_full"`
	SessionID                  int64              `json:"session_id"`
	SessionName                string             `json:"session_name"`
	SessionType                int64              `json:"session_type"`
	SessionTypes               []SessionType      `json:"session_types"`
	ShortParadeLap             bool               `json:"short_parade_lap"`
	StartOnQualTire            bool               `json:"start_on_qual_tire"`
	StartZone                  bool               `json:"start_zone"`
	Status                     int64              `json:"status"`
	SubsessionID               int64              `json:"subsession_id"`
	TeamEntryCount             int64              `json:"team_entry_count"`
	TelemetryForceToDisk       int64              `json:"telemetry_force_to_disk"`
	TelemetryRestriction       int64              `json:"telemetry_restriction"`
	TimeLimit                  int64              `json:"time_limit"`
	Track                      Track              `json:"track"`
	TrackState                 TrackState         `json:"track_state"`
	TrackTypes                 []TrackType        `json:"track_types"`
	UnsportConductRuleMode     int64              `json:"unsport_conduct_rule_mode"`
	WarmupLength               int64              `json:"warmup_length"`
	Weather                    Weather            `json:"weather"`
	SessionDesc                *string            `json:"session_desc,omitempty"`
	HeatSesInfo                *HeatSesInfo       `json:"heat_ses_info,omitempty"`
	AIMaxSkill                 *int64             `json:"ai_max_skill,omitempty"`
	AIMinSkill                 *int64             `json:"ai_min_skill,omitempty"`
	AltAssetID                 *int64             `json:"alt_asset_id,omitempty"`
}

type Host struct {
	CustID      int64  `json:"cust_id"`
	DisplayName string `json:"display_name"`
	Helmet      Helmet `json:"helmet"`
}

type Helmet struct {
	Pattern    int64  `json:"pattern"`
	Color1     string `json:"color1"`
	Color2     string `json:"color2"`
	Color3     string `json:"color3"`
	FaceType   int64  `json:"face_type"`
	HelmetType int64  `json:"helmet_type"`
}

type CarType struct {
	CarType string `json:"car_type"`
}

type Car struct {
	CarID             int64   `json:"car_id"`
	CarName           string  `json:"car_name"`
	CarClassID        int64   `json:"car_class_id"`
	CarClassName      string  `json:"car_class_name"`
	MaxPctFuelFill    int64   `json:"max_pct_fuel_fill"`
	WeightPenaltyKg   int64   `json:"weight_penalty_kg"`
	PowerAdjustPct    int64   `json:"power_adjust_pct"`
	MaxDryTireSets    int64   `json:"max_dry_tire_sets"`
	PackageID         int64   `json:"package_id"`
	QualSetupID       *int64  `json:"qual_setup_id,omitempty"`
	QualSetupFilename *string `json:"qual_setup_filename,omitempty"`
	RaceSetupID       *int64  `json:"race_setup_id,omitempty"`
	RaceSetupFilename *string `json:"race_setup_filename,omitempty"`
}

type Elig struct {
	SessionFull     bool    `json:"session_full"`
	CanSpot         bool    `json:"can_spot"`
	CanWatch        bool    `json:"can_watch"`
	CanDrive        bool    `json:"can_drive"`
	HasSessPassword bool    `json:"has_sess_password"`
	NeedsPurchase   bool    `json:"needs_purchase"`
	OwnCar          bool    `json:"own_car"`
	OwnTrack        bool    `json:"own_track"`
	PurchaseSkus    []int64 `json:"purchase_skus"`
	Registered      bool    `json:"registered"`
}

type EventType struct {
	EventType int64 `json:"event_type"`
}

type Farm struct {
	FarmID      int64  `json:"farm_id"`
	DisplayName string `json:"display_name"`
	ImagePath   string `json:"image_path"`
	Displayed   bool   `json:"displayed"`
}

type HeatSesInfo struct {
	ConsolationDeltaMaxFieldSize         int64     `json:"consolation_delta_max_field_size"`
	ConsolationDeltaSessionLaps          int64     `json:"consolation_delta_session_laps"`
	ConsolationDeltaSessionLengthMinutes int64     `json:"consolation_delta_session_length_minutes"`
	ConsolationFirstMaxFieldSize         int64     `json:"consolation_first_max_field_size"`
	ConsolationFirstSessionLaps          int64     `json:"consolation_first_session_laps"`
	ConsolationFirstSessionLengthMinutes int64     `json:"consolation_first_session_length_minutes"`
	ConsolationNumPositionToInvert       int64     `json:"consolation_num_position_to_invert"`
	ConsolationNumToConsolation          int64     `json:"consolation_num_to_consolation"`
	ConsolationNumToMain                 int64     `json:"consolation_num_to_main"`
	ConsolationRunAlways                 bool      `json:"consolation_run_always"`
	ConsolationScoresChampPoints         bool      `json:"consolation_scores_champ_points"`
	Created                              time.Time `json:"created"`
	CustID                               int64     `json:"cust_id"`
	HeatCautionType                      int64     `json:"heat_caution_type"`
	HeatInfoID                           int64     `json:"heat_info_id"`
	HeatInfoName                         string    `json:"heat_info_name"`
	HeatLaps                             int64     `json:"heat_laps"`
	HeatLengthMinutes                    int64     `json:"heat_length_minutes"`
	HeatMaxFieldSize                     int64     `json:"heat_max_field_size"`
	HeatNumFromEachToMain                int64     `json:"heat_num_from_each_to_main"`
	HeatNumPositionToInvert              int64     `json:"heat_num_position_to_invert"`
	HeatScoresChampPoints                bool      `json:"heat_scores_champ_points"`
	HeatSessionMinutesEstimate           int64     `json:"heat_session_minutes_estimate"`
	Hidden                               bool      `json:"hidden"`
	MainLaps                             int64     `json:"main_laps"`
	MainLengthMinutes                    int64     `json:"main_length_minutes"`
	MainMaxFieldSize                     int64     `json:"main_max_field_size"`
	MainNumPositionToInvert              int64     `json:"main_num_position_to_invert"`
	MaxEntrants                          int64     `json:"max_entrants"`
	OpenPractice                         bool      `json:"open_practice"`
	PreMainPracticeLengthMinutes         int64     `json:"pre_main_practice_length_minutes"`
	PreQualNumToMain                     int64     `json:"pre_qual_num_to_main"`
	PreQualPracticeLengthMinutes         int64     `json:"pre_qual_practice_length_minutes"`
	QualCautionType                      int64     `json:"qual_caution_type"`
	QualLaps                             int64     `json:"qual_laps"`
	QualLengthMinutes                    int64     `json:"qual_length_minutes"`
	QualNumToMain                        int64     `json:"qual_num_to_main"`
	QualOpenDelaySeconds                 int64     `json:"qual_open_delay_seconds"`
	QualScoresChampPoints                bool      `json:"qual_scores_champ_points"`
	QualScoring                          int64     `json:"qual_scoring"`
	QualStyle                            int64     `json:"qual_style"`
	RaceStyle                            int64     `json:"race_style"`
}

type LicenseGroupType struct {
	LicenseGroupType int64 `json:"license_group_type"`
}

type SessionType struct {
	SessionType int64 `json:"session_type"`
}

type Track struct {
	CategoryID int64   `json:"category_id"`
	TrackID    int64   `json:"track_id"`
	TrackName  string  `json:"track_name"`
	ConfigName *string `json:"config_name,omitempty"`
}

type TrackState struct {
	LeaveMarbles   bool  `json:"leave_marbles"`
	PracticeRubber int64 `json:"practice_rubber"`
	QualifyRubber  int64 `json:"qualify_rubber"`
	RaceRubber     int64 `json:"race_rubber"`
	WarmupRubber   int64 `json:"warmup_rubber"`
}

type TrackType struct {
	TrackType string `json:"track_type"`
}

type Weather struct {
	AllowFog                bool             `json:"allow_fog"`
	Fog                     *int64           `json:"fog,omitempty"`
	PrecipOption            int64            `json:"precip_option"`
	RelHumidity             int64            `json:"rel_humidity"`
	SimulatedStartTime      time.Time        `json:"simulated_start_time"`
	SimulatedTimeMultiplier int64            `json:"simulated_time_multiplier"`
	SimulatedTimeOffsets    []int64          `json:"simulated_time_offsets"`
	Skies                   int64            `json:"skies"`
	TempUnits               int64            `json:"temp_units"`
	TempValue               int64            `json:"temp_value"`
	TimeOfDay               int64            `json:"time_of_day"`
	TrackWater              int64            `json:"track_water"`
	Type                    *int64           `json:"type,omitempty"`
	Version                 int64            `json:"version"`
	WindDir                 int64            `json:"wind_dir"`
	WindUnits               int64            `json:"wind_units"`
	WindValue               int64            `json:"wind_value"`
	ForecastOptions         *ForecastOptions `json:"forecast_options,omitempty"`
	WeatherSummary          *WeatherSummary  `json:"weather_summary,omitempty"`
	WeatherURL              *string          `json:"weather_url,omitempty"`
}

type ForecastOptions struct {
	ForecastType  int64 `json:"forecast_type"`
	Precipitation int64 `json:"precipitation"`
	Skies         int64 `json:"skies"`
	StopPrecip    int64 `json:"stop_precip"`
	Temperature   int64 `json:"temperature"`
	WeatherSeed   int64 `json:"weather_seed"`
	WindDir       int64 `json:"wind_dir"`
	WindSpeed     int64 `json:"wind_speed"`
}

type WeatherSummary struct {
	MaxPrecipRate     *float64 `json:"max_precip_rate,omitempty"`
	MaxPrecipRateDesc string   `json:"max_precip_rate_desc"`
	PrecipChance      float64  `json:"precip_chance"`
	SkiesHigh         *int64   `json:"skies_high,omitempty"`
	SkiesLow          *int64   `json:"skies_low,omitempty"`
	TempHigh          *float64 `json:"temp_high,omitempty"`
	TempLow           *float64 `json:"temp_low,omitempty"`
	TempUnits         *int64   `json:"temp_units,omitempty"`
	WindDir           *int64   `json:"wind_dir,omitempty"`
	WindHigh          *float64 `json:"wind_high,omitempty"`
	WindLow           *float64 `json:"wind_low,omitempty"`
	WindUnits         *int64   `json:"wind_units,omitempty"`
}
