package spectator_subsessionids

import ()

type SeasonSpectatorSubsessionidsParams struct {
	EventTypes []int `url:"event_types,omitempty,comma"`
}

type SeasonSpectatorSubsessionidsResponse struct {
	EventTypes    []int64 `json:"event_types"`
	Success       bool    `json:"success"`
	SubsessionIDS []int64 `json:"subsession_ids"`
}
