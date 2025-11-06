package lookup

import (
	"encoding/json"
)

type LookupFlairsResponse struct {
	Flairs  []Flair `json:"flairs"`
	Success bool    `json:"success"`
}

type Flair struct {
	FlairID        int64   `json:"flair_id"`
	FlairName      string  `json:"flair_name"`
	Seq            int64   `json:"seq"`
	FlairShortname *string `json:"flair_shortname,omitempty"`
	CountryCode    *string `json:"country_code,omitempty"`
}

func (api *LookupApi) Flairs() (*LookupFlairsResponse, error) {
	return api.GetJson[LookupFlairsResponse]("/data/lookup/flairs")
}
