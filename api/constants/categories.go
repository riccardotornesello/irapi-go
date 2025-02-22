package constants

import (
	"encoding/json"
)

type ConstantsCategoriesResponse []struct {
	Label string `json:"label"`
	Value int    `json:"value"`
}

// Constant; returned directly as an array of objects
func (api *ConstantsApi) GetConstantsCategories() (*ConstantsCategoriesResponse, error) {
	url := "https://members-ng.iracing.com/data/constants/categories"

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &ConstantsCategoriesResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
