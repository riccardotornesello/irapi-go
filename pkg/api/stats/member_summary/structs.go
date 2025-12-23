package member_summary

type StatsMemberSummaryParams struct {
	CustId *int `json:"cust_id,omitempty,comma"`
}

type StatsMemberSummaryResponse struct {
	ThisYear ThisYear `json:"this_year"`
	CustID   int64    `json:"cust_id"`
}

type ThisYear struct {
	NumOfficialSessions int64 `json:"num_official_sessions"`
	NumLeagueSessions   int64 `json:"num_league_sessions"`
	NumOfficialWINS     int64 `json:"num_official_wins"`
	NumLeagueWINS       int64 `json:"num_league_wins"`
}
