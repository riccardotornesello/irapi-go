package lookup

import (
	"github.com/riccardotornesello/irapi-go/api/lookup/countries"
	"github.com/riccardotornesello/irapi-go/api/lookup/drivers"
	"github.com/riccardotornesello/irapi-go/api/lookup/flairs"
	"github.com/riccardotornesello/irapi-go/api/lookup/get"
	"github.com/riccardotornesello/irapi-go/api/lookup/licenses"
	"github.com/riccardotornesello/irapi-go/client"
)

type LookupApi struct {
	client *client.ApiClient
}

func NewLookupApi(client *client.ApiClient) *LookupApi {
	return &LookupApi{
		client: client,
	}
}

func (api *LookupApi) Countries() (*countries.LookupCountriesResponse, error) {
	resp, err := client.GetJson[countries.LookupCountriesResponse](api.client, "/data/lookup/countries", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *LookupApi) Drivers(parameters *drivers.LookupDriversParams) (*drivers.LookupDriversResponse, error) {
	resp, err := client.GetJson[drivers.LookupDriversResponse](api.client, "/data/lookup/drivers", parameters)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *LookupApi) Flairs() (*flairs.LookupFlairsResponse, error) {
	resp, err := client.GetJson[flairs.LookupFlairsResponse](api.client, "/data/lookup/flairs", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *LookupApi) Get() (*get.LookupGetResponse, error) {
	resp, err := client.GetJson[get.LookupGetResponse](api.client, "/data/lookup/get", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *LookupApi) Licenses() (*licenses.LookupLicensesResponse, error) {
	resp, err := client.GetJson[licenses.LookupLicensesResponse](api.client, "/data/lookup/licenses", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
