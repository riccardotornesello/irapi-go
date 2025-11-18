package season_standings

type LeagueSeasonStandingsParams struct {
	LeagueId   int `url:"league_id,omitempty,comma"`
	SeasonId   int `url:"season_id,omitempty,comma"`
	CarClassId int `url:"car_class_id,omitempty,comma"`
	CarId      int `url:"car_id,omitempty,comma"`
}

type LeagueSeasonStandingsResponse struct {
	CarClassID int64     `json:"car_class_id"`
	Success    bool      `json:"success"`
	SeasonID   int64     `json:"season_id"`
	CarID      int64     `json:"car_id"`
	Standings  Standings `json:"standings"`
	LeagueID   int64     `json:"league_id"`
}

type Standings struct {
	DriverStandings       []DriverStanding `json:"driver_standings"`
	TeamStandings         []interface{}    `json:"team_standings"`
	DriverStandingsCSVURL string           `json:"driver_standings_csv_url"`
	TeamStandingsCSVURL   string           `json:"team_standings_csv_url"`
}

type DriverStanding struct {
	Rownum              int64   `json:"rownum"`
	Position            int64   `json:"position"`
	Driver              Driver  `json:"driver"`
	CarNumber           *string `json:"car_number"`
	DriverNickname      *string `json:"driver_nickname"`
	WINS                int64   `json:"wins"`
	AverageStart        int64   `json:"average_start"`
	AverageFinish       int64   `json:"average_finish"`
	BasePoints          int64   `json:"base_points"`
	NegativeAdjustments int64   `json:"negative_adjustments"`
	PositiveAdjustments int64   `json:"positive_adjustments"`
	TotalAdjustments    int64   `json:"total_adjustments"`
	TotalPoints         int64   `json:"total_points"`
}

type Driver struct {
	CustID      int64  `json:"cust_id"`
	DisplayName string `json:"display_name"`
	Helmet      Helmet `json:"helmet"`
}

type Helmet struct {
	Pattern    int64  `json:"pattern"`
	Color1     string `json:"color1"`
	Color2     string `json:"color2"`
	Color3     string `json:"color3"`
	FaceType   int64  `json:"face_type"`
	HelmetType int64  `json:"helmet_type"`
}
