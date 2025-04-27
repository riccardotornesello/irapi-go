package season

import (
	"encoding/json"

	"github.com/google/go-querystring/query"
)

type SeasonSpectatorSubsessionidsParams struct {
	EventTypes *[]int `url:"event_types"` // Types of events to include in the search. Defaults to all. ?event_types=2,3,4,5
}

type SeasonSpectatorSubsessionidsResponse struct {
	EventTypes    []int `json:"event_types"`
	Success       bool  `json:"success"`
	SubsessionIds []int `json:"subsession_ids"`
}

func (api *SeasonApi) GetSeasonSpectatorSubsessionids(params SeasonSpectatorSubsessionidsParams) (*SeasonSpectatorSubsessionidsResponse, error) {
	paramsString, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	url := "/data/season/spectator_subsessionids?" + paramsString.Encode()

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &SeasonSpectatorSubsessionidsResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
