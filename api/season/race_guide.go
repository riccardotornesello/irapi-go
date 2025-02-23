package season

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/markphelps/optional"
)

type SeasonRaceGuideParams struct {
	From                *optional.String `url:"from,omitempty"`                   // ISO-8601 offset format. Defaults to the current time. Include sessions with start times up to 3 hours after this time. Times in the past will be rewritten to the current time.
	IncludeEndAfterFrom *optional.Bool   `url:"include_end_after_from,omitempty"` // Include sessions which start before 'from' but end after.
}

type SeasonRaceGuideResponse struct {
	Subscribed bool `json:"subscribed"`
	Sessions   []struct {
		SeasonId     int    `json:"season_id"`
		StartTime    string `json:"start_time"`
		SuperSession bool   `json:"super_session"`
		SeriesId     int    `json:"series_id"`
		RaceWeekNum  int    `json:"race_week_num"`
		EndTime      string `json:"end_time"`
		SessionId    int    `json:"session_id"`
		EntryCount   int    `json:"entry_count"`
	} `json:"sessions"`
	BlockBeginTime string `json:"block_begin_time"`
	BlockEndTime   string `json:"block_end_time"`
	Success        bool   `json:"success"`
}

func (api *SeasonApi) GetSeasonRaceGuide(params SeasonRaceGuideParams) (*SeasonRaceGuideResponse, error) {
	paramsString, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	url := "/data/season/race_guide?" + paramsString.Encode()

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &SeasonRaceGuideResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
