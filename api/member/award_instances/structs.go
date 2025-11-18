package award_instances

type MemberAwardInstancesParams struct {
	CustId  int `url:"cust_id,omitempty,comma"`
	AwardId int `url:"award_id,omitempty,comma"`
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
