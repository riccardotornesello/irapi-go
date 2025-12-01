package lap_data

import (
	"github.com/riccardotornesello/irapi-go/client"
	"github.com/riccardotornesello/irapi-go/pkg/types"
)

type ResultsLapDataParams struct {
	SubsessionId     int  `url:"subsession_id,comma"`
	SimsessionNumber int  `url:"simsession_number,comma"`
	CustId           *int `url:"cust_id,omitempty,comma"`
	TeamId           *int `url:"team_id,omitempty,comma"`
}

type ResultsLapDataResponse struct {
	Success         bool                          `json:"success"`
	SessionInfo     SessionInfo                   `json:"session_info"`
	BestLapNum      int64                         `json:"best_lap_num"`
	BestLapTime     int64                         `json:"best_lap_time"`
	BestNlapsNum    int64                         `json:"best_nlaps_num"`
	BestNlapsTime   int64                         `json:"best_nlaps_time"`
	BestQualLapNum  int64                         `json:"best_qual_lap_num"`
	BestQualLapTime int64                         `json:"best_qual_lap_time"`
	BestQualLapAt   interface{}                   `json:"best_qual_lap_at"`
	ChunkInfo       client.IRacingChunkInfo       `json:"chunk_info"`
	LastUpdated     types.DateTime                `json:"last_updated"`
	GroupID         int64                         `json:"group_id"`
	CustID          int64                         `json:"cust_id"`
	Name            string                        `json:"name"`
	CarID           int64                         `json:"car_id"`
	LicenseLevel    int64                         `json:"license_level"`
	Livery          Livery                        `json:"livery"`
	Chunks          []ResultsLapDataResponseChunk `json:"chunks"`
}

type Livery struct {
	CarID        int64  `json:"car_id"`
	Pattern      int64  `json:"pattern"`
	Color1       string `json:"color1"`
	Color2       string `json:"color2"`
	Color3       string `json:"color3"`
	NumberFont   int64  `json:"number_font"`
	NumberColor1 string `json:"number_color1"`
	NumberColor2 string `json:"number_color2"`
	NumberColor3 string `json:"number_color3"`
	NumberSlant  int64  `json:"number_slant"`
	Sponsor1     int64  `json:"sponsor1"`
	Sponsor2     int64  `json:"sponsor2"`
	CarNumber    string `json:"car_number"`
	WheelColor   string `json:"wheel_color"`
	RimType      int64  `json:"rim_type"`
}

type SessionInfo struct {
	SubsessionID          int64          `json:"subsession_id"`
	SessionID             int64          `json:"session_id"`
	SimsessionNumber      int64          `json:"simsession_number"`
	SimsessionType        int64          `json:"simsession_type"`
	SimsessionName        string         `json:"simsession_name"`
	NumLapsForQualAverage int64          `json:"num_laps_for_qual_average"`
	NumLapsForSoloAverage int64          `json:"num_laps_for_solo_average"`
	EventType             int64          `json:"event_type"`
	EventTypeName         string         `json:"event_type_name"`
	PrivateSessionID      int64          `json:"private_session_id"`
	SeasonName            string         `json:"season_name"`
	SeasonShortName       string         `json:"season_short_name"`
	SeriesName            string         `json:"series_name"`
	SeriesShortName       string         `json:"series_short_name"`
	SessionName           string         `json:"session_name"`
	RestrictResults       bool           `json:"restrict_results"`
	StartTime             types.DateTime `json:"start_time"`
	Track                 Track          `json:"track"`
}

type Track struct {
	ConfigName string `json:"config_name"`
	TrackID    int64  `json:"track_id"`
	TrackName  string `json:"track_name"`
}

type ResultsLapDataResponseChunk struct {
	GroupID          int64       `json:"group_id"`
	Name             string      `json:"name"`
	CustID           int64       `json:"cust_id"`
	DisplayName      string      `json:"display_name"`
	LapNumber        int64       `json:"lap_number"`
	Flags            int64       `json:"flags"`
	Incident         bool        `json:"incident"`
	SessionTime      int64       `json:"session_time"`
	SessionStartTime interface{} `json:"session_start_time"`
	LapTime          int64       `json:"lap_time"`
	TeamFastestLap   bool        `json:"team_fastest_lap"`
	PersonalBestLap  bool        `json:"personal_best_lap"`
	Helmet           Helmet      `json:"helmet"`
	LicenseLevel     int64       `json:"license_level"`
	CarNumber        string      `json:"car_number"`
	LapEvents        []string    `json:"lap_events"`
	AI               bool        `json:"ai"`
}

type Helmet struct {
	Pattern    int64  `json:"pattern"`
	Color1     string `json:"color1"`
	Color2     string `json:"color2"`
	Color3     string `json:"color3"`
	FaceType   int64  `json:"face_type"`
	HelmetType int64  `json:"helmet_type"`
}
