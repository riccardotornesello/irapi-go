package results

import (
	"github.com/riccardotornesello/irapi-go/api/results/event_log"
	"github.com/riccardotornesello/irapi-go/api/results/get"
	"github.com/riccardotornesello/irapi-go/api/results/lap_chart_data"
	"github.com/riccardotornesello/irapi-go/api/results/lap_data"
	"github.com/riccardotornesello/irapi-go/api/results/search_hosted"
	"github.com/riccardotornesello/irapi-go/api/results/search_series"
	"github.com/riccardotornesello/irapi-go/api/results/season_results"
	"github.com/riccardotornesello/irapi-go/client"
)

type ResultsApi struct {
	Client *client.ApiClient
}

func (api *ResultsApi) Get(parameters *get.ResultsGetParams) (*get.ResultsGetResponse, error) {
	return client.GetJson[get.ResultsGetResponse](api.Client, "/data/results/get", parameters)
}

func (api *ResultsApi) EventLog(parameters *event_log.ResultsEventLogParams) (*event_log.None, error) {
	return client.GetJson[event_log.None](api.Client, "/data/results/event_log", parameters)
}

func (api *ResultsApi) LapChartData(parameters *lap_chart_data.ResultsLapChartDataParams) (*lap_chart_data.None, error) {
	return client.GetJson[lap_chart_data.None](api.Client, "/data/results/lap_chart_data", parameters)
}

func (api *ResultsApi) LapData(parameters *lap_data.ResultsLapDataParams) (*lap_data.None, error) {
	return client.GetJson[lap_data.None](api.Client, "/data/results/lap_data", parameters)
}

func (api *ResultsApi) SearchHosted(parameters *search_hosted.ResultsSearchHostedParams) (*search_hosted.None, error) {
	return client.GetJson[search_hosted.None](api.Client, "/data/results/search_hosted", parameters)
}

func (api *ResultsApi) SearchSeries(parameters *search_series.ResultsSearchSeriesParams) (*search_series.None, error) {
	return client.GetJson[search_series.None](api.Client, "/data/results/search_series", parameters)
}

func (api *ResultsApi) SeasonResults(parameters *season_results.ResultsSeasonResultsParams) (*season_results.None, error) {
	return client.GetJson[season_results.None](api.Client, "/data/results/season_results", parameters)
}
