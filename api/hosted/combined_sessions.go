package hosted

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/markphelps/optional"
)

type HostedCombinedSessionsParams struct {
	PackageId *optional.Int `url:"package_id,omitempty"`
}

type HostedCombinedSessionsResponse struct {
	Sessions []struct {
		CountByCarId                      map[string]int `json:"count_by_car_id"`
		CountByCarClassId                 map[string]int `json:"count_by_car_class_id"`
		NumDrivers                        int            `json:"num_drivers"`
		NumSpotters                       int            `json:"num_spotters"`
		NumSpectators                     int            `json:"num_spectators"`
		NumBroadcasters                   int            `json:"num_broadcasters"`
		AvailableReservedBroadcasterSlots int            `json:"available_reserved_broadcaster_slots"`
		NumSpectatorSlots                 int            `json:"num_spectator_slots"`
		AvailableSpectatorSlots           int            `json:"available_spectator_slots"`
		CanBroadcast                      bool           `json:"can_broadcast"`
		CanWatch                          bool           `json:"can_watch"`
		CanSpot                           bool           `json:"can_spot"`
		Elig                              struct {
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
		DriverChanges              bool   `json:"driver_changes"`
		RestrictViewing            bool   `json:"restrict_viewing"`
		MaxUsers                   int    `json:"max_users"`
		PrivateSessionId           int    `json:"private_session_id"`
		SessionId                  int    `json:"session_id"`
		SubsessionId               int    `json:"subsession_id"`
		PasswordProtected          bool   `json:"password_protected"`
		SessionName                string `json:"session_name"`
		SessionDesc                string `json:"session_desc"`
		OpenRegExpires             string `json:"open_reg_expires"`
		LaunchAt                   string `json:"launch_at"`
		FullCourseCautions         bool   `json:"full_course_cautions"`
		NumFastTows                int    `json:"num_fast_tows"`
		RollingStarts              bool   `json:"rolling_starts"`
		Restarts                   int    `json:"restarts"`
		MulticlassType             int    `json:"multiclass_type"`
		PitsInUse                  int    `json:"pits_in_use"`
		CarsLeft                   int    `json:"cars_left"`
		MaxDrivers                 int    `json:"max_drivers"`
		HardcoreLevel              int    `json:"hardcore_level"`
		PracticeLength             int    `json:"practice_length"`
		LoneQualify                bool   `json:"lone_qualify"`
		QualifyLaps                int    `json:"qualify_laps"`
		QualifyLength              int    `json:"qualify_length"`
		WarmupLength               int    `json:"warmup_length"`
		RaceLaps                   int    `json:"race_laps"`
		RaceLength                 int    `json:"race_length"`
		TimeLimit                  int    `json:"time_limit"`
		RestrictResults            bool   `json:"restrict_results"`
		IncidentLimit              int    `json:"incident_limit"`
		IncidentWarnMode           int    `json:"incident_warn_mode"`
		IncidentWarnParam1         int    `json:"incident_warn_param1"`
		IncidentWarnParam2         int    `json:"incident_warn_param2"`
		UnsportConductRuleMode     int    `json:"unsport_conduct_rule_mode"`
		LuckyDog                   bool   `json:"lucky_dog"`
		MinTeamDrivers             int    `json:"min_team_drivers"`
		MaxTeamDrivers             int    `json:"max_team_drivers"`
		QualifierMustStartRace     bool   `json:"qualifier_must_start_race"`
		DriverChangeRule           int    `json:"driver_change_rule"`
		FixedSetup                 bool   `json:"fixed_setup"`
		EntryCount                 int    `json:"entry_count"`
		LeagueId                   int    `json:"league_id"`
		LeagueSeasonId             int    `json:"league_season_id"`
		SessionType                int    `json:"session_type"`
		OrderId                    int    `json:"order_id"`
		MinLicenseLevel            int    `json:"min_license_level"`
		MaxLicenseLevel            int    `json:"max_license_level"`
		Status                     int    `json:"status"`
		PaceCarId                  *int   `json:"pace_car_id"`
		PaceCarClassId             *int   `json:"pace_car_class_id"`
		NumOptLaps                 int    `json:"num_opt_laps"`
		DamageModel                int    `json:"damage_model"`
		DoNotPaintCars             bool   `json:"do_not_paint_cars"`
		GreenWhiteCheckeredLimit   int    `json:"green_white_checkered_limit"`
		DoNotCountCautionLaps      bool   `json:"do_not_count_caution_laps"`
		ConsecCautionsSingleFile   bool   `json:"consec_cautions_single_file"`
		NoLapperWaveArounds        bool   `json:"no_lapper_wave_arounds"`
		ShortParadeLap             bool   `json:"short_parade_lap"`
		StartOnQualTire            bool   `json:"start_on_qual_tire"`
		TelemetryRestriction       int    `json:"telemetry_restriction"`
		TelemetryForceToDisk       int    `json:"telemetry_force_to_disk"`
		MaxAiDrivers               int    `json:"max_ai_drivers"`
		AiAvoidPlayers             bool   `json:"ai_avoid_players"`
		MustUseDiffTireTypesInRace bool   `json:"must_use_diff_tire_types_in_race"`
		StartZone                  bool   `json:"start_zone"`
		EnablePitlaneCollisions    bool   `json:"enable_pitlane_collisions"`
		DisallowVirtualMirror      bool   `json:"disallow_virtual_mirror"`
		MaxVisorTearoffs           int    `json:"max_visor_tearoffs"`
		CategoryId                 int    `json:"category_id"`
		Category                   string `json:"category"`
		SessionFull                bool   `json:"session_full"`
		Host                       struct {
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
		Track struct {
			CategoryId int    `json:"category_id"`
			ConfigName string `json:"config_name"`
			TrackId    int    `json:"track_id"`
			TrackName  string `json:"track_name"`
		} `json:"track"`
		Weather struct {
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
		Farm struct {
			FarmId      int    `json:"farm_id"`
			DisplayName string `json:"display_name"`
			ImagePath   string `json:"image_path"`
			Displayed   bool   `json:"displayed"`
		} `json:"farm"`
		Admins []struct {
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
		AllowedClubs   []interface{} `json:"allowed_clubs"`
		AllowedTeams   []interface{} `json:"allowed_teams"`
		AllowedLeagues []interface{} `json:"allowed_leagues"`
		Cars           []struct {
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
		RegisteredTeams []int `json:"registered_teams"`
		CarTypes        []struct {
			CarType string `json:"car_type"`
		} `json:"car_types"`
		TrackTypes []struct {
			TrackType string `json:"track_type"`
		} `json:"track_types"`
		LicenseGroupTypes []struct {
			LicenseGroupType int `json:"license_group_type"`
		} `json:"license_group_types"`
		EventTypes []struct {
			EventType int `json:"event_type"`
		} `json:"event_types"`
		SessionTypes []struct {
			SessionType int `json:"session_type"`
		} `json:"session_types"`
		CanJoin        bool          `json:"can_join"`
		SessAdmin      bool          `json:"sess_admin"`
		Friends        []interface{} `json:"friends"`
		Watched        []interface{} `json:"watched"`
		EndTime        string        `json:"end_time"`
		TeamEntryCount int           `json:"team_entry_count"`
		IsHeatRacing   bool          `json:"is_heat_racing"`
		Populated      bool          `json:"populated"`
		Broadcaster    bool          `json:"broadcaster"`
		MinIr          int           `json:"min_ir"`
		MaxIr          int           `json:"max_ir"`
		AltAssetId     int           `json:"alt_asset_id"`
		AiMinSkill     int           `json:"ai_min_skill"`
		AiMaxSkill     int           `json:"ai_max_skill"`
		AiRosterName   string        `json:"ai_roster_name"`
		HeatSesInfo    struct {
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
	Sequence   int  `json:"sequence"`
	Success    bool `json:"success"`
}

// Sessions that can be joined as a driver or spectator, and also includes non-league pending sessions for the user.
func (api *HostedApi) GetHostedCombinedSessions(params HostedCombinedSessionsParams) (*HostedCombinedSessionsResponse, error) {
	paramsString, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	url := "https://members-ng.iracing.com/data/hosted/combined_sessions?" + paramsString.Encode()

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &HostedCombinedSessionsResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
