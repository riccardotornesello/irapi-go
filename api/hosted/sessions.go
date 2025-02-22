package hosted

import (
	"encoding/json"
)

type HostedSessionsResponse struct {
	Sessions []struct {
		CountByCarId      map[string]int `json:"count_by_car_id"`
		CountByCarClassId map[string]int `json:"count_by_car_class_id"`
		Admins            []struct {
			CustId      int    `json:"cust_id"`
			DisplayName string `json:"display_name"`
			Helmet      struct {
				Pattern    int    `json:"pattern"`
				Color1     string `json:"color1"`
				Color2     string `json:"color2"`
				Color3     string `json:"color3"`
				FaceType   int    `json:"face_type"`
				HelmetType int    `json:"helmet_type"`
			} `json:"helmet"`
		} `json:"admins"`
		AiAvoidPlayers bool          `json:"ai_avoid_players"`
		AllowedClubs   []interface{} `json:"allowed_clubs"`
		AllowedLeagues []interface{} `json:"allowed_leagues"`
		AllowedTeams   []interface{} `json:"allowed_teams"`
		CarTypes       []struct {
			CarType string `json:"car_type"`
		} `json:"car_types"`
		Cars []struct {
			CarId             int     `json:"car_id"`
			CarName           string  `json:"car_name"`
			CarClassId        int     `json:"car_class_id"`
			CarClassName      string  `json:"car_class_name"`
			MaxPctFuelFill    int     `json:"max_pct_fuel_fill"`
			WeightPenaltyKg   int     `json:"weight_penalty_kg"`
			PowerAdjustPct    float64 `json:"power_adjust_pct"`
			MaxDryTireSets    int     `json:"max_dry_tire_sets"`
			PackageId         int     `json:"package_id"`
			QualSetupId       int     `json:"qual_setup_id"`
			QualSetupFilename string  `json:"qual_setup_filename"`
			RaceSetupId       int     `json:"race_setup_id"`
			RaceSetupFilename string  `json:"race_setup_filename"`
		} `json:"cars"`
		CarsLeft                 int    `json:"cars_left"`
		Category                 string `json:"category"`
		CategoryId               int    `json:"category_id"`
		ConsecCautionsSingleFile bool   `json:"consec_cautions_single_file"`
		DamageModel              int    `json:"damage_model"`
		DisallowVirtualMirror    bool   `json:"disallow_virtual_mirror"`
		DoNotCountCautionLaps    bool   `json:"do_not_count_caution_laps"`
		DoNotPaintCars           bool   `json:"do_not_paint_cars"`
		DriverChangeRule         int    `json:"driver_change_rule"`
		DriverChanges            bool   `json:"driver_changes"`
		Elig                     struct {
			SessionFull     bool  `json:"session_full"`
			CanSpot         bool  `json:"can_spot"`
			CanWatch        bool  `json:"can_watch"`
			CanDrive        bool  `json:"can_drive"`
			HasSessPassword bool  `json:"has_sess_password"`
			NeedsPurchase   bool  `json:"needs_purchase"`
			OwnCar          bool  `json:"own_car"`
			OwnTrack        bool  `json:"own_track"`
			PurchaseSkus    []int `json:"purchase_skus"`
			Registered      bool  `json:"registered"`
		} `json:"elig"`
		EnablePitlaneCollisions bool `json:"enable_pitlane_collisions"`
		EntryCount              int  `json:"entry_count"`
		EventTypes              []struct {
			EventType int `json:"event_type"`
		} `json:"event_types"`
		Farm struct {
			FarmId      int    `json:"farm_id"`
			DisplayName string `json:"display_name"`
			ImagePath   string `json:"image_path"`
			Displayed   bool   `json:"displayed"`
		} `json:"farm"`
		FixedSetup               bool `json:"fixed_setup"`
		FullCourseCautions       bool `json:"full_course_cautions"`
		GreenWhiteCheckeredLimit int  `json:"green_white_checkered_limit"`
		HardcoreLevel            int  `json:"hardcore_level"`
		Host                     struct {
			CustId      int    `json:"cust_id"`
			DisplayName string `json:"display_name"`
			Helmet      struct {
				Pattern    int    `json:"pattern"`
				Color1     string `json:"color1"`
				Color2     string `json:"color2"`
				Color3     string `json:"color3"`
				FaceType   int    `json:"face_type"`
				HelmetType int    `json:"helmet_type"`
			} `json:"helmet"`
		} `json:"host"`
		IncidentLimit      int    `json:"incident_limit"`
		IncidentWarnMode   int    `json:"incident_warn_mode"`
		IncidentWarnParam1 int    `json:"incident_warn_param1"`
		IncidentWarnParam2 int    `json:"incident_warn_param2"`
		LaunchAt           string `json:"launch_at"`
		LeagueId           int    `json:"league_id"`
		LeagueSeasonId     int    `json:"league_season_id"`
		LicenseGroupTypes  []struct {
			LicenseGroupType int `json:"license_group_type"`
		} `json:"license_group_types"`
		LoneQualify                bool   `json:"lone_qualify"`
		LuckyDog                   bool   `json:"lucky_dog"`
		MaxAiDrivers               int    `json:"max_ai_drivers"`
		MaxDrivers                 int    `json:"max_drivers"`
		MaxIr                      int    `json:"max_ir"`
		MaxLicenseLevel            int    `json:"max_license_level"`
		MaxTeamDrivers             int    `json:"max_team_drivers"`
		MaxVisorTearoffs           int    `json:"max_visor_tearoffs"`
		MinIr                      int    `json:"min_ir"`
		MinLicenseLevel            int    `json:"min_license_level"`
		MinTeamDrivers             int    `json:"min_team_drivers"`
		MulticlassType             int    `json:"multiclass_type"`
		MustUseDiffTireTypesInRace bool   `json:"must_use_diff_tire_types_in_race"`
		NoLapperWaveArounds        bool   `json:"no_lapper_wave_arounds"`
		NumFastTows                int    `json:"num_fast_tows"`
		NumOptLaps                 int    `json:"num_opt_laps"`
		OpenRegExpires             string `json:"open_reg_expires"`
		OrderId                    int    `json:"order_id"`
		PaceCarClassId             *int   `json:"pace_car_class_id"`
		PaceCarId                  *int   `json:"pace_car_id"`
		PasswordProtected          bool   `json:"password_protected"`
		PitsInUse                  int    `json:"pits_in_use"`
		PracticeLength             int    `json:"practice_length"`
		PrivateSessionId           int    `json:"private_session_id"`
		QualifierMustStartRace     bool   `json:"qualifier_must_start_race"`
		QualifyLaps                int    `json:"qualify_laps"`
		QualifyLength              int    `json:"qualify_length"`
		RaceLaps                   int    `json:"race_laps"`
		RaceLength                 int    `json:"race_length"`
		RegisteredTeams            []int  `json:"registered_teams"`
		Restarts                   int    `json:"restarts"`
		RestrictResults            bool   `json:"restrict_results"`
		RestrictViewing            bool   `json:"restrict_viewing"`
		RollingStarts              bool   `json:"rolling_starts"`
		SessionDesc                string `json:"session_desc"`
		SessionFull                bool   `json:"session_full"`
		SessionId                  int    `json:"session_id"`
		SessionName                string `json:"session_name"`
		SessionType                int    `json:"session_type"`
		SessionTypes               []struct {
			SessionType int `json:"session_type"`
		} `json:"session_types"`
		ShortParadeLap       bool `json:"short_parade_lap"`
		StartOnQualTire      bool `json:"start_on_qual_tire"`
		StartZone            bool `json:"start_zone"`
		Status               int  `json:"status"`
		SubsessionId         int  `json:"subsession_id"`
		TeamEntryCount       int  `json:"team_entry_count"`
		TelemetryForceToDisk int  `json:"telemetry_force_to_disk"`
		TelemetryRestriction int  `json:"telemetry_restriction"`
		TimeLimit            int  `json:"time_limit"`
		Track                struct {
			CategoryId int    `json:"category_id"`
			ConfigName string `json:"config_name"`
			TrackId    int    `json:"track_id"`
			TrackName  string `json:"track_name"`
		} `json:"track"`
		TrackState struct {
			LeaveMarbles         bool `json:"leave_marbles"`
			PracticeGripCompound int  `json:"practice_grip_compound"`
			PracticeRubber       int  `json:"practice_rubber"`
			QualifyGripCompound  int  `json:"qualify_grip_compound"`
			QualifyRubber        int  `json:"qualify_rubber"`
			RaceGripCompound     int  `json:"race_grip_compound"`
			RaceRubber           int  `json:"race_rubber"`
			WarmupGripCompound   int  `json:"warmup_grip_compound"`
			WarmupRubber         int  `json:"warmup_rubber"`
		} `json:"track_state"`
		TrackTypes []struct {
			TrackType string `json:"track_type"`
		} `json:"track_types"`
		UnsportConductRuleMode int `json:"unsport_conduct_rule_mode"`
		WarmupLength           int `json:"warmup_length"`
		Weather                struct {
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
				PrecipChance      int     `json:"precip_chance"`
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
			WeatherUrl string `json:"weather_url"`
			WindDir    int    `json:"wind_dir"`
			WindUnits  int    `json:"wind_units"`
			WindValue  int    `json:"wind_value"`
			Fog        int    `json:"fog"`
			Type       int    `json:"type"`
		} `json:"weather"`
		AltAssetId   int    `json:"alt_asset_id"`
		AiMaxSkill   int    `json:"ai_max_skill"`
		AiMinSkill   int    `json:"ai_min_skill"`
		AiRosterName string `json:"ai_roster_name"`
		HeatSesInfo  struct {
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
			Description                          string `json:"description"`
		} `json:"heat_ses_info"`
	} `json:"sessions"`
	Subscribed bool `json:"subscribed"`
	Success    bool `json:"success"`
}

// Sessions that can be joined as a driver. Without spectator and non-league pending sessions for the user.
func (api *HostedApi) GetHostedSessions() (*HostedSessionsResponse, error) {
	url := "https://members-ng.iracing.com/data/hosted/sessions"

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &HostedSessionsResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
