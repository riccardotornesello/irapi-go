package league

import (
	"encoding/json"
)

import "time"

type LeagueGetResponse struct {
	LeagueID           int64         `json:"league_id"`
	OwnerID            int64         `json:"owner_id"`
	LeagueName         string        `json:"league_name"`
	Created            time.Time     `json:"created"`
	Hidden             bool          `json:"hidden"`
	Message            string        `json:"message"`
	About              string        `json:"about"`
	URL                string        `json:"url"`
	Recruiting         bool          `json:"recruiting"`
	PrivateWall        bool          `json:"private_wall"`
	PrivateRoster      bool          `json:"private_roster"`
	PrivateSchedule    bool          `json:"private_schedule"`
	PrivateResults     bool          `json:"private_results"`
	IsOwner            bool          `json:"is_owner"`
	IsAdmin            bool          `json:"is_admin"`
	RosterCount        int64         `json:"roster_count"`
	Owner              Owner         `json:"owner"`
	Image              Image         `json:"image"`
	Tags               Tags          `json:"tags"`
	LeagueApplications []interface{} `json:"league_applications"`
	PendingRequests    []interface{} `json:"pending_requests"`
	IsMember           bool          `json:"is_member"`
	IsApplicant        bool          `json:"is_applicant"`
	IsInvite           bool          `json:"is_invite"`
	IsIgnored          bool          `json:"is_ignored"`
	Roster             []Roster      `json:"roster"`
}

type Image struct {
	SmallLogo string `json:"small_logo"`
	LargeLogo string `json:"large_logo"`
}

type Owner struct {
	CustID      int64       `json:"cust_id"`
	DisplayName string      `json:"display_name"`
	Helmet      Helmet      `json:"helmet"`
	CarNumber   interface{} `json:"car_number"`
	NickName    interface{} `json:"nick_name"`
}

type Helmet struct {
	Pattern    int64  `json:"pattern"`
	Color1     string `json:"color1"`
	Color2     string `json:"color2"`
	Color3     string `json:"color3"`
	FaceType   int64  `json:"face_type"`
	HelmetType int64  `json:"helmet_type"`
}

type Roster struct {
	CustID            int64     `json:"cust_id"`
	DisplayName       string    `json:"display_name"`
	Helmet            Helmet    `json:"helmet"`
	Owner             bool      `json:"owner"`
	Admin             bool      `json:"admin"`
	LeagueMailOptOut  bool      `json:"league_mail_opt_out"`
	LeaguePmOptOut    bool      `json:"league_pm_opt_out"`
	LeagueMemberSince time.Time `json:"league_member_since"`
	CarNumber         string    `json:"car_number"`
	NickName          *string   `json:"nick_name"`
}

type Tags struct {
	Categorized    []interface{} `json:"categorized"`
	NotCategorized []interface{} `json:"not_categorized"`
}

func (api *LeagueApi) Get() (*LeagueGetResponse, error) {
	return api.GetJson[LeagueGetResponse]("/data/league/get")
}
