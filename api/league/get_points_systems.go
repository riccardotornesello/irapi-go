package league

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/markphelps/optional"
)

type LeagueGetPointsSystemsParams struct {
	LeagueId int           `url:"league_id,omitempty"`
	SeasonId *optional.Int `url:"season_id,omitempty"`
}

type LeagueGetPointsSystemsResponse struct {
	Subscribed    bool `json:"subscribed"`
	Success       bool `json:"success"`
	PointsSystems []struct {
		PointsSystemId int    `json:"points_system_id"`
		Name           string `json:"name"`
		Description    string `json:"description"`
		LeagueId       int    `json:"league_id"`
		Retired        bool   `json:"retired"`
		IracingSystem  bool   `json:"iracing_system"`
	} `json:"points_systems"`
	LeagueId int `json:"league_id"`
}

func (api *LeagueApi) GetLeagueGetPointsSystems(params LeagueGetPointsSystemsParams) (*LeagueGetPointsSystemsResponse, error) {
	paramsString, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	url := "https://members-ng.iracing.com/data/league/get_points_systems?" + paramsString.Encode()

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &LeagueGetPointsSystemsResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
