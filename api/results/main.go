package results

import (
	"github.com/riccardotornesello/irapi-go/api/results/get"
	"github.com/riccardotornesello/irapi-go/client"
)

type ResultsApi struct {
	Client *client.ApiClient
}

func (api *ResultsApi) Get(parameters *get.ResultsGetParams) (*get.ResultsGetResponse, error) {
	return client.GetJson[get.ResultsGetResponse](api.Client, "/data/results/get", parameters)
}
