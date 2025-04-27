package stats

import (
	"encoding/json"

	"github.com/google/go-querystring/query"
	"github.com/markphelps/optional"
)

type StatsMemberRecentRacesParams struct {
	CustId *optional.Int `url:"cust_id"` // Defaults to the authenticated member.
}

type StatsMemberRecentRacesResponse struct {
	Races []struct {
		SeasonId   int    `json:"season_id"`
		SeriesId   int    `json:"series_id"`
		SeriesName string `json:"series_name"`
		CarId      int    `json:"car_id"`
		CarClassId int    `json:"car_class_id"`
		Livery     struct {
			CarId   int    `json:"car_id"`
			Pattern int    `json:"pattern"`
			Color1  string `json:"color1"`
			Color2  string `json:"color2"`
			Color3  string `json:"color3"`
		} `json:"livery"`
		LicenseLevel     int    `json:"license_level"`
		SessionStartTime string `json:"session_start_time"`
		WinnerGroupId    int    `json:"winner_group_id"`
		WinnerName       string `json:"winner_name"`
		WinnerHelmet     struct {
			Pattern    int    `json:"pattern"`
			Color1     string `json:"color1"`
			Color2     string `json:"color2"`
			Color3     string `json:"color3"`
			FaceType   int    `json:"face_type"`
			HelmetType int    `json:"helmet_type"`
		} `json:"winner_helmet"`
		WinnerLicenseLevel int `json:"winner_license_level"`
		StartPosition      int `json:"start_position"`
		FinishPosition     int `json:"finish_position"`
		QualifyingTime     int `json:"qualifying_time"`
		Laps               int `json:"laps"`
		LapsLed            int `json:"laps_led"`
		Incidents          int `json:"incidents"`
		ClubPoints         int `json:"club_points"`
		Points             int `json:"points"`
		StrengthOfField    int `json:"strength_of_field"`
		SubsessionId       int `json:"subsession_id"`
		OldSubLevel        int `json:"old_sub_level"`
		NewSubLevel        int `json:"new_sub_level"`
		OldiRating         int `json:"oldi_rating"`
		NewiRating         int `json:"newi_rating"`
		Track              struct {
			TrackId   int    `json:"track_id"`
			TrackName string `json:"track_name"`
		} `json:"track"`
		DropRace      bool `json:"drop_race"`
		SeasonYear    int  `json:"season_year"`
		SeasonQuarter int  `json:"season_quarter"`
		RaceWeekNum   int  `json:"race_week_num"`
	} `json:"races"`
	CustId int `json:"cust_id"`
}

func (api *StatsApi) GetStatsMemberRecentRaces(params StatsMemberRecentRacesParams) (*StatsMemberRecentRacesResponse, error) {
	paramsString, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	url := "/data/stats/member_recent_races?" + paramsString.Encode()

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &StatsMemberRecentRacesResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
