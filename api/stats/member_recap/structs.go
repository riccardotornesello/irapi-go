package member_recap

type StatsMemberRecapParams struct {
	CustId *int `url:"cust_id,omitempty,comma"`
	Year   *int `url:"year,omitempty,comma"`
	Season *int `url:"season,omitempty,comma"`
}

type StatsMemberRecapResponse struct {
	Year    int64       `json:"year"`
	Stats   Stats       `json:"stats"`
	Success bool        `json:"success"`
	Season  interface{} `json:"season"`
	CustID  int64       `json:"cust_id"`
}

type Stats struct {
	Starts            int64         `json:"starts"`
	WINS              int64         `json:"wins"`
	Top5              int64         `json:"top5"`
	AvgStartPosition  int64         `json:"avg_start_position"`
	AvgFinishPosition int64         `json:"avg_finish_position"`
	Laps              int64         `json:"laps"`
	LapsLED           int64         `json:"laps_led"`
	FavoriteCar       FavoriteCar   `json:"favorite_car"`
	FavoriteTrack     FavoriteTrack `json:"favorite_track"`
}

type FavoriteCar struct {
	CarID    int64  `json:"car_id"`
	CarName  string `json:"car_name"`
	CarImage string `json:"car_image"`
}

type FavoriteTrack struct {
	ConfigName string `json:"config_name"`
	TrackID    int64  `json:"track_id"`
	TrackLogo  string `json:"track_logo"`
	TrackName  string `json:"track_name"`
}
