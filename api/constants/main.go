package constants

import (
	"github.com/riccardotornesello/irapi-go/api/constants/categories"
	"github.com/riccardotornesello/irapi-go/api/constants/divisions"
	"github.com/riccardotornesello/irapi-go/api/constants/event_types"
	"github.com/riccardotornesello/irapi-go/client"
)

type ConstantsApi struct {
	Client *client.ApiClient
}

func (api *ConstantsApi) Categories() (*categories.ConstantsCategoriesResponse, error) {
	resp, err := client.GetJson[categories.ConstantsCategoriesResponse](api.Client, "/data/constants/categories", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *ConstantsApi) Divisions() (*divisions.ConstantsDivisionsResponse, error) {
	resp, err := client.GetJson[divisions.ConstantsDivisionsResponse](api.Client, "/data/constants/divisions", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *ConstantsApi) EventTypes() (*event_types.ConstantsEventTypesResponse, error) {
	resp, err := client.GetJson[event_types.ConstantsEventTypesResponse](api.Client, "/data/constants/event_types", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
