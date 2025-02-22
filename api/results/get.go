package results

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/markphelps/optional"
)

type ResultsGetParams struct {
	SubsessionId    int            `url:"subsession_id,omitempty"`
	IncludeLicenses *optional.Bool `url:"include_licenses,omitempty"`
}

type ResultsGetResponse struct {
	SubsessionId            int   `json:"subsession_id"`
	AssociatedSubsessionIds []int `json:"associated_subsession_ids"`
	CanProtest              bool  `json:"can_protest"`
	CarClasses              []struct {
		CarClassId      int    `json:"car_class_id"`
		ShortName       string `json:"short_name"`
		Name            string `json:"name"`
		StrengthOfField int    `json:"strength_of_field"`
		NumEntries      int    `json:"num_entries"`
		CarsInClass     []struct {
			CarId int `json:"car_id"`
		} `json:"cars_in_class"`
	} `json:"car_classes"`
	CautionType           int    `json:"caution_type"`
	CooldownMinutes       int    `json:"cooldown_minutes"`
	CornersPerLap         int    `json:"corners_per_lap"`
	DamageModel           int    `json:"damage_model"`
	DriverChangeParam1    int    `json:"driver_change_param1"`
	DriverChangeParam2    int    `json:"driver_change_param2"`
	DriverChangeRule      int    `json:"driver_change_rule"`
	DriverChanges         bool   `json:"driver_changes"`
	EndTime               string `json:"end_time"`
	EventAverageLap       int    `json:"event_average_lap"`
	EventBestLapTime      int    `json:"event_best_lap_time"`
	EventLapsComplete     int    `json:"event_laps_complete"`
	EventStrengthOfField  int    `json:"event_strength_of_field"`
	EventType             int    `json:"event_type"`
	EventTypeName         string `json:"event_type_name"`
	HeatInfoId            int    `json:"heat_info_id"`
	HostId                int    `json:"host_id"`
	LeagueId              int    `json:"league_id"`
	LeagueName            string `json:"league_name"`
	LeagueSeasonId        int    `json:"league_season_id"`
	LicenseCategory       string `json:"license_category"`
	LicenseCategoryId     int    `json:"license_category_id"`
	LimitMinutes          int    `json:"limit_minutes"`
	MaxTeamDrivers        int    `json:"max_team_drivers"`
	MaxWeeks              int    `json:"max_weeks"`
	MinTeamDrivers        int    `json:"min_team_drivers"`
	NumCautionLaps        int    `json:"num_caution_laps"`
	NumCautions           int    `json:"num_cautions"`
	NumDrivers            int    `json:"num_drivers"`
	NumLapsForQualAverage int    `json:"num_laps_for_qual_average"`
	NumLapsForSoloAverage int    `json:"num_laps_for_solo_average"`
	NumLeadChanges        int    `json:"num_lead_changes"`
	OfficialSession       bool   `json:"official_session"`
	PointsType            string `json:"points_type"`
	PrivateSessionId      int    `json:"private_session_id"`
	RaceWeekNum           int    `json:"race_week_num"`
	RestrictResults       bool   `json:"restrict_results"`
	ResultsRestricted     bool   `json:"results_restricted"`
	SeasonId              int    `json:"season_id"`
	SeasonName            string `json:"season_name"`
	SeasonQuarter         int    `json:"season_quarter"`
	SeasonShortName       string `json:"season_short_name"`
	SeasonYear            int    `json:"season_year"`
	SeriesId              int    `json:"series_id"`
	SeriesName            string `json:"series_name"`
	SeriesShortName       string `json:"series_short_name"`
	SessionId             int    `json:"session_id"`
	SessionName           string `json:"session_name"`
	SessionResults        []struct {
		SimsessionNumber   int    `json:"simsession_number"`
		SimsessionName     string `json:"simsession_name"`
		SimsessionType     int    `json:"simsession_type"`
		SimsessionTypeName string `json:"simsession_type_name"`
		SimsessionSubtype  int    `json:"simsession_subtype"`
		Results            []struct {
			CustId                int    `json:"cust_id"`
			DisplayName           string `json:"display_name"`
			AggregateChampPoints  int    `json:"aggregate_champ_points"`
			Ai                    bool   `json:"ai"`
			AverageLap            int    `json:"average_lap"`
			BestLapNum            int    `json:"best_lap_num"`
			BestLapTime           int    `json:"best_lap_time"`
			BestNlapsNum          int    `json:"best_nlaps_num"`
			BestNlapsTime         int    `json:"best_nlaps_time"`
			BestQualLapAt         string `json:"best_qual_lap_at"`
			BestQualLapNum        int    `json:"best_qual_lap_num"`
			BestQualLapTime       int    `json:"best_qual_lap_time"`
			CarClassId            int    `json:"car_class_id"`
			CarClassName          string `json:"car_class_name"`
			CarClassShortName     string `json:"car_class_short_name"`
			CarId                 int    `json:"car_id"`
			CarName               string `json:"car_name"`
			ChampPoints           int    `json:"champ_points"`
			ClassInterval         int    `json:"class_interval"`
			ClubId                int    `json:"club_id"`
			ClubName              string `json:"club_name"`
			ClubPoints            int    `json:"club_points"`
			ClubShortname         string `json:"club_shortname"`
			CountryCode           string `json:"country_code"`
			Division              int    `json:"division"`
			DropRace              bool   `json:"drop_race"`
			FinishPosition        int    `json:"finish_position"`
			FinishPositionInClass int    `json:"finish_position_in_class"`
			Friend                bool   `json:"friend"`
			Helmet                struct {
				Pattern    int    `json:"pattern"`
				Color1     string `json:"color1"`
				Color2     string `json:"color2"`
				Color3     string `json:"color3"`
				FaceType   int    `json:"face_type"`
				HelmetType int    `json:"helmet_type"`
			} `json:"helmet"`
			Incidents         int `json:"incidents"`
			Interval          int `json:"interval"`
			LapsComplete      int `json:"laps_complete"`
			LapsLead          int `json:"laps_lead"`
			LeagueAggPoints   int `json:"league_agg_points"`
			LeaguePoints      int `json:"league_points"`
			LicenseChangeOval int `json:"license_change_oval"`
			LicenseChangeRoad int `json:"license_change_road"`
			Livery            struct {
				CarId        int         `json:"car_id"`
				Pattern      int         `json:"pattern"`
				Color1       string      `json:"color1"`
				Color2       string      `json:"color2"`
				Color3       string      `json:"color3"`
				NumberFont   int         `json:"number_font"`
				NumberColor1 string      `json:"number_color1"`
				NumberColor2 string      `json:"number_color2"`
				NumberColor3 string      `json:"number_color3"`
				NumberSlant  int         `json:"number_slant"`
				Sponsor1     int         `json:"sponsor1"`
				Sponsor2     int         `json:"sponsor2"`
				CarNumber    string      `json:"car_number"`
				WheelColor   interface{} `json:"wheel_color"`
				RimType      int         `json:"rim_type"`
			} `json:"livery"`
			MaxPctFuelFill          int     `json:"max_pct_fuel_fill"`
			Multiplier              int     `json:"multiplier"`
			NewCpi                  float64 `json:"new_cpi"`
			NewLicenseLevel         int     `json:"new_license_level"`
			NewSubLevel             int     `json:"new_sub_level"`
			NewTtrating             int     `json:"new_ttrating"`
			NewiRating              int     `json:"newi_rating"`
			OldCpi                  float64 `json:"old_cpi"`
			OldLicenseLevel         int     `json:"old_license_level"`
			OldSubLevel             int     `json:"old_sub_level"`
			OldTtrating             int     `json:"old_ttrating"`
			OldiRating              int     `json:"oldi_rating"`
			OptLapsComplete         int     `json:"opt_laps_complete"`
			Position                int     `json:"position"`
			QualLapTime             int     `json:"qual_lap_time"`
			ReasonOut               string  `json:"reason_out"`
			ReasonOutId             int     `json:"reason_out_id"`
			StartingPosition        int     `json:"starting_position"`
			StartingPositionInClass int     `json:"starting_position_in_class"`
			Suit                    struct {
				Pattern int    `json:"pattern"`
				Color1  string `json:"color1"`
				Color2  string `json:"color2"`
				Color3  string `json:"color3"`
			} `json:"suit"`
			Watched         bool `json:"watched"`
			WeightPenaltyKg int  `json:"weight_penalty_kg"`
		} `json:"results"`
	} `json:"session_results"`
	SessionSplits []struct {
		SubsessionId         int `json:"subsession_id"`
		EventStrengthOfField int `json:"event_strength_of_field"`
	} `json:"session_splits"`
	SpecialEventType int    `json:"special_event_type"`
	StartTime        string `json:"start_time"`
	Track            struct {
		Category   string `json:"category"`
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
	Weather struct {
		AllowFog                      bool   `json:"allow_fog"`
		Fog                           int    `json:"fog"`
		PrecipMm2hrBeforeFinalSession int    `json:"precip_mm2hr_before_final_session"`
		PrecipMmFinalSession          int    `json:"precip_mm_final_session"`
		PrecipOption                  int    `json:"precip_option"`
		PrecipTimePct                 int    `json:"precip_time_pct"`
		RelHumidity                   int    `json:"rel_humidity"`
		SimulatedStartTime            string `json:"simulated_start_time"`
		Skies                         int    `json:"skies"`
		TempUnits                     int    `json:"temp_units"`
		TempValue                     int    `json:"temp_value"`
		TimeOfDay                     int    `json:"time_of_day"`
		TrackWater                    int    `json:"track_water"`
		Type                          int    `json:"type"`
		Version                       int    `json:"version"`
		WeatherVarInitial             int    `json:"weather_var_initial"`
		WeatherVarOngoing             int    `json:"weather_var_ongoing"`
		WindDir                       int    `json:"wind_dir"`
		WindUnits                     int    `json:"wind_units"`
		WindValue                     int    `json:"wind_value"`
	} `json:"weather"`
}

// Get the results of a subsession, if authorized to view them. series_logo image paths are relative to https://images-static.iracing.com/img/logos/series/
func (api *ResultsApi) GetResults(params ResultsGetParams) (*ResultsGetResponse, error) {
	paramsString, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	url := "https://members-ng.iracing.com/data/results/get?" + paramsString.Encode()

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &ResultsGetResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
