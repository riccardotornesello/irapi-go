package league

import (
	"encoding/json"
)

type LeagueDirectoryResponse struct {
	ResultsPage []ResultsPage `json:"results_page"`
	Success     bool          `json:"success"`
	Lowerbound  int64         `json:"lowerbound"`
	Upperbound  int64         `json:"upperbound"`
	RowCount    int64         `json:"row_count"`
}

import "time"

type ResultsPage struct {
	LeagueID           int64     `json:"league_id"`
	OwnerID            int64     `json:"owner_id"`
	LeagueName         string    `json:"league_name"`
	Created            time.Time `json:"created"`
	About              *string   `json:"about,omitempty"`
	URL                *string   `json:"url,omitempty"`
	RosterCount        int64     `json:"roster_count"`
	Recruiting         bool      `json:"recruiting"`
	IsAdmin            bool      `json:"is_admin"`
	IsMember           bool      `json:"is_member"`
	PendingApplication bool      `json:"pending_application"`
	PendingInvitation  bool      `json:"pending_invitation"`
	Owner              Owner     `json:"owner"`
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


func (api *LeagueApi) Directory() (*LeagueDirectoryResponse, error) {
	return api.GetJson[LeagueDirectoryResponse]("/data/league/directory")
}