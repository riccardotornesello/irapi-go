package seasons

import ()

type LeagueSeasonsParams struct {
	LeagueId int  `url:"league_id,omitempty,comma"`
	Retired  bool `url:"retired,omitempty,comma"`
}

type LeagueSeasonsResponse struct {
	Subscribed bool     `json:"subscribed"`
	Seasons    []Season `json:"seasons"`
	Success    bool     `json:"success"`
	Retired    bool     `json:"retired"`
	LeagueID   int64    `json:"league_id"`
}

type Season struct {
	LeagueID                int64            `json:"league_id"`
	SeasonID                int64            `json:"season_id"`
	PointsSystemID          int64            `json:"points_system_id"`
	SeasonName              string           `json:"season_name"`
	Active                  bool             `json:"active"`
	Hidden                  bool             `json:"hidden"`
	NumDrops                int64            `json:"num_drops"`
	NoDropsOnOrAfterRaceNum int64            `json:"no_drops_on_or_after_race_num"`
	PointsCars              []PointsCar      `json:"points_cars"`
	DriverPointsCarClasses  []PointsCarClass `json:"driver_points_car_classes"`
	TeamPointsCarClasses    []PointsCarClass `json:"team_points_car_classes"`
	PointsSystemName        string           `json:"points_system_name"`
	PointsSystemDesc        string           `json:"points_system_desc"`
}

type PointsCarClass struct {
	CarClassID  int64       `json:"car_class_id"`
	Name        string      `json:"name"`
	CarsInClass []PointsCar `json:"cars_in_class"`
}

type PointsCar struct {
	CarID   int64  `json:"car_id"`
	CarName string `json:"car_name"`
}
