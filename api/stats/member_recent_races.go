package stats

import (
	"encoding/json"
)

type StatsMemberRecentRacesResponse struct {
	Races  []Race `json:"races"`
	CustID int64  `json:"cust_id"`
}

import "time"

type Race struct {
	SeasonID           int64        `json:"season_id"`
	SeriesID           int64        `json:"series_id"`
	SeriesName         string       `json:"series_name"`
	CarID              int64        `json:"car_id"`
	CarClassID         int64        `json:"car_class_id"`
	Livery             Livery       `json:"livery"`
	LicenseLevel       int64        `json:"license_level"`
	SessionStartTime   time.Time    `json:"session_start_time"`
	WinnerGroupID      int64        `json:"winner_group_id"`
	WinnerName         string       `json:"winner_name"`
	WinnerHelmet       WinnerHelmet `json:"winner_helmet"`
	WinnerLicenseLevel int64        `json:"winner_license_level"`
	StartPosition      int64        `json:"start_position"`
	FinishPosition     int64        `json:"finish_position"`
	QualifyingTime     int64        `json:"qualifying_time"`
	Laps               int64        `json:"laps"`
	LapsLED            int64        `json:"laps_led"`
	Incidents          int64        `json:"incidents"`
	Points             int64        `json:"points"`
	StrengthOfField    int64        `json:"strength_of_field"`
	SubsessionID       int64        `json:"subsession_id"`
	OldSubLevel        int64        `json:"old_sub_level"`
	NewSubLevel        int64        `json:"new_sub_level"`
	OldiRating         int64        `json:"oldi_rating"`
	NewiRating         int64        `json:"newi_rating"`
	Track              Track        `json:"track"`
	DropRace           bool         `json:"drop_race"`
	SeasonYear         int64        `json:"season_year"`
	SeasonQuarter      int64        `json:"season_quarter"`
	RaceWeekNum        int64        `json:"race_week_num"`
}

type Livery struct {
	CarID   int64  `json:"car_id"`
	Pattern int64  `json:"pattern"`
	Color1  Color1 `json:"color1"`
	Color2  Color2 `json:"color2"`
	Color3  Color3 `json:"color3"`
}

type Track struct {
	TrackID   int64     `json:"track_id"`
	TrackName TrackName `json:"track_name"`
}

type WinnerHelmet struct {
	Pattern    int64  `json:"pattern"`
	Color1     string `json:"color1"`
	Color2     string `json:"color2"`
	Color3     string `json:"color3"`
	FaceType   int64  `json:"face_type"`
	HelmetType int64  `json:"helmet_type"`
}

type Color1 string

const (
	Ff0000    Color1 = "FF0000"
	The000000 Color1 = "000000"
	The14213D Color1 = "14213d"
)

type Color2 string

const (
	Fca311    Color2 = "fca311"
	Ffec41    Color2 = "ffec41"
	The00Ff00 Color2 = "00FF00"
)

type Color3 string

const (
	E5E5E5    Color3 = "e5e5e5"
	Fc1212    Color3 = "fc1212"
	The0000Ff Color3 = "0000FF"
)

type TrackName string

const (
	CircuitDeSPAFrancorchamps   TrackName = "Circuit de Spa-Francorchamps"
	OultonParkCircuit           TrackName = "Oulton Park Circuit"
	SebringInternationalRaceway TrackName = "Sebring International Raceway"
)


func (api *StatsApi) MemberRecentRaces() (*StatsMemberRecentRacesResponse, error) {
	return api.GetJson[StatsMemberRecentRacesResponse]("/data/stats/member_recent_races")
}