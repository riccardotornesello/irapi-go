package team

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/markphelps/optional"
)

type TeamGetParams struct {
	TeamId          int            `url:"team_id,omitempty"`
	IncludeLicenses *optional.Bool `url:"include_licenses,omitempty"`
}

type TeamGetResponse struct {
	TeamId      int    `json:"team_id"`
	OwnerId     int    `json:"owner_id"`
	TeamName    string `json:"team_name"`
	Created     string `json:"created"`
	Hidden      bool   `json:"hidden"`
	About       string `json:"about"`
	Url         string `json:"url"`
	RosterCount int    `json:"roster_count"`
	Recruiting  bool   `json:"recruiting"`
	PrivateWall bool   `json:"private_wall"`
	IsDefault   bool   `json:"is_default"`
	IsOwner     bool   `json:"is_owner"`
	IsAdmin     bool   `json:"is_admin"`
	Suit        struct {
		Pattern int    `json:"pattern"`
		Color1  string `json:"color1"`
		Color2  string `json:"color2"`
		Color3  string `json:"color3"`
	} `json:"suit"`
	Owner struct {
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
		Owner bool `json:"owner"`
		Admin bool `json:"admin"`
	} `json:"owner"`
	Tags struct {
		Categorized    []interface{} `json:"categorized"`
		NotCategorized []interface{} `json:"not_categorized"`
	} `json:"tags"`
	TeamApplications []interface{} `json:"team_applications"`
	PendingRequests  []interface{} `json:"pending_requests"`
	IsMember         bool          `json:"is_member"`
	IsApplicant      bool          `json:"is_applicant"`
	IsInvite         bool          `json:"is_invite"`
	IsIgnored        bool          `json:"is_ignored"`
	Roster           []struct {
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
		Owner bool `json:"owner"`
		Admin bool `json:"admin"`
	} `json:"roster"`
}

func (api *TeamApi) GetTeam(params TeamGetParams) (*TeamGetResponse, error) {
	paramsString, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	url := "https://members-ng.iracing.com/data/team/get?" + paramsString.Encode()

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &TeamGetResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
