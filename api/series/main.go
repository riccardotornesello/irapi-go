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
	Client *client.ApiClient
}

func (api *SeriesApi) Assets() (*assets.SeriesAssetsResponse, error) {
	return client.GetJson[assets.SeriesAssetsResponse](api.Client, "/data/series/assets", nil)
}

func (api *SeriesApi) Get() (*get.SeriesGetResponse, error) {
	return client.GetJson[get.SeriesGetResponse](api.Client, "/data/series/get", nil)
}

func (api *SeriesApi) Seasons(parameters *seasons.SeriesSeasonsParams) (*seasons.SeriesSeasonsResponse, error) {
	return client.GetJson[seasons.SeriesSeasonsResponse](api.Client, "/data/series/seasons", parameters)
}

func (api *SeriesApi) SeasonList(parameters *season_list.SeriesSeasonListParams) (*season_list.SeriesSeasonListResponse, error) {
	return client.GetJson[season_list.SeriesSeasonListResponse](api.Client, "/data/series/season_list", parameters)
}

func (api *SeriesApi) StatsSeries() (*stats_series.SeriesStatsSeriesResponse, error) {
	return client.GetJson[stats_series.SeriesStatsSeriesResponse](api.Client, "/data/series/stats_series", nil)
}
