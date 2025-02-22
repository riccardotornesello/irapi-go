package series

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/markphelps/optional"
)

type SeriesSeasonsParams struct {
	IncludeSeries *optional.Bool `url:"include_series,omitempty"`
}

type SeriesSeasonsResponse []struct {
	AllowedSeasonMembers *map[string]struct {
		CustId      int    `json:"cust_id"`
		DisplayName string `json:"display_name"`
		SeasonId    int    `json:"season_id"`
	} `json:"allowed_season_members"`
	SeasonId     int    `json:"season_id"`
	SeasonName   string `json:"season_name"`
	Active       bool   `json:"active"`
	CarClassIds  []int  `json:"car_class_ids"`
	CarSwitching bool   `json:"car_switching"`
	CarTypes     []struct {
		CarType string `json:"car_type"`
	} `json:"car_types"`
	CautionLapsDoNotCount    bool `json:"caution_laps_do_not_count"`
	Complete                 bool `json:"complete"`
	CrossLicense             bool `json:"cross_license"`
	DistributedMatchmaking   bool `json:"distributed_matchmaking"`
	DriverChangeRule         int  `json:"driver_change_rule"`
	DriverChanges            bool `json:"driver_changes"`
	Drops                    int  `json:"drops"`
	EnablePitlaneCollisions  bool `json:"enable_pitlane_collisions"`
	FixedSetup               bool `json:"fixed_setup"`
	GreenWhiteCheckeredLimit int  `json:"green_white_checkered_limit"`
	GridByClass              bool `json:"grid_by_class"`
	HardcoreLevel            int  `json:"hardcore_level"`
	HasSupersessions         bool `json:"has_supersessions"`
	IgnoreLicenseForPractice bool `json:"ignore_license_for_practice"`
	IncidentLimit            int  `json:"incident_limit"`
	IncidentWarnMode         int  `json:"incident_warn_mode"`
	IncidentWarnParam1       int  `json:"incident_warn_param1"`
	IncidentWarnParam2       int  `json:"incident_warn_param2"`
	IsHeatRacing             bool `json:"is_heat_racing"`
	LicenseGroup             int  `json:"license_group"`
	LicenseGroupTypes        []struct {
		LicenseGroupType int `json:"license_group_type"`
	} `json:"license_group_types"`
	LuckyDog                   bool        `json:"lucky_dog"`
	MaxTeamDrivers             int         `json:"max_team_drivers"`
	MaxWeeks                   int         `json:"max_weeks"`
	MinTeamDrivers             int         `json:"min_team_drivers"`
	Multiclass                 bool        `json:"multiclass"`
	MustUseDiffTireTypesInRace bool        `json:"must_use_diff_tire_types_in_race"`
	NextRaceSession            interface{} `json:"next_race_session"`
	NumFastTows                int         `json:"num_fast_tows"`
	NumOptLaps                 int         `json:"num_opt_laps"`
	Official                   bool        `json:"official"`
	OpDuration                 int         `json:"op_duration"`
	OpenPracticeSessionTypeId  int         `json:"open_practice_session_type_id"`
	QualifierMustStartRace     bool        `json:"qualifier_must_start_race"`
	RaceWeek                   int         `json:"race_week"`
	RaceWeekToMakeDivisions    int         `json:"race_week_to_make_divisions"`
	RegUserCount               int         `json:"reg_user_count"`
	RegionCompetition          bool        `json:"region_competition"`
	RestrictByMember           bool        `json:"restrict_by_member"`
	RestrictToCar              bool        `json:"restrict_to_car"`
	RestrictViewing            bool        `json:"restrict_viewing"`
	ScheduleDescription        string      `json:"schedule_description"`
	Schedules                  []struct {
		SeasonId        int `json:"season_id"`
		RaceWeekNum     int `json:"race_week_num"`
		CarRestrictions []struct {
			CarId           int     `json:"car_id"`
			MaxDryTireSets  int     `json:"max_dry_tire_sets"`
			MaxPctFuelFill  int     `json:"max_pct_fuel_fill"`
			PowerAdjustPct  float64 `json:"power_adjust_pct"`
			WeightPenaltyKg int     `json:"weight_penalty_kg"`
			RaceSetupId     int     `json:"race_setup_id"`
			QualSetupId     int     `json:"qual_setup_id"`
		} `json:"car_restrictions"`
		Category                string `json:"category"`
		CategoryId              int    `json:"category_id"`
		EnablePitlaneCollisions bool   `json:"enable_pitlane_collisions"`
		FullCourseCautions      bool   `json:"full_course_cautions"`
		PracticeLength          int    `json:"practice_length"`
		QualAttached            bool   `json:"qual_attached"`
		QualifyLaps             int    `json:"qualify_laps"`
		QualifyLength           int    `json:"qualify_length"`
		RaceLapLimit            *int   `json:"race_lap_limit"`
		RaceTimeDescriptors     []struct {
			DayOffset        []int    `json:"day_offset"`
			FirstSessionTime string   `json:"first_session_time"`
			RepeatMinutes    int      `json:"repeat_minutes"`
			Repeating        bool     `json:"repeating"`
			SessionMinutes   int      `json:"session_minutes"`
			StartDate        string   `json:"start_date"`
			SuperSession     bool     `json:"super_session"`
			SessionTimes     []string `json:"session_times"`
		} `json:"race_time_descriptors"`
		RaceTimeLimit *int `json:"race_time_limit"`
		RaceWeekCars  []struct {
			CarId              int    `json:"car_id"`
			CarName            string `json:"car_name"`
			CarNameAbbreviated string `json:"car_name_abbreviated"`
		} `json:"race_week_cars"`
		RestartType      string      `json:"restart_type"`
		ScheduleName     string      `json:"schedule_name"`
		SeasonName       string      `json:"season_name"`
		SeriesId         int         `json:"series_id"`
		SeriesName       string      `json:"series_name"`
		ShortParadeLap   bool        `json:"short_parade_lap"`
		SpecialEventType interface{} `json:"special_event_type"`
		StartDate        string      `json:"start_date"`
		StartType        string      `json:"start_type"`
		StartZone        bool        `json:"start_zone"`
		Track            struct {
			Category   string `json:"category"`
			CategoryId int    `json:"category_id"`
			ConfigName string `json:"config_name"`
			TrackId    int    `json:"track_id"`
			TrackName  string `json:"track_name"`
		} `json:"track"`
		TrackState struct {
			LeaveMarbles   bool `json:"leave_marbles"`
			PracticeRubber int  `json:"practice_rubber"`
			RaceRubber     int  `json:"race_rubber"`
		} `json:"track_state"`
		WarmupLength int `json:"warmup_length"`
		Weather      struct {
			AllowFog        bool `json:"allow_fog"`
			ForecastOptions struct {
				ForecastType  int `json:"forecast_type"`
				Precipitation int `json:"precipitation"`
				Skies         int `json:"skies"`
				StopPrecip    int `json:"stop_precip"`
				Temperature   int `json:"temperature"`
				WeatherSeed   int `json:"weather_seed"`
				WindDir       int `json:"wind_dir"`
				WindSpeed     int `json:"wind_speed"`
			} `json:"forecast_options"`
			PrecipOption            int    `json:"precip_option"`
			RelHumidity             int    `json:"rel_humidity"`
			SimulatedStartTime      string `json:"simulated_start_time"`
			SimulatedTimeMultiplier int    `json:"simulated_time_multiplier"`
			SimulatedTimeOffsets    []int  `json:"simulated_time_offsets"`
			Skies                   int    `json:"skies"`
			TempUnits               int    `json:"temp_units"`
			TempValue               int    `json:"temp_value"`
			TimeOfDay               int    `json:"time_of_day"`
			TrackWater              int    `json:"track_water"`
			Version                 int    `json:"version"`
			WeatherSummary          struct {
				MaxPrecipRate     float64 `json:"max_precip_rate"`
				MaxPrecipRateDesc string  `json:"max_precip_rate_desc"`
				PrecipChance      float64 `json:"precip_chance"`
				SkiesHigh         int     `json:"skies_high"`
				SkiesLow          int     `json:"skies_low"`
				TempHigh          float64 `json:"temp_high"`
				TempLow           float64 `json:"temp_low"`
				TempUnits         int     `json:"temp_units"`
				WindDir           int     `json:"wind_dir"`
				WindHigh          float64 `json:"wind_high"`
				WindLow           float64 `json:"wind_low"`
				WindUnits         int     `json:"wind_units"`
			} `json:"weather_summary"`
			WeatherUrl        string `json:"weather_url"`
			WindDir           int    `json:"wind_dir"`
			WindUnits         int    `json:"wind_units"`
			WindValue         int    `json:"wind_value"`
			Fog               int    `json:"fog"`
			Type              int    `json:"type"`
			WeatherVarInitial int    `json:"weather_var_initial"`
			WeatherVarOngoing int    `json:"weather_var_ongoing"`
		} `json:"weather"`
		WeekEndTime string `json:"week_end_time"`
	} `json:"schedules"`
	SeasonQuarter      int    `json:"season_quarter"`
	SeasonShortName    string `json:"season_short_name"`
	SeasonYear         int    `json:"season_year"`
	SendToOpenPractice bool   `json:"send_to_open_practice"`
	SeriesId           int    `json:"series_id"`
	ShortParadeLap     bool   `json:"short_parade_lap"`
	StartDate          string `json:"start_date"`
	StartOnQualTire    bool   `json:"start_on_qual_tire"`
	StartZone          bool   `json:"start_zone"`
	TrackTypes         []struct {
		TrackType string `json:"track_type"`
	} `json:"track_types"`
	UnsportConductRuleMode int `json:"unsport_conduct_rule_mode"`
	HeatSesInfo            struct {
		ConsolationDeltaMaxFieldSize         int    `json:"consolation_delta_max_field_size"`
		ConsolationDeltaSessionLaps          int    `json:"consolation_delta_session_laps"`
		ConsolationDeltaSessionLengthMinutes int    `json:"consolation_delta_session_length_minutes"`
		ConsolationFirstMaxFieldSize         int    `json:"consolation_first_max_field_size"`
		ConsolationFirstSessionLaps          int    `json:"consolation_first_session_laps"`
		ConsolationFirstSessionLengthMinutes int    `json:"consolation_first_session_length_minutes"`
		ConsolationNumPositionToInvert       int    `json:"consolation_num_position_to_invert"`
		ConsolationNumToConsolation          int    `json:"consolation_num_to_consolation"`
		ConsolationNumToMain                 int    `json:"consolation_num_to_main"`
		ConsolationRunAlways                 bool   `json:"consolation_run_always"`
		ConsolationScoresChampPoints         bool   `json:"consolation_scores_champ_points"`
		Created                              string `json:"created"`
		CustId                               int    `json:"cust_id"`
		Description                          string `json:"description"`
		HeatCautionType                      int    `json:"heat_caution_type"`
		HeatInfoId                           int    `json:"heat_info_id"`
		HeatInfoName                         string `json:"heat_info_name"`
		HeatLaps                             int    `json:"heat_laps"`
		HeatLengthMinutes                    int    `json:"heat_length_minutes"`
		HeatMaxFieldSize                     int    `json:"heat_max_field_size"`
		HeatNumFromEachToMain                int    `json:"heat_num_from_each_to_main"`
		HeatNumPositionToInvert              int    `json:"heat_num_position_to_invert"`
		HeatScoresChampPoints                bool   `json:"heat_scores_champ_points"`
		HeatSessionMinutesEstimate           int    `json:"heat_session_minutes_estimate"`
		Hidden                               bool   `json:"hidden"`
		MainLaps                             int    `json:"main_laps"`
		MainLengthMinutes                    int    `json:"main_length_minutes"`
		MainMaxFieldSize                     int    `json:"main_max_field_size"`
		MainNumPositionToInvert              int    `json:"main_num_position_to_invert"`
		MaxEntrants                          int    `json:"max_entrants"`
		OpenPractice                         bool   `json:"open_practice"`
		PreMainPracticeLengthMinutes         int    `json:"pre_main_practice_length_minutes"`
		PreQualNumToMain                     int    `json:"pre_qual_num_to_main"`
		PreQualPracticeLengthMinutes         int    `json:"pre_qual_practice_length_minutes"`
		QualCautionType                      int    `json:"qual_caution_type"`
		QualLaps                             int    `json:"qual_laps"`
		QualLengthMinutes                    int    `json:"qual_length_minutes"`
		QualNumToMain                        int    `json:"qual_num_to_main"`
		QualOpenDelaySeconds                 int    `json:"qual_open_delay_seconds"`
		QualScoresChampPoints                bool   `json:"qual_scores_champ_points"`
		QualScoring                          int    `json:"qual_scoring"`
		QualStyle                            int    `json:"qual_style"`
		RaceStyle                            int    `json:"race_style"`
	} `json:"heat_ses_info"`
	RacePoints     int    `json:"race_points"`
	RookieSeason   string `json:"rookie_season"`
	RegOpenMinutes int    `json:"reg_open_minutes"`
}

func (api *SeriesApi) GetSeriesSeasons(params SeriesSeasonsParams) (*SeriesSeasonsResponse, error) {
	paramsString, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	url := "/data/series/seasons?" + paramsString.Encode()

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &SeriesSeasonsResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
