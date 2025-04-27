package league

import (
	"encoding/json"

	"github.com/google/go-querystring/query"
	"github.com/markphelps/optional"
)

type LeagueMembershipParams struct {
	CustId        *optional.Int  `url:"cust_id"` // If different from the authenticated member, the following restrictions apply: - Caller cannot be on requested customer's block list or an empty list will result; - Requested customer cannot have their online activity preference set to hidden or an empty list will result; - Only leagues for which the requested customer is an admin and the league roster is not private are returned.
	IncludeLeague *optional.Bool `url:"include_league"`
}

type LeagueMembershipResponse []struct {
	LeagueId         int         `json:"league_id"`
	LeagueName       string      `json:"league_name"`
	Owner            bool        `json:"owner"`
	Admin            bool        `json:"admin"`
	LeagueMailOptOut bool        `json:"league_mail_opt_out"`
	LeaguePmOptOut   bool        `json:"league_pm_opt_out"`
	CarNumber        *string     `json:"car_number"`
	NickName         interface{} `json:"nick_name"`
}

func (api *LeagueApi) GetLeagueMembership(params LeagueMembershipParams) (*LeagueMembershipResponse, error) {
	paramsString, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	url := "/data/league/membership?" + paramsString.Encode()

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &LeagueMembershipResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
