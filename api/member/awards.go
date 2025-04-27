package member

import (
	"encoding/json"

	"github.com/google/go-querystring/query"
	"github.com/markphelps/optional"
)

type MemberAwardsParams struct {
	CustId *optional.Int `url:"cust_id"` // Defaults to the authenticated member.
}

type MemberAwardsResponse struct {
	Type string `json:"type"`
	Data struct {
		Success    bool `json:"success"`
		CustId     int  `json:"cust_id"`
		AwardCount int  `json:"award_count"`
	} `json:"data"`
	DataUrl string `json:"data_url"`
}

func (api *MemberApi) GetMemberAwards(params MemberAwardsParams) (*MemberAwardsResponse, error) {
	paramsString, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	url := "/data/member/awards?" + paramsString.Encode()

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &MemberAwardsResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
