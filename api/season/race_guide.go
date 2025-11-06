package season

import (
	"encoding/json"
)

import "time"

type SeasonRaceGuideResponse struct {
	Subscribed     bool      `json:"subscribed"`
	Sessions       []Session `json:"sessions"`
	BlockBeginTime time.Time `json:"block_begin_time"`
	BlockEndTime   time.Time `json:"block_end_time"`
	Success        bool      `json:"success"`
}

type Session struct {
	SeasonID     int64     `json:"season_id"`
	StartTime    time.Time `json:"start_time"`
	SuperSession bool      `json:"super_session"`
	SeriesID     int64     `json:"series_id"`
	RaceWeekNum  int64     `json:"race_week_num"`
	EndTime      time.Time `json:"end_time"`
	SessionID    *int64    `json:"session_id,omitempty"`
	EntryCount   int64     `json:"entry_count"`
}

func (api *SeasonApi) RaceGuide() (*SeasonRaceGuideResponse, error) {
	return api.GetJson[SeasonRaceGuideResponse]("/data/season/race_guide")
}
