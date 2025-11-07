package directory

import (
	"time"
)

type LeagueDirectoryParams struct {
	Search               string `url:"search,omitempty,comma"`
	Tag                  string `url:"tag,omitempty,comma"`
	RestrictToMember     bool   `url:"restrict_to_member,omitempty,comma"`
	RestrictToRecruiting bool   `url:"restrict_to_recruiting,omitempty,comma"`
	RestrictToFriends    bool   `url:"restrict_to_friends,omitempty,comma"`
	RestrictToWatched    bool   `url:"restrict_to_watched,omitempty,comma"`
	MinimumRosterCount   int    `url:"minimum_roster_count,omitempty,comma"`
	MaximumRosterCount   int    `url:"maximum_roster_count,omitempty,comma"`
	Lowerbound           int    `url:"lowerbound,omitempty,comma"`
	Upperbound           int    `url:"upperbound,omitempty,comma"`
	Sort                 string `url:"sort,omitempty,comma"`
	Order                string `url:"order,omitempty,comma"`
}

type LeagueDirectoryResponse struct {
	ResultsPage []ResultsPage `json:"results_page"`
	Success     bool          `json:"success"`
	Lowerbound  int64         `json:"lowerbound"`
	Upperbound  int64         `json:"upperbound"`
	RowCount    int64         `json:"row_count"`
}

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
