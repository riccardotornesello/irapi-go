package league

import (
	"encoding/json"
)

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

func (api *LeagueApi) GetPointsSystems() (*LeagueGetPointsSystemsResponse, error) {
	return api.GetJson[LeagueGetPointsSystemsResponse]("/data/league/get_points_systems")
}
