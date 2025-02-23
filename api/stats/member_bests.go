package stats

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/markphelps/optional"
)

type StatsMemberBestsParams struct {
	CustId *optional.Int `url:"cust_id,omitempty"` // Defaults to the authenticated member.
	CarId  *optional.Int `url:"car_id,omitempty"`  // First call should exclude car_id; use cars_driven list in return for subsequent calls.
}

type StatsMemberBestsResponse struct {
	CarsDriven []struct {
		CarId   int    `json:"car_id"`
		CarName string `json:"car_name"`
	} `json:"cars_driven"`
	Bests []struct {
		Track struct {
			ConfigName string `json:"config_name"`
			TrackId    int    `json:"track_id"`
			TrackName  string `json:"track_name"`
		} `json:"track"`
		EventType     string `json:"event_type"`
		BestLapTime   int    `json:"best_lap_time"`
		SubsessionId  int    `json:"subsession_id"`
		EndTime       string `json:"end_time"`
		SeasonYear    int    `json:"season_year"`
		SeasonQuarter int    `json:"season_quarter"`
	} `json:"bests"`
	CustId int `json:"cust_id"`
	CarId  int `json:"car_id"`
}

func (api *StatsApi) GetStatsMemberBests(params StatsMemberBestsParams) (*StatsMemberBestsResponse, error) {
	paramsString, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	url := "/data/stats/member_bests?" + paramsString.Encode()

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &StatsMemberBestsResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
