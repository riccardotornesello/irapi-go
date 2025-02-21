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
	Subscribed bool `json:"subscribed"`
	Sequence   int  `json:"sequence"`
	Sessions   []struct {
		NumDrivers                        int  `json:"num_drivers"`
		NumSpotters                       int  `json:"num_spotters"`
		NumSpectators                     int  `json:"num_spectators"`
		NumBroadcasters                   int  `json:"num_broadcasters"`
		AvailableReservedBroadcasterSlots int  `json:"available_reserved_broadcaster_slots"`
		NumSpectatorSlots                 int  `json:"num_spectator_slots"`
		AvailableSpectatorSlots           int  `json:"available_spectator_slots"`
		CanBroadcast                      bool `json:"can_broadcast"`
		CanWatch                          bool `json:"can_watch"`
		CanSpot                           bool `json:"can_spot"`
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
			AllowFog                bool   `json:"allow_fog"`
			Fog                     int    `json:"fog"`
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
			Type                    int    `json:"type"`
			Version                 int    `json:"version"`
			WindDir                 int    `json:"wind_dir"`
			WindUnits               int    `json:"wind_units"`
			WindValue               int    `json:"wind_value"`
			ForecastOptions         struct {
				ForecastType  int `json:"forecast_type"`
				Precipitation int `json:"precipitation"`
				Skies         int `json:"skies"`
				StopPrecip    int `json:"stop_precip"`
				Temperature   int `json:"temperature"`
				WeatherSeed   int `json:"weather_seed"`
				WindDir       int `json:"wind_dir"`
				WindSpeed     int `json:"wind_speed"`
			} `json:"forecast_options"`
			WeatherSummary struct {
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
			WeatherUrl string `json:"weather_url"`
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
		CountByCarId struct {
			Field147 int `json:"147"`
			Field132 int `json:"132"`
			Field133 int `json:"133"`
			Field156 int `json:"156"`
			Field169 int `json:"169"`
			Field173 int `json:"173"`
			Field176 int `json:"176"`
			Field184 int `json:"184"`
			Field185 int `json:"185"`
			Field188 int `json:"188"`
			Field194 int `json:"194"`
			Field172 int `json:"172"`
			Field43  int `json:"43"`
			Field67  int `json:"67"`
			Field72  int `json:"72"`
			Field73  int `json:"73"`
			Field106 int `json:"106"`
			Field135 int `json:"135"`
			Field152 int `json:"152"`
			Field157 int `json:"157"`
			Field160 int `json:"160"`
			Field178 int `json:"178"`
			Field189 int `json:"189"`
			Field103 int `json:"103"`
			Field96  int `json:"96"`
			Field143 int `json:"143"`
			Field13  int `json:"13"`
			Field119 int `json:"119"`
			Field195 int `json:"195"`
			Field164 int `json:"164"`
			Field131 int `json:"131"`
			Field56  int `json:"56"`
			Field110 int `json:"110"`
			Field111 int `json:"111"`
			Field114 int `json:"114"`
			Field115 int `json:"115"`
			Field116 int `json:"116"`
			Field123 int `json:"123"`
			Field139 int `json:"139"`
			Field140 int `json:"140"`
			Field141 int `json:"141"`
			Field155 int `json:"155"`
			Field167 int `json:"167"`
		} `json:"count_by_car_id"`
		CountByCarClassId struct {
			Field0     int `json:"0"`
			Field2708  int `json:"2708"`
			Field4072  int `json:"4072"`
			Field59285 int `json:"59285"`
			Field59287 int `json:"59287"`
			Field59288 int `json:"59288"`
			Field59568 int `json:"59568"`
			Field15    int `json:"15"`
			Field55293 int `json:"55293"`
			Field60653 int `json:"60653"`
			Field60654 int `json:"60654"`
			Field73998 int `json:"73998"`
			Field73999 int `json:"73999"`
			Field74721 int `json:"74721"`
			Field2705  int `json:"2705"`
			Field2742  int `json:"2742"`
			Field4000  int `json:"4000"`
			Field4025  int `json:"4025"`
			Field4036  int `json:"4036"`
			Field4046  int `json:"4046"`
			Field4060  int `json:"4060"`
			Field4061  int `json:"4061"`
			Field4066  int `json:"4066"`
		} `json:"count_by_car_class_id"`
		CarTypes []struct {
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
		CanJoin         bool          `json:"can_join"`
		SessAdmin       bool          `json:"sess_admin"`
		Friends         []interface{} `json:"friends"`
		Watched         []interface{} `json:"watched"`
		EndTime         string        `json:"end_time"`
		TeamEntryCount  int           `json:"team_entry_count"`
		IsHeatRacing    bool          `json:"is_heat_racing"`
		Populated       bool          `json:"populated"`
		Broadcaster     bool          `json:"broadcaster"`
		MinIr           int           `json:"min_ir"`
		MaxIr           int           `json:"max_ir"`
		SessionDesc     string        `json:"session_desc"`
		AiMinSkill      int           `json:"ai_min_skill"`
		AiMaxSkill      int           `json:"ai_max_skill"`
		AiRosterName    string        `json:"ai_roster_name"`
		RegisteredTeams []int         `json:"registered_teams"`
	} `json:"sessions"`
	Success bool `json:"success"`
}

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
