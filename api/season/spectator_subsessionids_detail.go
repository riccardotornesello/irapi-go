package season

import (
	"encoding/json"
)

type SeasonSpectatorSubsessionidsDetailResponse struct {
	Success     bool         `json:"success"`
	SeasonIDS   []int64      `json:"season_ids"`
	EventTypes  []int64      `json:"event_types"`
	Subsessions []Subsession `json:"subsessions"`
}

import "time"

type Subsession struct {
	SubsessionID int64     `json:"subsession_id"`
	SessionID    int64     `json:"session_id"`
	SeasonID     int64     `json:"season_id"`
	StartTime    time.Time `json:"start_time"`
	RaceWeekNum  int64     `json:"race_week_num"`
	EventType    int64     `json:"event_type"`
}


func (api *SeasonApi) SpectatorSubsessionidsDetail() (*SeasonSpectatorSubsessionidsDetailResponse, error) {
	return api.GetJson[SeasonSpectatorSubsessionidsDetailResponse]("/data/season/spectator_subsessionids_detail")
}