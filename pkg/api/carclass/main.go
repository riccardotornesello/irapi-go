package carclass

import (
	"github.com/riccardotornesello/irapi-go/pkg/api/carclass/get"
	"github.com/riccardotornesello/irapi-go/pkg/client"
)

type CarclassApi struct {
	client *client.ApiClient
}

func NewCarclassApi(client *client.ApiClient) *CarclassApi {
	return &CarclassApi{
		client: client,
	}
}

func (api *CarclassApi) Get() (*get.CarclassGetResponse, error) {
	resp, err := client.GetJson[get.CarclassGetResponse](api.client, "/data/carclass/get", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
