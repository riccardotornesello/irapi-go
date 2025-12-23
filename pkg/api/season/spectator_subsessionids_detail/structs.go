package spectator_subsessionids_detail

import (
	"github.com/riccardotornesello/irapi-go/pkg/types"
)

type SeasonSpectatorSubsessionidsDetailParams struct {
	EventTypes *[]int `json:"event_types,omitempty,comma"`
	SeasonIds  *[]int `json:"season_ids,omitempty,comma"`
}

type SeasonSpectatorSubsessionidsDetailResponse struct {
	Success     bool         `json:"success"`
	SeasonIDS   []int64      `json:"season_ids"`
	EventTypes  []int64      `json:"event_types"`
	Subsessions []Subsession `json:"subsessions"`
}

type Subsession struct {
	SubsessionID int64          `json:"subsession_id"`
	SessionID    int64          `json:"session_id"`
	SeasonID     int64          `json:"season_id"`
	StartTime    types.DateTime `json:"start_time"`
	RaceWeekNum  int64          `json:"race_week_num"`
	EventType    int64          `json:"event_type"`
}
