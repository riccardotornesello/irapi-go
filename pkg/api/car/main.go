package car

import (
	"github.com/riccardotornesello/irapi-go/pkg/api/car/assets"
	"github.com/riccardotornesello/irapi-go/pkg/api/car/get"
	"github.com/riccardotornesello/irapi-go/pkg/client"
)

type CarApi struct {
	client *client.ApiClient
}

func NewCarApi(client *client.ApiClient) *CarApi {
	return &CarApi{
		client: client,
	}
}

func (api *CarApi) Assets() (*assets.CarAssetsResponse, error) {
	resp, err := client.GetJson[assets.CarAssetsResponse](api.client, "/data/car/assets", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *CarApi) Get() (*get.CarGetResponse, error) {
	resp, err := client.GetJson[get.CarGetResponse](api.client, "/data/car/get", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
