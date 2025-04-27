package stats

import (
	"encoding/json"

	"github.com/google/go-querystring/query"
	"github.com/markphelps/optional"
)

type StatsMemberYearlyParams struct {
	CustId *optional.Int `url:"cust_id"` // Defaults to the authenticated member.
}

type StatsMemberYearlyResponse struct {
	Stats []struct {
		CategoryId        int     `json:"category_id"`
		Category          string  `json:"category"`
		Starts            int     `json:"starts"`
		Wins              int     `json:"wins"`
		Top5              int     `json:"top5"`
		Poles             int     `json:"poles"`
		AvgStartPosition  int     `json:"avg_start_position"`
		AvgFinishPosition int     `json:"avg_finish_position"`
		Laps              int     `json:"laps"`
		LapsLed           int     `json:"laps_led"`
		AvgIncidents      float64 `json:"avg_incidents"`
		AvgPoints         int     `json:"avg_points"`
		WinPercentage     float64 `json:"win_percentage"`
		Top5Percentage    float64 `json:"top5_percentage"`
		LapsLedPercentage float64 `json:"laps_led_percentage"`
		TotalClubPoints   int     `json:"total_club_points"`
		Year              int     `json:"year"`
		PolesPercentage   float64 `json:"poles_percentage"`
	} `json:"stats"`
	CustId int `json:"cust_id"`
}

func (api *StatsApi) GetStatsMemberYearly(params StatsMemberYearlyParams) (*StatsMemberYearlyResponse, error) {
	paramsString, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	url := "/data/stats/member_yearly?" + paramsString.Encode()

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &StatsMemberYearlyResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
