package member

import (
	"encoding/json"
)

type MemberChartDataResponse struct {
	Blackout   bool    `json:"blackout"`
	CategoryID int64   `json:"category_id"`
	ChartType  int64   `json:"chart_type"`
	Data       []Datum `json:"data"`
	Success    bool    `json:"success"`
	CustID     int64   `json:"cust_id"`
}

type Datum struct {
	When  string `json:"when"`
	Value int64  `json:"value"`
}

func (api *MemberApi) ChartData() (*MemberChartDataResponse, error) {
	return api.GetJson[MemberChartDataResponse]("/data/member/chart_data")
}
