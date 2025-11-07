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

func (api *ConstantsApi) Categories() (*categories.None, error) {
	return client.GetJson[categories.None](api.Client, "/data/constants/categories", nil)
}

func (api *ConstantsApi) Divisions() (*divisions.None, error) {
	return client.GetJson[divisions.None](api.Client, "/data/constants/divisions", nil)
}

func (api *ConstantsApi) EventTypes() (*event_types.None, error) {
	return client.GetJson[event_types.None](api.Client, "/data/constants/event_types", nil)
}
