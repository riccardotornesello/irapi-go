package league

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/markphelps/optional"
)

type LeagueSeasonStandingsParams struct {
	LeagueId   int           `url:"league_id,omitempty"`
	SeasonId   int           `url:"season_id,omitempty"`
	CarClassId *optional.Int `url:"car_class_id,omitempty"`
	CarId      *optional.Int `url:"car_id,omitempty"` // If car_class_id is included then the standings are for the car in that car class, otherwise they are for the car across car classes.
}

type LeagueSeasonStandingsResponse struct {
	CarClassId int  `json:"car_class_id"`
	Success    bool `json:"success"`
	SeasonId   int  `json:"season_id"`
	CarId      int  `json:"car_id"`
	Standings  struct {
		DriverStandings []struct {
			Rownum   int `json:"rownum"`
			Position int `json:"position"`
			Driver   struct {
				CustId      int    `json:"cust_id"`
				DisplayName string `json:"display_name"`
				Helmet      struct {
					Pattern    int    `json:"pattern"`
					Color1     string `json:"color1"`
					Color2     string `json:"color2"`
					Color3     string `json:"color3"`
					FaceType   int    `json:"face_type"`
					HelmetType int    `json:"helmet_type"`
				} `json:"helmet"`
			} `json:"driver"`
			CarNumber           *string     `json:"car_number"`
			DriverNickname      interface{} `json:"driver_nickname"`
			Wins                int         `json:"wins"`
			AverageStart        int         `json:"average_start"`
			AverageFinish       int         `json:"average_finish"`
			BasePoints          int         `json:"base_points"`
			NegativeAdjustments int         `json:"negative_adjustments"`
			PositiveAdjustments int         `json:"positive_adjustments"`
			TotalAdjustments    int         `json:"total_adjustments"`
			TotalPoints         int         `json:"total_points"`
		} `json:"driver_standings"`
		TeamStandings         []interface{} `json:"team_standings"`
		DriverStandingsCsvUrl string        `json:"driver_standings_csv_url"`
		TeamStandingsCsvUrl   string        `json:"team_standings_csv_url"`
	} `json:"standings"`
	LeagueId int `json:"league_id"`
}

func (api *LeagueApi) GetLeagueSeasonStandings(params LeagueSeasonStandingsParams) (*LeagueSeasonStandingsResponse, error) {
	paramsString, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	url := "/data/league/season_standings?" + paramsString.Encode()

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &LeagueSeasonStandingsResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
