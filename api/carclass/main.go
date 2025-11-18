package carclass

import (
	"github.com/riccardotornesello/irapi-go/api/carclass/get"
	"github.com/riccardotornesello/irapi-go/client"
)

type CarclassApi struct {
	Client *client.ApiClient
}

func (api *CarclassApi) Get() (*get.CarclassGetResponse, error) {
	resp, err := client.GetJson[get.CarclassGetResponse](api.Client, "/data/carclass/get", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
