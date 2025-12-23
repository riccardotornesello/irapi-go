package roster

type LeagueRosterParams struct {
	LeagueId        int   `json:"league_id,comma"`
	IncludeLicenses *bool `json:"include_licenses,omitempty,comma"`
}

type LeagueRosterResponse struct {
	Type    string `json:"type"`
	Data    Data   `json:"data"`
	DataURL string `json:"data_url"`
}

type Data struct {
	Subscribed  bool  `json:"subscribed"`
	Success     bool  `json:"success"`
	RosterCount int64 `json:"roster_count"`
	LeagueID    int64 `json:"league_id"`
}
