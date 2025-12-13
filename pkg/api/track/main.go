package track

import (
	"github.com/riccardotornesello/irapi-go/pkg/api/track/assets"
	"github.com/riccardotornesello/irapi-go/pkg/api/track/get"
	"github.com/riccardotornesello/irapi-go/pkg/client"
)

type TrackApi struct {
	client *client.ApiClient
}

func NewTrackApi(client *client.ApiClient) *TrackApi {
	return &TrackApi{
		client: client,
	}
}

func (api *TrackApi) Assets() (*assets.TrackAssetsResponse, error) {
	resp, err := client.GetJson[assets.TrackAssetsResponse](api.client, "/data/track/assets", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *TrackApi) Get() (*get.TrackGetResponse, error) {
	resp, err := client.GetJson[get.TrackGetResponse](api.client, "/data/track/get", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
