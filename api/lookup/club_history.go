package lookup

import (
	"encoding/json"

	"github.com/google/go-querystring/query"
)

type LookupClubHistoryParams struct {
	SeasonYear    int `url:"season_year,omitempty"`
	SeasonQuarter int `url:"season_quarter,omitempty"`
}

type LookupClubHistoryResponse []struct {
	ClubId        int    `json:"club_id"`
	ClubName      string `json:"club_name"`
	SeasonYear    int    `json:"season_year"`
	SeasonQuarter int    `json:"season_quarter"`
	Region        string `json:"region"`
}

// Returns an earlier history if requested quarter does not have a club history.
func (api *LookupApi) GetLookupClubHistory(params LookupClubHistoryParams) (*LookupClubHistoryResponse, error) {
	paramsString, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	url := "https://members-ng.iracing.com/data/lookup/club_history?" + paramsString.Encode()

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &LookupClubHistoryResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
