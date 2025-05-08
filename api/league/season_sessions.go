package league

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
)

type LeagueSeasonSessionsParams struct {
	LeagueId    int   `url:"league_id"`
	SeasonId    int   `url:"season_id"`
	ResultsOnly *bool `url:"results_only,omitempty"` // If true include only sessions for which results are available.
}

type LeagueSeasonSessionsResponse struct {
	Sessions []struct {
		Cars []struct {
			CarId        int    `json:"car_id"`
			CarName      string `json:"car_name"`
			CarClassId   int    `json:"car_class_id"`
			CarClassName string `json:"car_class_name"`
		} `json:"cars"`
		DriverChanges     bool        `json:"driver_changes"`
		EntryCount        int         `json:"entry_count"`
		HasResults        bool        `json:"has_results"`
		LaunchAt          string      `json:"launch_at"`
		LeagueId          int         `json:"league_id"`
		LeagueSeasonId    int         `json:"league_season_id"`
		LoneQualify       bool        `json:"lone_qualify"`
		PaceCarClassId    interface{} `json:"pace_car_class_id"`
		PaceCarId         interface{} `json:"pace_car_id"`
		PasswordProtected bool        `json:"password_protected"`
		PracticeLength    int         `json:"practice_length"`
		PrivateSessionId  int         `json:"private_session_id"`
		QualifyLaps       int         `json:"qualify_laps"`
		QualifyLength     int         `json:"qualify_length"`
		RaceLaps          int         `json:"race_laps"`
		RaceLength        int         `json:"race_length"`
		SessionId         int         `json:"session_id"`
		Status            int         `json:"status"`
		SubsessionId      int         `json:"subsession_id"`
		TeamEntryCount    int         `json:"team_entry_count"`
		TimeLimit         int         `json:"time_limit"`
		Track             struct {
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
			AllowFog       bool `json:"allow_fog"`
			Fog            int  `json:"fog"`
			PrecipOption   int  `json:"precip_option"`
			RelHumidity    int  `json:"rel_humidity"`
			Skies          int  `json:"skies"`
			TempUnits      int  `json:"temp_units"`
			TempValue      int  `json:"temp_value"`
			TrackWater     int  `json:"track_water"`
			Type           int  `json:"type"`
			Version        int  `json:"version"`
			WeatherSummary struct {
				MaxPrecipRateDesc string `json:"max_precip_rate_desc"`
				PrecipChance      int    `json:"precip_chance"`
			} `json:"weather_summary"`
			WeatherVarInitial int `json:"weather_var_initial"`
			WeatherVarOngoing int `json:"weather_var_ongoing"`
			WindDir           int `json:"wind_dir"`
			WindUnits         int `json:"wind_units"`
			WindValue         int `json:"wind_value"`
		} `json:"weather"`
		WinnerId   int    `json:"winner_id"`
		WinnerName string `json:"winner_name"`
	} `json:"sessions"`
	Success  bool `json:"success"`
	SeasonId int  `json:"season_id"`
	LeagueId int  `json:"league_id"`
}

func (api *LeagueApi) GetLeagueSeasonSessions(params LeagueSeasonSessionsParams) (*LeagueSeasonSessionsResponse, error) {
	paramsString, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	url := "/data/league/season_sessions?" + paramsString.Encode()

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &LeagueSeasonSessionsResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
