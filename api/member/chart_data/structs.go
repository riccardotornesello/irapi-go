package chart_data

type MemberChartDataParams struct {
	CustId     int `url:"cust_id,omitempty,comma"`
	CategoryId int `url:"category_id,omitempty,comma"`
	ChartType  int `url:"chart_type,omitempty,comma"`
}

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
