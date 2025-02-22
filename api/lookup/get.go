package lookup

import (
	"encoding/json"
)

type LookupGetResponse []interface{}

// ?weather=weather_wind_speed_units&weather=weather_wind_speed_max&weather=weather_wind_speed_min&licenselevels=licenselevels
func (api *LookupApi) GetLookup() (*LookupGetResponse, error) {
	url := "https://members-ng.iracing.com/data/lookup/get"

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &LookupGetResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
