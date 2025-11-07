package session

import (
	"github.com/riccardotornesello/irapi-go/api/session/reg_drivers_list"
	"github.com/riccardotornesello/irapi-go/client"
)

type SessionApi struct {
	Client *client.ApiClient
}

func (api *SessionApi) RegDriversList(parameters *reg_drivers_list.SessionRegDriversListParams) (*reg_drivers_list.None, error) {
	return client.GetJson[reg_drivers_list.None](api.Client, "/data/session/reg_drivers_list", parameters)
}
