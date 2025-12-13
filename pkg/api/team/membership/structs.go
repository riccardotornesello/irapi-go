package membership

type TeamMembershipResponse []TeamMembershipResponseElement

type TeamMembershipResponseElement struct {
	TeamID      int64  `json:"team_id"`
	TeamName    string `json:"team_name"`
	Owner       bool   `json:"owner"`
	Admin       bool   `json:"admin"`
	DefaultTeam bool   `json:"default_team"`
}
