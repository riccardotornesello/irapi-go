package lookup

import (
	"encoding/json"
)

type LookupCountriesResponse []LookupCountriesResponseElement

type LookupCountriesResponseElement struct {
	CountryName string `json:"country_name"`
	CountryCode string `json:"country_code"`
}

func (api *LookupApi) Countries() (*LookupCountriesResponse, error) {
	return api.GetJson[LookupCountriesResponse]("/data/lookup/countries")
}
