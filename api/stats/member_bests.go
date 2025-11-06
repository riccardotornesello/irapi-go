package stats

import (
	"encoding/json"
)

type StatsMemberBestsResponse struct {
	CarsDriven []CarsDriven `json:"cars_driven"`
	Bests      []Best       `json:"bests"`
	CustID     int64        `json:"cust_id"`
	CarID      int64        `json:"car_id"`
}

import "time"

type Best struct {
	Track         Track     `json:"track"`
	EventType     EventType `json:"event_type"`
	BestLapTime   int64     `json:"best_lap_time"`
	SubsessionID  int64     `json:"subsession_id"`
	EndTime       time.Time `json:"end_time"`
	SeasonYear    int64     `json:"season_year"`
	SeasonQuarter int64     `json:"season_quarter"`
}

type Track struct {
	ConfigName *string `json:"config_name,omitempty"`
	TrackID    int64   `json:"track_id"`
	TrackName  string  `json:"track_name"`
}

type CarsDriven struct {
	CarID   int64  `json:"car_id"`
	CarName string `json:"car_name"`
}

type EventType string

const (
	Practice EventType = "Practice"
	Race     EventType = "Race"
)


func (api *StatsApi) MemberBests() (*StatsMemberBestsResponse, error) {
	return api.GetJson[StatsMemberBestsResponse]("/data/stats/member_bests")
}