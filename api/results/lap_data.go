package results

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
)

type ResultsLapDataParams struct {
	SubsessionId     int  `url:"subsession_id"`
	SimsessionNumber int  `url:"simsession_number"` // The main event is 0; the preceding event is -1, and so on.
	CustId           *int `url:"cust_id,omitempty"` // Required if the subsession was a single-driver event. Optional for team events. If omitted for a team event then the laps driven by all the team's drivers will be included.
	TeamId           *int `url:"team_id,omitempty"` // Required if the subsession was a team event.
}

type ResultsLapDataResponseChunks []struct {
	GroupId          int         `json:"group_id"`
	Name             string      `json:"name"`
	CustId           int         `json:"cust_id"`
	DisplayName      string      `json:"display_name"`
	LapNumber        int         `json:"lap_number"`
	Flags            int         `json:"flags"`
	Incident         bool        `json:"incident"`
	SessionTime      int         `json:"session_time"`
	SessionStartTime interface{} `json:"session_start_time"`
	LapTime          int         `json:"lap_time"`
	TeamFastestLap   bool        `json:"team_fastest_lap"`
	PersonalBestLap  bool        `json:"personal_best_lap"`
	Helmet           struct {
		Pattern    int    `json:"pattern"`
		Color1     string `json:"color1"`
		Color2     string `json:"color2"`
		Color3     string `json:"color3"`
		FaceType   int    `json:"face_type"`
		HelmetType int    `json:"helmet_type"`
	} `json:"helmet"`
	LicenseLevel int      `json:"license_level"`
	CarNumber    string   `json:"car_number"`
	LapEvents    []string `json:"lap_events"`
	Ai           bool     `json:"ai"`
}

type ResultsLapDataResponse struct {
	Success     bool `json:"success"`
	SessionInfo struct {
		SubsessionId          int    `json:"subsession_id"`
		SessionId             int    `json:"session_id"`
		SimsessionNumber      int    `json:"simsession_number"`
		SimsessionType        int    `json:"simsession_type"`
		SimsessionName        string `json:"simsession_name"`
		NumLapsForQualAverage int    `json:"num_laps_for_qual_average"`
		NumLapsForSoloAverage int    `json:"num_laps_for_solo_average"`
		EventType             int    `json:"event_type"`
		EventTypeName         string `json:"event_type_name"`
		PrivateSessionId      int    `json:"private_session_id"`
		SeasonName            string `json:"season_name"`
		SeasonShortName       string `json:"season_short_name"`
		SeriesName            string `json:"series_name"`
		SeriesShortName       string `json:"series_short_name"`
		SessionName           string `json:"session_name"`
		RestrictResults       bool   `json:"restrict_results"`
		StartTime             string `json:"start_time"`
		Track                 struct {
			ConfigName string `json:"config_name"`
			TrackId    int    `json:"track_id"`
			TrackName  string `json:"track_name"`
		} `json:"track"`
	} `json:"session_info"`
	BestLapNum      int         `json:"best_lap_num"`
	BestLapTime     int         `json:"best_lap_time"`
	BestNlapsNum    int         `json:"best_nlaps_num"`
	BestNlapsTime   int         `json:"best_nlaps_time"`
	BestQualLapNum  int         `json:"best_qual_lap_num"`
	BestQualLapTime int         `json:"best_qual_lap_time"`
	BestQualLapAt   interface{} `json:"best_qual_lap_at"`
	ChunkInfo       struct {
		ChunkSize       int      `json:"chunk_size"`
		NumChunks       int      `json:"num_chunks"`
		Rows            int      `json:"rows"`
		BaseDownloadUrl string   `json:"base_download_url"`
		ChunkFileNames  []string `json:"chunk_file_names"`
	} `json:"chunk_info"`
	LastUpdated  string `json:"last_updated"`
	GroupId      int    `json:"group_id"`
	CustId       int    `json:"cust_id"`
	Name         string `json:"name"`
	CarId        int    `json:"car_id"`
	LicenseLevel int    `json:"license_level"`
	Livery       struct {
		CarId        int    `json:"car_id"`
		Pattern      int    `json:"pattern"`
		Color1       string `json:"color1"`
		Color2       string `json:"color2"`
		Color3       string `json:"color3"`
		NumberFont   int    `json:"number_font"`
		NumberColor1 string `json:"number_color1"`
		NumberColor2 string `json:"number_color2"`
		NumberColor3 string `json:"number_color3"`
		NumberSlant  int    `json:"number_slant"`
		Sponsor1     int    `json:"sponsor1"`
		Sponsor2     int    `json:"sponsor2"`
		CarNumber    string `json:"car_number"`
		WheelColor   string `json:"wheel_color"`
		RimType      int    `json:"rim_type"`
	} `json:"livery"`
	Chunks ResultsLapDataResponseChunks
}

func (api *ResultsApi) GetResultsLapData(params ResultsLapDataParams) (*ResultsLapDataResponse, error) {
	paramsString, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	url := "/data/results/lap_data?" + paramsString.Encode()

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &ResultsLapDataResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
