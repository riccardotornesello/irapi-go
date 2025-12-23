package award_instances

type MemberAwardInstancesParams struct {
	CustId  *int `json:"cust_id,omitempty,comma"`
	AwardId int  `json:"award_id,comma"`
}

type MemberAwardInstancesResponse struct {
	Type    string `json:"type"`
	Data    Data   `json:"data"`
	DataURL string `json:"data_url"`
}

type Data struct {
	Success    bool  `json:"success"`
	CustID     int64 `json:"cust_id"`
	AwardID    int64 `json:"award_id"`
	AwardCount int64 `json:"award_count"`
}
