package results

import (
	"github.com/riccardotornesello/irapi-go/api/results/get"
	"github.com/riccardotornesello/irapi-go/api/results/lap_data"
	"github.com/riccardotornesello/irapi-go/client"
)

type ResultsApi struct {
	Client *client.ApiClient
}

func (api *ResultsApi) Get(parameters *get.ResultsGetParams) (*get.ResultsGetResponse, error) {
	resp, err := client.GetJson[get.ResultsGetResponse](api.Client, "/data/results/get", parameters)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *ResultsApi) LapData(parameters *lap_data.ResultsLapDataParams) (*lap_data.ResultsLapDataResponse, error) {
	resp, err := client.GetJson[lap_data.ResultsLapDataResponse](api.Client, "/data/results/lap_data", parameters)
	if err != nil {
		return nil, err
	}

	chunks, err := client.GetChunks[lap_data.ResultsLapDataResponseChunk](&resp.ChunkInfo)
	if err != nil {
		return nil, err
	}

	resp.Chunks = chunks

	return resp, nil
}
