package series

import (
	"github.com/riccardotornesello/irapi-go/api/series/assets"
	"github.com/riccardotornesello/irapi-go/api/series/get"
	"github.com/riccardotornesello/irapi-go/api/series/season_list"
	"github.com/riccardotornesello/irapi-go/api/series/seasons"
	"github.com/riccardotornesello/irapi-go/api/series/stats_series"
	"github.com/riccardotornesello/irapi-go/client"
)

type SeriesApi struct {
	client *client.ApiClient
}

func NewSeriesApi(client *client.ApiClient) *SeriesApi {
	return &SeriesApi{
		client: client,
	}
}

func (api *SeriesApi) Assets() (*assets.SeriesAssetsResponse, error) {
	resp, err := client.GetJson[assets.SeriesAssetsResponse](api.client, "/data/series/assets", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *SeriesApi) Get() (*get.SeriesGetResponse, error) {
	resp, err := client.GetJson[get.SeriesGetResponse](api.client, "/data/series/get", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *SeriesApi) Seasons(parameters *seasons.SeriesSeasonsParams) (*seasons.SeriesSeasonsResponse, error) {
	resp, err := client.GetJson[seasons.SeriesSeasonsResponse](api.client, "/data/series/seasons", parameters)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *SeriesApi) SeasonList(parameters *season_list.SeriesSeasonListParams) (*season_list.SeriesSeasonListResponse, error) {
	resp, err := client.GetJson[season_list.SeriesSeasonListResponse](api.client, "/data/series/season_list", parameters)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *SeriesApi) StatsSeries() (*stats_series.SeriesStatsSeriesResponse, error) {
	resp, err := client.GetJson[stats_series.SeriesStatsSeriesResponse](api.client, "/data/series/stats_series", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
