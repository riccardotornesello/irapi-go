package constants

import (
	"encoding/json"
)

type ConstantsEventTypesResponse []struct {
	Label string `json:"label"`
	Value int    `json:"value"`
}

// Constant; returned directly as an array of objects
func (api *ConstantsApi) GetConstantsEventTypes() (*ConstantsEventTypesResponse, error) {
	url := "https://members-ng.iracing.com/data/constants/event_types"

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &ConstantsEventTypesResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
