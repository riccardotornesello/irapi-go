package constants

import (
	"github.com/riccardotornesello/irapi-go/pkg/api/constants/categories"
	"github.com/riccardotornesello/irapi-go/pkg/api/constants/divisions"
	"github.com/riccardotornesello/irapi-go/pkg/api/constants/event_types"
	"github.com/riccardotornesello/irapi-go/pkg/client"
)

type ConstantsApi struct {
	client *client.ApiClient
}

func NewConstantsApi(client *client.ApiClient) *ConstantsApi {
	return &ConstantsApi{
		client: client,
	}
}

func (api *ConstantsApi) Categories() (*categories.ConstantsCategoriesResponse, error) {
	resp, err := client.GetJson[categories.ConstantsCategoriesResponse](api.client, "/data/constants/categories", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *ConstantsApi) Divisions() (*divisions.ConstantsDivisionsResponse, error) {
	resp, err := client.GetJson[divisions.ConstantsDivisionsResponse](api.client, "/data/constants/divisions", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *ConstantsApi) EventTypes() (*event_types.ConstantsEventTypesResponse, error) {
	resp, err := client.GetJson[event_types.ConstantsEventTypesResponse](api.client, "/data/constants/event_types", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
