package stats

import (
	"encoding/json"

	"github.com/google/go-querystring/query"
	"github.com/markphelps/optional"
)

type StatsMemberSummaryParams struct {
	CustId *optional.Int `url:"cust_id"` // Defaults to the authenticated member.
}

type StatsMemberSummaryResponse struct {
	ThisYear struct {
		NumOfficialSessions int `json:"num_official_sessions"`
		NumLeagueSessions   int `json:"num_league_sessions"`
		NumOfficialWins     int `json:"num_official_wins"`
		NumLeagueWins       int `json:"num_league_wins"`
	} `json:"this_year"`
	CustId int `json:"cust_id"`
}

func (api *StatsApi) GetStatsMemberSummary(params StatsMemberSummaryParams) (*StatsMemberSummaryResponse, error) {
	paramsString, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	url := "/data/stats/member_summary?" + paramsString.Encode()

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &StatsMemberSummaryResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
