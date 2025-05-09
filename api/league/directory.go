package league

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
)

type LeagueDirectoryParams struct {
	Search               *string `url:"search,omitempty"`                 // Will search against league name, description, owner, and league ID.
	Tag                  *string `url:"tag,omitempty"`                    // One or more tags, comma-separated.
	RestrictToMember     *bool   `url:"restrict_to_member,omitempty"`     // If true include only leagues for which customer is a member.
	RestrictToRecruiting *bool   `url:"restrict_to_recruiting,omitempty"` // If true include only leagues which are recruiting.
	RestrictToFriends    *bool   `url:"restrict_to_friends,omitempty"`    // If true include only leagues owned by a friend.
	RestrictToWatched    *bool   `url:"restrict_to_watched,omitempty"`    // If true include only leagues owned by a watched member.
	MinimumRosterCount   *int    `url:"minimum_roster_count,omitempty"`   // If set include leagues with at least this number of members.
	MaximumRosterCount   *int    `url:"maximum_roster_count,omitempty"`   // If set include leagues with no more than this number of members.
	Lowerbound           *int    `url:"lowerbound,omitempty"`             // First row of results to return.  Defaults to 1.
	Upperbound           *int    `url:"upperbound,omitempty"`             // Last row of results to return. Defaults to lowerbound + 39.
	Sort                 *string `url:"sort,omitempty"`                   // One of relevance, leaguename, displayname, rostercount. displayname is owners's name. Defaults to relevance.
	Order                *string `url:"order,omitempty"`                  // One of asc or desc.  Defaults to asc.
}

type LeagueDirectoryResponse struct {
	ResultsPage []struct {
		LeagueId           int    `json:"league_id"`
		OwnerId            int    `json:"owner_id"`
		LeagueName         string `json:"league_name"`
		Created            string `json:"created"`
		About              string `json:"about"`
		Url                string `json:"url"`
		RosterCount        int    `json:"roster_count"`
		Recruiting         bool   `json:"recruiting"`
		IsAdmin            bool   `json:"is_admin"`
		IsMember           bool   `json:"is_member"`
		PendingApplication bool   `json:"pending_application"`
		PendingInvitation  bool   `json:"pending_invitation"`
		Owner              struct {
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
	} `json:"results_page"`
	Success    bool `json:"success"`
	Lowerbound int  `json:"lowerbound"`
	Upperbound int  `json:"upperbound"`
	RowCount   int  `json:"row_count"`
}

func (api *LeagueApi) GetLeagueDirectory(params LeagueDirectoryParams) (*LeagueDirectoryResponse, error) {
	paramsString, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	url := "/data/league/directory?" + paramsString.Encode()

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &LeagueDirectoryResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
