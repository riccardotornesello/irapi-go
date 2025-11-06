package team

import (
	"encoding/json"
)

type TeamMembershipResponse []TeamMembershipResponseElement

type TeamMembershipResponseElement struct {
	TeamID      int64  `json:"team_id"`
	TeamName    string `json:"team_name"`
	Owner       bool   `json:"owner"`
	Admin       bool   `json:"admin"`
	DefaultTeam bool   `json:"default_team"`
}

func (api *TeamApi) Membership() (*TeamMembershipResponse, error) {
	return api.GetJson[TeamMembershipResponse]("/data/team/membership")
}
