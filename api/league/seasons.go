package league

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
)

type LeagueSeasonsParams struct {
	LeagueId int   `url:"league_id"`
	Retired  *bool `url:"retired,omitempty"` // If true include seasons which are no longer active.
}

type LeagueSeasonsResponse struct {
	Subscribed bool `json:"subscribed"`
	Seasons    []struct {
		LeagueId                int    `json:"league_id"`
		SeasonId                int    `json:"season_id"`
		PointsSystemId          int    `json:"points_system_id"`
		SeasonName              string `json:"season_name"`
		Active                  bool   `json:"active"`
		Hidden                  bool   `json:"hidden"`
		NumDrops                int    `json:"num_drops"`
		NoDropsOnOrAfterRaceNum int    `json:"no_drops_on_or_after_race_num"`
		PointsCars              []struct {
			CarId   int    `json:"car_id"`
			CarName string `json:"car_name"`
		} `json:"points_cars"`
		DriverPointsCarClasses []struct {
			CarClassId  int    `json:"car_class_id"`
			Name        string `json:"name"`
			CarsInClass []struct {
				CarId   int    `json:"car_id"`
				CarName string `json:"car_name"`
			} `json:"cars_in_class"`
		} `json:"driver_points_car_classes"`
		TeamPointsCarClasses []struct {
			CarClassId  int    `json:"car_class_id"`
			Name        string `json:"name"`
			CarsInClass []struct {
				CarId   int    `json:"car_id"`
				CarName string `json:"car_name"`
			} `json:"cars_in_class"`
		} `json:"team_points_car_classes"`
		PointsSystemName string `json:"points_system_name"`
		PointsSystemDesc string `json:"points_system_desc"`
	} `json:"seasons"`
	Success  bool `json:"success"`
	Retired  bool `json:"retired"`
	LeagueId int  `json:"league_id"`
}

func (api *LeagueApi) GetLeagueSeasons(params LeagueSeasonsParams) (*LeagueSeasonsResponse, error) {
	paramsString, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	url := "/data/league/seasons?" + paramsString.Encode()

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &LeagueSeasonsResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
