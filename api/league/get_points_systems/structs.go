package get_points_systems

type LeagueGetPointsSystemsParams struct {
	LeagueId int `url:"league_id,omitempty,comma"`
	SeasonId int `url:"season_id,omitempty,comma"`
}

type LeagueGetPointsSystemsResponse struct {
	Subscribed    bool           `json:"subscribed"`
	Success       bool           `json:"success"`
	PointsSystems []PointsSystem `json:"points_systems"`
	LeagueID      int64          `json:"league_id"`
}

type PointsSystem struct {
	PointsSystemID int64  `json:"points_system_id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	LeagueID       int64  `json:"league_id"`
	Retired        bool   `json:"retired"`
	IracingSystem  bool   `json:"iracing_system"`
}
