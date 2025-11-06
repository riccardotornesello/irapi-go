package season

import (
	"encoding/json"
)

type SeasonSpectatorSubsessionidsResponse struct {
	EventTypes    []int64 `json:"event_types"`
	Success       bool    `json:"success"`
	SubsessionIDS []int64 `json:"subsession_ids"`
}

func (api *SeasonApi) SpectatorSubsessionids() (*SeasonSpectatorSubsessionidsResponse, error) {
	return api.GetJson[SeasonSpectatorSubsessionidsResponse]("/data/season/spectator_subsessionids")
}
