package track

import (
	"github.com/riccardotornesello/irapi-go/api/track/assets"
	"github.com/riccardotornesello/irapi-go/api/track/get"
	"github.com/riccardotornesello/irapi-go/client"
)

type TrackApi struct {
	Client *client.ApiClient
}

func (api *TrackApi) Assets() (*assets.TrackAssetsResponse, error) {
	resp, err := client.GetJson[assets.TrackAssetsResponse](api.Client, "/data/track/assets", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *TrackApi) Get() (*get.TrackGetResponse, error) {
	resp, err := client.GetJson[get.TrackGetResponse](api.Client, "/data/track/get", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
