package member

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/markphelps/optional"
)

type MemberAwardInstancesParams struct {
	CustId  *optional.Int `url:"cust_id,omitempty"` // Defaults to the authenticated member.
	AwardId int           `url:"award_id,omitempty"`
}

type MemberAwardInstancesResponse struct {
	Type string `json:"type"`
	Data struct {
		Success    bool `json:"success"`
		CustId     int  `json:"cust_id"`
		AwardId    int  `json:"award_id"`
		AwardCount int  `json:"award_count"`
	} `json:"data"`
	DataUrl string `json:"data_url"`
}

func (api *MemberApi) GetMemberAwardInstances(params MemberAwardInstancesParams) (*MemberAwardInstancesResponse, error) {
	paramsString, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	url := "/data/member/award_instances?" + paramsString.Encode()

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &MemberAwardInstancesResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
