package car

import (
	"github.com/riccardotornesello/irapi-go/api/car/assets"
	"github.com/riccardotornesello/irapi-go/api/car/get"
	"github.com/riccardotornesello/irapi-go/client"
)

type CarApi struct {
	Client *client.ApiClient
}

func (api *CarApi) Assets() (*assets.CarAssetsResponse, error) {
	resp, err := client.GetJson[assets.CarAssetsResponse](api.Client, "/data/car/assets", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *CarApi) Get() (*get.CarGetResponse, error) {
	resp, err := client.GetJson[get.CarGetResponse](api.Client, "/data/car/get", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
