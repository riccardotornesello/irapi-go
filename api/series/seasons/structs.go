package seasons

import (
	"github.com/riccardotornesello/irapi-go/pkg/types"
)

type SeriesSeasonsParams struct {
	IncludeSeries *bool `url:"include_series,omitempty,comma"`
	SeasonYear    *int  `url:"season_year,omitempty,comma"`
	SeasonQuarter *int  `url:"season_quarter,omitempty,comma"`
}

type SeriesSeasonsResponse []SeriesSeasonsResponseElement

type SeriesSeasonsResponseElement struct {
	SeasonID                   int64              `json:"season_id"`
	SeasonName                 string             `json:"season_name"`
	Active                     bool               `json:"active"`
	AllowedSeasonMembers       interface{}        `json:"allowed_season_members"`
	CarClassIDS                []int64            `json:"car_class_ids"`
	CarSwitching               bool               `json:"car_switching"`
	CarTypes                   []CarType          `json:"car_types"`
	CautionLapsDoNotCount      bool               `json:"caution_laps_do_not_count"`
	Complete                   bool               `json:"complete"`
	ConnectionBlackFlag        bool               `json:"connection_black_flag"`
	ConsecCautionWithinNlaps   int64              `json:"consec_caution_within_nlaps"`
	ConsecCautionsSingleFile   bool               `json:"consec_cautions_single_file"`
	CrossLicense               bool               `json:"cross_license"`
	DistributedMatchmaking     bool               `json:"distributed_matchmaking"`
	DriverChangeRule           int64              `json:"driver_change_rule"`
	DriverChanges              bool               `json:"driver_changes"`
	Drops                      int64              `json:"drops"`
	EnablePitlaneCollisions    bool               `json:"enable_pitlane_collisions"`
	FixedSetup                 bool               `json:"fixed_setup"`
	GreenWhiteCheckeredLimit   int64              `json:"green_white_checkered_limit"`
	GridByClass                bool               `json:"grid_by_class"`
	HardcoreLevel              int64              `json:"hardcore_level"`
	HasSupersessions           bool               `json:"has_supersessions"`
	IgnoreLicenseForPractice   bool               `json:"ignore_license_for_practice"`
	IncidentLimit              int64              `json:"incident_limit"`
	IncidentWarnMode           int64              `json:"incident_warn_mode"`
	IncidentWarnParam1         int64              `json:"incident_warn_param1"`
	IncidentWarnParam2         int64              `json:"incident_warn_param2"`
	IsHeatRacing               bool               `json:"is_heat_racing"`
	LicenseGroup               int64              `json:"license_group"`
	LicenseGroupTypes          []LicenseGroupType `json:"license_group_types"`
	LuckyDog                   bool               `json:"lucky_dog"`
	MaxTeamDrivers             int64              `json:"max_team_drivers"`
	MaxWeeks                   int64              `json:"max_weeks"`
	MinTeamDrivers             int64              `json:"min_team_drivers"`
	Multiclass                 bool               `json:"multiclass"`
	MustUseDiffTireTypesInRace bool               `json:"must_use_diff_tire_types_in_race"`
	NextRaceSession            interface{}        `json:"next_race_session"`
	NumFastTows                int64              `json:"num_fast_tows"`
	NumOptLaps                 int64              `json:"num_opt_laps"`
	Official                   bool               `json:"official"`
	OpDuration                 int64              `json:"op_duration"`
	OpenPracticeSessionTypeID  int64              `json:"open_practice_session_type_id"`
	QualifierMustStartRace     bool               `json:"qualifier_must_start_race"`
	RaceWeek                   int64              `json:"race_week"`
	RaceWeekToMakeDivisions    int64              `json:"race_week_to_make_divisions"`
	RegUserCount               int64              `json:"reg_user_count"`
	RegionCompetition          bool               `json:"region_competition"`
	RestrictByMember           bool               `json:"restrict_by_member"`
	RestrictToCar              bool               `json:"restrict_to_car"`
	RestrictViewing            bool               `json:"restrict_viewing"`
	ScheduleDescription        string             `json:"schedule_description"`
	Schedules                  []Schedule         `json:"schedules"`
	ScoreAsCarclassid          *int64             `json:"score_as_carclassid,omitempty"`
	SeasonQuarter              int64              `json:"season_quarter"`
	SeasonShortName            string             `json:"season_short_name"`
	SeasonYear                 int64              `json:"season_year"`
	SendToOpenPractice         bool               `json:"send_to_open_practice"`
	SeriesID                   int64              `json:"series_id"`
	ShortParadeLap             bool               `json:"short_parade_lap"`
	StartDate                  types.DateTime     `json:"start_date"`
	StartOnQualTire            bool               `json:"start_on_qual_tire"`
	StartZone                  bool               `json:"start_zone"`
	TrackTypes                 []TrackType        `json:"track_types"`
	UnsportConductRuleMode     int64              `json:"unsport_conduct_rule_mode"`
	RookieSeason               *string            `json:"rookie_season,omitempty"`
	RegOpenMinutes             *int64             `json:"reg_open_minutes,omitempty"`
	HeatSesInfo                *HeatSesInfo       `json:"heat_ses_info,omitempty"`
}

type CarType struct {
	CarType string `json:"car_type"`
}

type HeatSesInfo struct {
	ConsolationDeltaMaxFieldSize         int64          `json:"consolation_delta_max_field_size"`
	ConsolationDeltaSessionLaps          int64          `json:"consolation_delta_session_laps"`
	ConsolationDeltaSessionLengthMinutes int64          `json:"consolation_delta_session_length_minutes"`
	ConsolationFirstMaxFieldSize         int64          `json:"consolation_first_max_field_size"`
	ConsolationFirstSessionLaps          int64          `json:"consolation_first_session_laps"`
	ConsolationFirstSessionLengthMinutes int64          `json:"consolation_first_session_length_minutes"`
	ConsolationNumPositionToInvert       int64          `json:"consolation_num_position_to_invert"`
	ConsolationNumToConsolation          int64          `json:"consolation_num_to_consolation"`
	ConsolationNumToMain                 int64          `json:"consolation_num_to_main"`
	ConsolationRunAlways                 bool           `json:"consolation_run_always"`
	ConsolationScoresChampPoints         bool           `json:"consolation_scores_champ_points"`
	Created                              types.DateTime `json:"created"`
	CustID                               int64          `json:"cust_id"`
	Description                          string         `json:"description"`
	HeatCautionType                      int64          `json:"heat_caution_type"`
	HeatInfoID                           int64          `json:"heat_info_id"`
	HeatInfoName                         string         `json:"heat_info_name"`
	HeatLaps                             int64          `json:"heat_laps"`
	HeatLengthMinutes                    int64          `json:"heat_length_minutes"`
	HeatMaxFieldSize                     int64          `json:"heat_max_field_size"`
	HeatNumFromEachToMain                int64          `json:"heat_num_from_each_to_main"`
	HeatNumPositionToInvert              int64          `json:"heat_num_position_to_invert"`
	HeatScoresChampPoints                bool           `json:"heat_scores_champ_points"`
	HeatSessionMinutesEstimate           int64          `json:"heat_session_minutes_estimate"`
	Hidden                               bool           `json:"hidden"`
	MainLaps                             int64          `json:"main_laps"`
	MainLengthMinutes                    int64          `json:"main_length_minutes"`
	MainMaxFieldSize                     int64          `json:"main_max_field_size"`
	MainNumPositionToInvert              int64          `json:"main_num_position_to_invert"`
	MaxEntrants                          int64          `json:"max_entrants"`
	OpenPractice                         bool           `json:"open_practice"`
	PreMainPracticeLengthMinutes         int64          `json:"pre_main_practice_length_minutes"`
	PreQualNumToMain                     int64          `json:"pre_qual_num_to_main"`
	PreQualPracticeLengthMinutes         int64          `json:"pre_qual_practice_length_minutes"`
	QualCautionType                      int64          `json:"qual_caution_type"`
	QualLaps                             int64          `json:"qual_laps"`
	QualLengthMinutes                    int64          `json:"qual_length_minutes"`
	QualNumToMain                        int64          `json:"qual_num_to_main"`
	QualOpenDelaySeconds                 int64          `json:"qual_open_delay_seconds"`
	QualScoresChampPoints                bool           `json:"qual_scores_champ_points"`
	QualScoring                          int64          `json:"qual_scoring"`
	QualStyle                            int64          `json:"qual_style"`
	RaceStyle                            int64          `json:"race_style"`
}

type LicenseGroupType struct {
	LicenseGroupType int64 `json:"license_group_type"`
}

type Schedule struct {
	SeasonID                int64                `json:"season_id"`
	RaceWeekNum             int64                `json:"race_week_num"`
	CarRestrictions         []CarRestriction     `json:"car_restrictions"`
	Category                string               `json:"category"`
	CategoryID              int64                `json:"category_id"`
	EnablePitlaneCollisions bool                 `json:"enable_pitlane_collisions"`
	FullCourseCautions      bool                 `json:"full_course_cautions"`
	PracticeLength          *int64               `json:"practice_length,omitempty"`
	QualAttached            bool                 `json:"qual_attached"`
	QualifyLaps             *int64               `json:"qualify_laps,omitempty"`
	QualifyLength           *int64               `json:"qualify_length,omitempty"`
	RaceLapLimit            *int64               `json:"race_lap_limit"`
	RaceTimeDescriptors     []RaceTimeDescriptor `json:"race_time_descriptors"`
	RaceTimeLimit           *int64               `json:"race_time_limit"`
	RaceWeekCarClassIDS     []int64              `json:"race_week_car_class_ids"`
	RaceWeekCars            []RaceWeekCar        `json:"race_week_cars"`
	RestartType             string               `json:"restart_type"`
	ScheduleName            string               `json:"schedule_name"`
	SeasonName              string               `json:"season_name"`
	SeriesID                int64                `json:"series_id"`
	SeriesName              string               `json:"series_name"`
	ShortParadeLap          bool                 `json:"short_parade_lap"`
	SpecialEventType        interface{}          `json:"special_event_type"`
	StartDate               string               `json:"start_date"`
	StartType               string               `json:"start_type"`
	StartZone               bool                 `json:"start_zone"`
	Track                   Track                `json:"track"`
	TrackState              TrackState           `json:"track_state"`
	WarmupLength            *int64               `json:"warmup_length,omitempty"`
	Weather                 Weather              `json:"weather"`
	WeekEndTime             types.DateTime       `json:"week_end_time"`
}

type CarRestriction struct {
	CarID           int64   `json:"car_id"`
	MaxDryTireSets  int64   `json:"max_dry_tire_sets"`
	MaxPctFuelFill  int64   `json:"max_pct_fuel_fill"`
	PowerAdjustPct  float64 `json:"power_adjust_pct"`
	RaceSetupID     *int64  `json:"race_setup_id,omitempty"`
	WeightPenaltyKg int64   `json:"weight_penalty_kg"`
	QualSetupID     *int64  `json:"qual_setup_id,omitempty"`
}

type RaceTimeDescriptor struct {
	Repeating        bool             `json:"repeating"`
	SessionMinutes   int64            `json:"session_minutes"`
	SessionTimes     []types.DateTime `json:"session_times,omitempty"`
	SuperSession     bool             `json:"super_session"`
	DayOffset        []int64          `json:"day_offset,omitempty"`
	FirstSessionTime *string          `json:"first_session_time,omitempty"`
	RepeatMinutes    *int64           `json:"repeat_minutes,omitempty"`
	StartDate        *string          `json:"start_date,omitempty"`
}

type RaceWeekCar struct {
	CarID              int64  `json:"car_id"`
	CarName            string `json:"car_name"`
	CarNameAbbreviated string `json:"car_name_abbreviated"`
}

type Track struct {
	Category   string  `json:"category"`
	CategoryID int64   `json:"category_id"`
	TrackID    int64   `json:"track_id"`
	TrackName  string  `json:"track_name"`
	ConfigName *string `json:"config_name,omitempty"`
}

type TrackState struct {
	LeaveMarbles   bool   `json:"leave_marbles"`
	PracticeRubber *int64 `json:"practice_rubber,omitempty"`
}

type Weather struct {
	AllowFog                *bool            `json:"allow_fog,omitempty"`
	ForecastOptions         *ForecastOptions `json:"forecast_options,omitempty"`
	PrecipOption            *int64           `json:"precip_option,omitempty"`
	RelHumidity             int64            `json:"rel_humidity"`
	SimulatedStartTime      types.DateTime   `json:"simulated_start_time"`
	SimulatedTimeMultiplier int64            `json:"simulated_time_multiplier"`
	SimulatedTimeOffsets    []int64          `json:"simulated_time_offsets"`
	Skies                   int64            `json:"skies"`
	TempUnits               int64            `json:"temp_units"`
	TempValue               int64            `json:"temp_value"`
	TimeOfDay               int64            `json:"time_of_day"`
	TrackWater              *int64           `json:"track_water,omitempty"`
	Version                 int64            `json:"version"`
	WeatherSummary          *WeatherSummary  `json:"weather_summary,omitempty"`
	WeatherURL              *string          `json:"weather_url,omitempty"`
	WindDir                 int64            `json:"wind_dir"`
	WindUnits               int64            `json:"wind_units"`
	WindValue               int64            `json:"wind_value"`
	Fog                     *int64           `json:"fog,omitempty"`
	Type                    *int64           `json:"type,omitempty"`
	WeatherVarInitial       *int64           `json:"weather_var_initial,omitempty"`
	WeatherVarOngoing       *int64           `json:"weather_var_ongoing,omitempty"`
}

type ForecastOptions struct {
	AllowFog      bool  `json:"allow_fog"`
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
	MaxPrecipRate     float64 `json:"max_precip_rate"`
	MaxPrecipRateDesc string  `json:"max_precip_rate_desc"`
	PrecipChance      int64   `json:"precip_chance"`
	SkiesHigh         int64   `json:"skies_high"`
	SkiesLow          int64   `json:"skies_low"`
	TempHigh          float64 `json:"temp_high"`
	TempLow           float64 `json:"temp_low"`
	TempUnits         int64   `json:"temp_units"`
	WindDir           int64   `json:"wind_dir"`
	WindHigh          float64 `json:"wind_high"`
	WindLow           float64 `json:"wind_low"`
	WindUnits         int64   `json:"wind_units"`
}

type TrackType struct {
	TrackType string `json:"track_type"`
}
