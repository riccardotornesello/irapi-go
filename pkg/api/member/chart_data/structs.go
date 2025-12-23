package chart_data

type MemberChartDataParams struct {
	CustId     *int `json:"cust_id,omitempty,comma"`
	CategoryId int  `json:"category_id,comma"`
	ChartType  int  `json:"chart_type,comma"`
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
