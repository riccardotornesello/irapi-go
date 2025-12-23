package member_bests

import (
	"github.com/riccardotornesello/irapi-go/pkg/types"
)

type StatsMemberBestsParams struct {
	CustId *int `json:"cust_id,omitempty,comma"`
	CarId  *int `json:"car_id,omitempty,comma"`
}

type StatsMemberBestsResponse struct {
	CarsDriven []CarsDriven `json:"cars_driven"`
	Bests      []Best       `json:"bests"`
	CustID     int64        `json:"cust_id"`
	CarID      int64        `json:"car_id"`
}

type Best struct {
	Track         Track          `json:"track"`
	EventType     string         `json:"event_type"`
	BestLapTime   int64          `json:"best_lap_time"`
	SubsessionID  int64          `json:"subsession_id"`
	EndTime       types.DateTime `json:"end_time"`
	SeasonYear    int64          `json:"season_year"`
	SeasonQuarter int64          `json:"season_quarter"`
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
