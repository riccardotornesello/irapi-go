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
	Client *client.ApiClient
}

func (api *LookupApi) Countries() (*countries.LookupCountriesResponse, error) {
	resp, err := client.GetJson[countries.LookupCountriesResponse](api.Client, "/data/lookup/countries", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *LookupApi) Drivers(parameters *drivers.LookupDriversParams) (*drivers.LookupDriversResponse, error) {
	resp, err := client.GetJson[drivers.LookupDriversResponse](api.Client, "/data/lookup/drivers", parameters)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *LookupApi) Flairs() (*flairs.LookupFlairsResponse, error) {
	resp, err := client.GetJson[flairs.LookupFlairsResponse](api.Client, "/data/lookup/flairs", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *LookupApi) Get() (*get.LookupGetResponse, error) {
	resp, err := client.GetJson[get.LookupGetResponse](api.Client, "/data/lookup/get", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *LookupApi) Licenses() (*licenses.LookupLicensesResponse, error) {
	resp, err := client.GetJson[licenses.LookupLicensesResponse](api.Client, "/data/lookup/licenses", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
