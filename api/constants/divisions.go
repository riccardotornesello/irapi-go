package constants

import (
	"encoding/json"
)

type ConstantsDivisionsResponse []struct {
	Label string `json:"label"`
	Value int    `json:"value"`
}

func (api *ConstantsApi) GetConstantsDivisions() (*ConstantsDivisionsResponse, error) {
	url := "https://members-ng.iracing.com/data/constants/divisions"

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &ConstantsDivisionsResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
