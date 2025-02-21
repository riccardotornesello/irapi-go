package lookup

import (
	"encoding/json"
)

type LookupCountriesResponse []struct {
	CountryName string `json:"country_name"`
	CountryCode string `json:"country_code"`
}

func (api *LookupApi) GetLookupCountries() (*LookupCountriesResponse, error) {
	url := "https://members-ng.iracing.com/data/lookup/countries"

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &LookupCountriesResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
