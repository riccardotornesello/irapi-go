package league

import (
	"encoding/json"
)

type LeagueMembershipResponse []LeagueMembershipResponseElement

type LeagueMembershipResponseElement struct {
	LeagueID         int64   `json:"league_id"`
	LeagueName       string  `json:"league_name"`
	Owner            bool    `json:"owner"`
	Admin            bool    `json:"admin"`
	LeagueMailOptOut bool    `json:"league_mail_opt_out"`
	LeaguePmOptOut   bool    `json:"league_pm_opt_out"`
	CarNumber        *string `json:"car_number"`
	NickName         *string `json:"nick_name"`
}

func (api *LeagueApi) Membership() (*LeagueMembershipResponse, error) {
	return api.GetJson[LeagueMembershipResponse]("/data/league/membership")
}
