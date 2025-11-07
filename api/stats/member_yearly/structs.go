package member_yearly

import ()

type StatsMemberYearlyParams struct {
	CustId int `url:"cust_id,omitempty,comma"`
}

type StatsMemberYearlyResponse struct {
	Stats  []Stat `json:"stats"`
	CustID int64  `json:"cust_id"`
}

type Stat struct {
	CategoryID        int64   `json:"category_id"`
	Category          string  `json:"category"`
	Starts            int64   `json:"starts"`
	WINS              int64   `json:"wins"`
	Top5              int64   `json:"top5"`
	Poles             int64   `json:"poles"`
	AvgStartPosition  int64   `json:"avg_start_position"`
	AvgFinishPosition int64   `json:"avg_finish_position"`
	Laps              int64   `json:"laps"`
	LapsLED           int64   `json:"laps_led"`
	AvgIncidents      float64 `json:"avg_incidents"`
	AvgPoints         int64   `json:"avg_points"`
	WinPercentage     float64 `json:"win_percentage"`
	Top5Percentage    float64 `json:"top5_percentage"`
	LapsLEDPercentage float64 `json:"laps_led_percentage"`
	Year              int64   `json:"year"`
	PolesPercentage   float64 `json:"poles_percentage"`
}
