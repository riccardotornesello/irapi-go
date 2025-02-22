package league

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/markphelps/optional"
)

type LeagueRosterParams struct {
	LeagueId        int            `url:"league_id,omitempty"`
	IncludeLicenses *optional.Bool `url:"include_licenses,omitempty"`
}

type LeagueRosterResponse struct {
	Type string `json:"type"`
	Data struct {
		Subscribed  bool `json:"subscribed"`
		Success     bool `json:"success"`
		RosterCount int  `json:"roster_count"`
		LeagueId    int  `json:"league_id"`
	} `json:"data"`
	DataUrl string `json:"data_url"`
}

func (api *LeagueApi) GetLeagueRoster(params LeagueRosterParams) (*LeagueRosterResponse, error) {
	paramsString, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	url := "/data/league/roster?" + paramsString.Encode()

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &LeagueRosterResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
