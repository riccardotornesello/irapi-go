package member

import (
	"encoding/json"

	"github.com/google/go-querystring/query"
	"github.com/markphelps/optional"
)

type MemberChartDataParams struct {
	CustId     *optional.Int `url:"cust_id"`     // Defaults to the authenticated member.
	CategoryId int           `url:"category_id"` // 1 - Oval; 2 - Road; 3 - Dirt oval; 4 - Dirt road
	ChartType  int           `url:"chart_type"`  // 1 - iRating; 2 - TT Rating; 3 - License/SR
}

type MemberChartDataResponse struct {
	Blackout   bool `json:"blackout"`
	CategoryId int  `json:"category_id"`
	ChartType  int  `json:"chart_type"`
	Data       []struct {
		When  string `json:"when"`
		Value int    `json:"value"`
	} `json:"data"`
	Success bool `json:"success"`
	CustId  int  `json:"cust_id"`
}

func (api *MemberApi) GetMemberChartData(params MemberChartDataParams) (*MemberChartDataResponse, error) {
	paramsString, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	url := "/data/member/chart_data?" + paramsString.Encode()

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &MemberChartDataResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
