package stats

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/markphelps/optional"
)

type StatsMemberRecapParams struct {
	CustId *optional.Int `url:"cust_id,omitempty"`
	Year   *optional.Int `url:"year,omitempty"`
	Season *optional.Int `url:"season,omitempty"`
}

type StatsMemberRecapResponse struct {
	Year  int `json:"year"`
	Stats struct {
		Starts            int `json:"starts"`
		Wins              int `json:"wins"`
		Top5              int `json:"top5"`
		AvgStartPosition  int `json:"avg_start_position"`
		AvgFinishPosition int `json:"avg_finish_position"`
		Laps              int `json:"laps"`
		LapsLed           int `json:"laps_led"`
		FavoriteCar       struct {
			CarId    int    `json:"car_id"`
			CarName  string `json:"car_name"`
			CarImage string `json:"car_image"`
		} `json:"favorite_car"`
		FavoriteTrack struct {
			ConfigName string `json:"config_name"`
			TrackId    int    `json:"track_id"`
			TrackLogo  string `json:"track_logo"`
			TrackName  string `json:"track_name"`
		} `json:"favorite_track"`
	} `json:"stats"`
	Success bool        `json:"success"`
	Season  interface{} `json:"season"`
	CustId  int         `json:"cust_id"`
}

func (api *StatsApi) GetStatsMemberRecap(params StatsMemberRecapParams) (*StatsMemberRecapResponse, error) {
	paramsString, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	url := "https://members-ng.iracing.com/data/stats/member_recap?" + paramsString.Encode()

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &StatsMemberRecapResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
