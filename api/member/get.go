package member

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/markphelps/optional"
)

type MemberGetParams struct {
	CustIds         []int          `url:"cust_ids,omitempty"`
	IncludeLicenses *optional.Bool `url:"include_licenses,omitempty"`
}

type MemberGetResponse struct {
	Success bool  `json:"success"`
	CustIds []int `json:"cust_ids"`
	Members []struct {
		CustId      int    `json:"cust_id"`
		DisplayName string `json:"display_name"`
		Helmet      struct {
			Pattern    int    `json:"pattern"`
			Color1     string `json:"color1"`
			Color2     string `json:"color2"`
			Color3     string `json:"color3"`
			FaceType   int    `json:"face_type"`
			HelmetType int    `json:"helmet_type"`
		} `json:"helmet"`
		LastLogin   string `json:"last_login"`
		MemberSince string `json:"member_since"`
		ClubId      int    `json:"club_id"`
		ClubName    string `json:"club_name"`
		Ai          bool   `json:"ai"`
	} `json:"members"`
}

func (api *MemberApi) GetMember(params MemberGetParams) (*MemberGetResponse, error) {
	paramsString, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	url := "/data/member/get?" + paramsString.Encode()

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &MemberGetResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
