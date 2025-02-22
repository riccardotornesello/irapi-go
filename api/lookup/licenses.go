package lookup

import (
	"encoding/json"
)

type LookupLicensesResponse []struct {
	LicenseGroup         int    `json:"license_group"`
	GroupName            string `json:"group_name"`
	MinNumRaces          *int   `json:"min_num_races"`
	ParticipationCredits int    `json:"participation_credits"`
	MinSrToFastTrack     *int   `json:"min_sr_to_fast_track"`
	Levels               []struct {
		LicenseId     int    `json:"license_id"`
		LicenseGroup  int    `json:"license_group"`
		License       string `json:"license"`
		ShortName     string `json:"short_name"`
		LicenseLetter string `json:"license_letter"`
		Color         string `json:"color"`
	} `json:"levels"`
	MinNumTt *int `json:"min_num_tt"`
}

func (api *LookupApi) GetLookupLicenses() (*LookupLicensesResponse, error) {
	url := "/data/lookup/licenses"

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &LookupLicensesResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
