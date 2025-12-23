package awards

type MemberAwardsParams struct {
	CustId *int `json:"cust_id,omitempty,comma"`
}

type MemberAwardsResponse struct {
	Type    string `json:"type"`
	Data    Data   `json:"data"`
	DataURL string `json:"data_url"`
}

type Data struct {
	Success    bool  `json:"success"`
	CustID     int64 `json:"cust_id"`
	AwardCount int64 `json:"award_count"`
}
