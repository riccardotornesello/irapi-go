package member_recent_races

import (
	"github.com/riccardotornesello/irapi-go/pkg/types"
)

type StatsMemberRecentRacesParams struct {
	CustId *int `json:"cust_id,omitempty,comma"`
}

type StatsMemberRecentRacesResponse struct {
	Races  []Race `json:"races"`
	CustID int64  `json:"cust_id"`
}

type Race struct {
	SeasonID           int64          `json:"season_id"`
	SeriesID           int64          `json:"series_id"`
	SeriesName         string         `json:"series_name"`
	CarID              int64          `json:"car_id"`
	CarClassID         int64          `json:"car_class_id"`
	Livery             Livery         `json:"livery"`
	LicenseLevel       int64          `json:"license_level"`
	SessionStartTime   types.DateTime `json:"session_start_time"`
	WinnerGroupID      int64          `json:"winner_group_id"`
	WinnerName         string         `json:"winner_name"`
	WinnerHelmet       WinnerHelmet   `json:"winner_helmet"`
	WinnerLicenseLevel int64          `json:"winner_license_level"`
	StartPosition      int64          `json:"start_position"`
	FinishPosition     int64          `json:"finish_position"`
	QualifyingTime     int64          `json:"qualifying_time"`
	Laps               int64          `json:"laps"`
	LapsLED            int64          `json:"laps_led"`
	Incidents          int64          `json:"incidents"`
	Points             int64          `json:"points"`
	StrengthOfField    int64          `json:"strength_of_field"`
	SubsessionID       int64          `json:"subsession_id"`
	OldSubLevel        int64          `json:"old_sub_level"`
	NewSubLevel        int64          `json:"new_sub_level"`
	OldiRating         int64          `json:"oldi_rating"`
	NewiRating         int64          `json:"newi_rating"`
	Track              Track          `json:"track"`
	DropRace           bool           `json:"drop_race"`
	SeasonYear         int64          `json:"season_year"`
	SeasonQuarter      int64          `json:"season_quarter"`
	RaceWeekNum        int64          `json:"race_week_num"`
}

type Livery struct {
	CarID   int64  `json:"car_id"`
	Pattern int64  `json:"pattern"`
	Color1  string `json:"color1"`
	Color2  string `json:"color2"`
	Color3  string `json:"color3"`
}

type Track struct {
	TrackID   int64  `json:"track_id"`
	TrackName string `json:"track_name"`
}

type WinnerHelmet struct {
	Pattern    int64  `json:"pattern"`
	Color1     string `json:"color1"`
	Color2     string `json:"color2"`
	Color3     string `json:"color3"`
	FaceType   int64  `json:"face_type"`
	HelmetType int64  `json:"helmet_type"`
}
