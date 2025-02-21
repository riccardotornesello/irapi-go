package league

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/markphelps/optional"
)

type LeagueDirectoryParams struct {
	Search               *optional.String `url:"search,omitempty"`
	Tag                  *optional.String `url:"tag,omitempty"`
	RestrictToMember     *optional.Bool   `url:"restrict_to_member,omitempty"`
	RestrictToRecruiting *optional.Bool   `url:"restrict_to_recruiting,omitempty"`
	RestrictToFriends    *optional.Bool   `url:"restrict_to_friends,omitempty"`
	RestrictToWatched    *optional.Bool   `url:"restrict_to_watched,omitempty"`
	MinimumRosterCount   *optional.Int    `url:"minimum_roster_count,omitempty"`
	MaximumRosterCount   *optional.Int    `url:"maximum_roster_count,omitempty"`
	Lowerbound           *optional.Int    `url:"lowerbound,omitempty"`
	Upperbound           *optional.Int    `url:"upperbound,omitempty"`
	Sort                 *optional.String `url:"sort,omitempty"`
	Order                *optional.String `url:"order,omitempty"`
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

	url := "https://members-ng.iracing.com/data/league/directory?" + paramsString.Encode()

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
