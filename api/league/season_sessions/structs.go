package season_sessions

import (
	"github.com/riccardotornesello/irapi-go/pkg/types"
)

type LeagueSeasonSessionsParams struct {
	LeagueId    int  `url:"league_id,omitempty,comma"`
	SeasonId    int  `url:"season_id,omitempty,comma"`
	ResultsOnly bool `url:"results_only,omitempty,comma"`
}

type LeagueSeasonSessionsResponse struct {
	Success    bool      `json:"success"`
	Subscribed bool      `json:"subscribed"`
	LeagueID   int64     `json:"league_id"`
	SeasonID   int64     `json:"season_id"`
	Sessions   []Session `json:"sessions"`
}

type Session struct {
	Cars              []Car          `json:"cars"`
	DriverChanges     bool           `json:"driver_changes"`
	EntryCount        int64          `json:"entry_count"`
	HasResults        bool           `json:"has_results"`
	LaunchAt          types.DateTime `json:"launch_at"`
	LeagueID          int64          `json:"league_id"`
	LeagueSeasonID    int64          `json:"league_season_id"`
	LoneQualify       bool           `json:"lone_qualify"`
	PaceCarClassID    interface{}    `json:"pace_car_class_id"`
	PaceCarID         interface{}    `json:"pace_car_id"`
	PasswordProtected bool           `json:"password_protected"`
	PracticeLength    int64          `json:"practice_length"`
	PrivateSessionID  int64          `json:"private_session_id"`
	QualifyLaps       int64          `json:"qualify_laps"`
	QualifyLength     int64          `json:"qualify_length"`
	RaceLaps          int64          `json:"race_laps"`
	RaceLength        int64          `json:"race_length"`
	SessionID         int64          `json:"session_id"`
	Status            int64          `json:"status"`
	SubsessionID      int64          `json:"subsession_id"`
	TeamEntryCount    int64          `json:"team_entry_count"`
	TimeLimit         int64          `json:"time_limit"`
	Track             Track          `json:"track"`
	TrackState        TrackState     `json:"track_state"`
	Weather           Weather        `json:"weather"`
	WinnerID          int64          `json:"winner_id"`
	WinnerName        string         `json:"winner_name"`
}

type Car struct {
	CarID        int64  `json:"car_id"`
	CarName      string `json:"car_name"`
	CarClassID   int64  `json:"car_class_id"`
	CarClassName string `json:"car_class_name"`
}

type Track struct {
	ConfigName string `json:"config_name"`
	TrackID    int64  `json:"track_id"`
	TrackName  string `json:"track_name"`
}

type TrackState struct {
	LeaveMarbles         bool  `json:"leave_marbles"`
	PracticeGripCompound int64 `json:"practice_grip_compound"`
	PracticeRubber       int64 `json:"practice_rubber"`
	QualifyGripCompound  int64 `json:"qualify_grip_compound"`
	QualifyRubber        int64 `json:"qualify_rubber"`
	RaceGripCompound     int64 `json:"race_grip_compound"`
	RaceRubber           int64 `json:"race_rubber"`
	WarmupGripCompound   int64 `json:"warmup_grip_compound"`
	WarmupRubber         int64 `json:"warmup_rubber"`
}

type Weather struct {
	AllowFog          bool           `json:"allow_fog"`
	Fog               int64          `json:"fog"`
	PrecipOption      int64          `json:"precip_option"`
	RelHumidity       int64          `json:"rel_humidity"`
	Skies             int64          `json:"skies"`
	TempUnits         int64          `json:"temp_units"`
	TempValue         int64          `json:"temp_value"`
	TrackWater        int64          `json:"track_water"`
	Type              int64          `json:"type"`
	Version           int64          `json:"version"`
	WeatherSummary    WeatherSummary `json:"weather_summary"`
	WeatherVarInitial int64          `json:"weather_var_initial"`
	WeatherVarOngoing int64          `json:"weather_var_ongoing"`
	WindDir           int64          `json:"wind_dir"`
	WindUnits         int64          `json:"wind_units"`
	WindValue         int64          `json:"wind_value"`
}

type WeatherSummary struct {
	MaxPrecipRateDesc string `json:"max_precip_rate_desc"`
	PrecipChance      int64  `json:"precip_chance"`
}
