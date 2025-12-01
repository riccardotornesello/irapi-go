package get

import (
	"github.com/riccardotornesello/irapi-go/pkg/types"
)

type TeamGetParams struct {
	TeamId          int   `url:"team_id,comma"`
	IncludeLicenses *bool `url:"include_licenses,omitempty,comma"`
}

type TeamGetResponse struct {
	TeamID           int64          `json:"team_id"`
	OwnerID          int64          `json:"owner_id"`
	TeamName         string         `json:"team_name"`
	Created          types.DateTime `json:"created"`
	Hidden           bool           `json:"hidden"`
	About            string         `json:"about"`
	URL              string         `json:"url"`
	RosterCount      int64          `json:"roster_count"`
	Recruiting       bool           `json:"recruiting"`
	PrivateWall      bool           `json:"private_wall"`
	IsDefault        bool           `json:"is_default"`
	IsOwner          bool           `json:"is_owner"`
	IsAdmin          bool           `json:"is_admin"`
	Suit             Suit           `json:"suit"`
	Owner            Owner          `json:"owner"`
	Tags             Tags           `json:"tags"`
	TeamApplications []interface{}  `json:"team_applications"`
	PendingRequests  []interface{}  `json:"pending_requests"`
	IsMember         bool           `json:"is_member"`
	IsApplicant      bool           `json:"is_applicant"`
	IsInvite         bool           `json:"is_invite"`
	IsIgnored        bool           `json:"is_ignored"`
	Roster           []Owner        `json:"roster"`
}

type Owner struct {
	CustID      int64  `json:"cust_id"`
	DisplayName string `json:"display_name"`
	Helmet      Helmet `json:"helmet"`
	Owner       bool   `json:"owner"`
	Admin       bool   `json:"admin"`
}

type Helmet struct {
	Pattern    int64  `json:"pattern"`
	Color1     string `json:"color1"`
	Color2     string `json:"color2"`
	Color3     string `json:"color3"`
	FaceType   int64  `json:"face_type"`
	HelmetType int64  `json:"helmet_type"`
}

type Suit struct {
	Pattern int64  `json:"pattern"`
	Color1  string `json:"color1"`
	Color2  string `json:"color2"`
	Color3  string `json:"color3"`
}

type Tags struct {
	Categorized    []interface{} `json:"categorized"`
	NotCategorized []interface{} `json:"not_categorized"`
}
