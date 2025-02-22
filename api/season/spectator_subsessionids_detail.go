package season

import (
	"encoding/json"

	"github.com/google/go-querystring/query"
)

type SeasonSpectatorSubsessionidsDetailParams struct {
	EventTypes *[]int `url:"event_types,omitempty"`
	SeasonIds  *[]int `url:"season_ids,omitempty"`
}

type SeasonSpectatorSubsessionidsDetailResponse struct {
	Success     bool  `json:"success"`
	SeasonIds   []int `json:"season_ids"`
	EventTypes  []int `json:"event_types"`
	Subsessions []struct {
		SubsessionId int    `json:"subsession_id"`
		SessionId    int    `json:"session_id"`
		SeasonId     int    `json:"season_id"`
		StartTime    string `json:"start_time"`
		RaceWeekNum  int    `json:"race_week_num"`
		EventType    int    `json:"event_type"`
	} `json:"subsessions"`
}

func (api *SeasonApi) GetSeasonSpectatorSubsessionidsDetail(params SeasonSpectatorSubsessionidsDetailParams) (*SeasonSpectatorSubsessionidsDetailResponse, error) {
	paramsString, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	url := "/data/season/spectator_subsessionids_detail?" + paramsString.Encode()

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &SeasonSpectatorSubsessionidsDetailResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
