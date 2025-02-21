package season

import (
	"encoding/json"

	"github.com/google/go-querystring/query"
)

type SeasonSpectatorSubsessionidsParams struct {
	EventTypes *[]int `url:"event_types,omitempty"`
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

	url := "https://members-ng.iracing.com/data/season/spectator_subsessionids?" + paramsString.Encode()

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
