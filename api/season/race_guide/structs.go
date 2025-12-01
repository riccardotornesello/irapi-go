package race_guide

import (
	"github.com/riccardotornesello/irapi-go/pkg/types"
)

type SeasonRaceGuideParams struct {
	From                string `url:"from,omitempty,comma"`
	IncludeEndAfterFrom bool   `url:"include_end_after_from,omitempty,comma"`
}

type SeasonRaceGuideResponse struct {
	Subscribed     bool           `json:"subscribed"`
	Sessions       []Session      `json:"sessions"`
	BlockBeginTime types.DateTime `json:"block_begin_time"`
	BlockEndTime   types.DateTime `json:"block_end_time"`
	Success        bool           `json:"success"`
}

type Session struct {
	SeasonID     int64          `json:"season_id"`
	StartTime    types.DateTime `json:"start_time"`
	SuperSession bool           `json:"super_session"`
	SeriesID     int64          `json:"series_id"`
	RaceWeekNum  int64          `json:"race_week_num"`
	EndTime      types.DateTime `json:"end_time"`
	SessionID    *int64         `json:"session_id,omitempty"`
	EntryCount   int64          `json:"entry_count"`
}
