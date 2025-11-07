package carclass

import (
	"github.com/riccardotornesello/irapi-go/api/carclass/get"
	"github.com/riccardotornesello/irapi-go/client"
)

type CarclassApi struct {
	Client *client.ApiClient
}

func (api *CarclassApi) Get() (*get.CarclassGetResponse, error) {
	return client.GetJson[get.CarclassGetResponse](api.Client, "/data/carclass/get", nil)
}
