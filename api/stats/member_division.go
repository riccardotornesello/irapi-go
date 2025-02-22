package stats

import (
	"encoding/json"

	"github.com/google/go-querystring/query"
)

type StatsMemberDivisionParams struct {
	SeasonId  int `url:"season_id,omitempty"`
	EventType int `url:"event_type,omitempty"`
}

type StatsMemberDivisionResponse struct {
	Division  int  `json:"division"`
	Projected bool `json:"projected"`
	EventType int  `json:"event_type"`
	Success   bool `json:"success"`
	SeasonId  int  `json:"season_id"`
}

// Divisions are 0-based: 0 is Division 1, 10 is Rookie. See /data/constants/divisons for more information. Always for the authenticated member.
func (api *StatsApi) GetStatsMemberDivision(params StatsMemberDivisionParams) (*StatsMemberDivisionResponse, error) {
	paramsString, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	url := "https://members-ng.iracing.com/data/stats/member_division?" + paramsString.Encode()

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &StatsMemberDivisionResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
