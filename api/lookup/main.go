package lookup

import (
	"github.com/riccardotornesello/irapi-go/api/lookup/countries"
	"github.com/riccardotornesello/irapi-go/api/lookup/flairs"
	"github.com/riccardotornesello/irapi-go/api/lookup/get"
	"github.com/riccardotornesello/irapi-go/api/lookup/licenses"
	"github.com/riccardotornesello/irapi-go/client"
)

type LookupApi struct {
	Client *client.ApiClient
}

func (api *LookupApi) Countries() (*countries.LookupCountriesResponse, error) {
	return client.GetJson[countries.LookupCountriesResponse](api.Client, "/data/lookup/countries", nil)
}

func (api *LookupApi) Flairs() (*flairs.LookupFlairsResponse, error) {
	return client.GetJson[flairs.LookupFlairsResponse](api.Client, "/data/lookup/flairs", nil)
}

func (api *LookupApi) Get() (*get.LookupGetResponse, error) {
	return client.GetJson[get.LookupGetResponse](api.Client, "/data/lookup/get", nil)
}

func (api *LookupApi) Licenses() (*licenses.LookupLicensesResponse, error) {
	return client.GetJson[licenses.LookupLicensesResponse](api.Client, "/data/lookup/licenses", nil)
}
