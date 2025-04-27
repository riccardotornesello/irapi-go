package league

import (
	"encoding/json"

	"github.com/google/go-querystring/query"
	"github.com/markphelps/optional"
)

type LeagueGetParams struct {
	LeagueId        int            `url:"league_id"`
	IncludeLicenses *optional.Bool `url:"include_licenses"` // For faster responses, only request when necessary.
}

type LeagueGetResponse struct {
	LeagueId        int    `json:"league_id"`
	OwnerId         int    `json:"owner_id"`
	LeagueName      string `json:"league_name"`
	Created         string `json:"created"`
	Hidden          bool   `json:"hidden"`
	Message         string `json:"message"`
	About           string `json:"about"`
	Url             string `json:"url"`
	Recruiting      bool   `json:"recruiting"`
	PrivateWall     bool   `json:"private_wall"`
	PrivateRoster   bool   `json:"private_roster"`
	PrivateSchedule bool   `json:"private_schedule"`
	PrivateResults  bool   `json:"private_results"`
	IsOwner         bool   `json:"is_owner"`
	IsAdmin         bool   `json:"is_admin"`
	RosterCount     int    `json:"roster_count"`
	Owner           struct {
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
		CarNumber interface{} `json:"car_number"`
		NickName  interface{} `json:"nick_name"`
	} `json:"owner"`
	Image struct {
		SmallLogo string `json:"small_logo"`
		LargeLogo string `json:"large_logo"`
	} `json:"image"`
	Tags struct {
		Categorized    []interface{} `json:"categorized"`
		NotCategorized []interface{} `json:"not_categorized"`
	} `json:"tags"`
	LeagueApplications []interface{} `json:"league_applications"`
	PendingRequests    []interface{} `json:"pending_requests"`
	IsMember           bool          `json:"is_member"`
	IsApplicant        bool          `json:"is_applicant"`
	IsInvite           bool          `json:"is_invite"`
	IsIgnored          bool          `json:"is_ignored"`
	Roster             []struct {
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
		Owner             bool        `json:"owner"`
		Admin             bool        `json:"admin"`
		LeagueMailOptOut  bool        `json:"league_mail_opt_out"`
		LeaguePmOptOut    bool        `json:"league_pm_opt_out"`
		LeagueMemberSince string      `json:"league_member_since"`
		CarNumber         *string     `json:"car_number"`
		NickName          interface{} `json:"nick_name"`
	} `json:"roster"`
}

func (api *LeagueApi) GetLeague(params LeagueGetParams) (*LeagueGetResponse, error) {
	paramsString, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	url := "/data/league/get?" + paramsString.Encode()

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &LeagueGetResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
