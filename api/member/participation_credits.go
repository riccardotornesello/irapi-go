package member

import (
	"encoding/json"
)

type MemberParticipationCreditsResponse []struct {
	CustId               int    `json:"cust_id"`
	SeasonId             int    `json:"season_id"`
	SeriesId             int    `json:"series_id"`
	SeriesName           string `json:"series_name"`
	LicenseGroup         int    `json:"license_group"`
	LicenseGroupName     string `json:"license_group_name"`
	ParticipationCredits int    `json:"participation_credits"`
	MinWeeks             int    `json:"min_weeks"`
	Weeks                int    `json:"weeks"`
	EarnedCredits        int    `json:"earned_credits"`
	TotalCredits         int    `json:"total_credits"`
}

// Always the authenticated member.
func (api *MemberApi) GetMemberParticipationCredits() (*MemberParticipationCreditsResponse, error) {
	url := "/data/member/participation_credits"

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &MemberParticipationCreditsResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
