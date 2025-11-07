package participation_credits

import ()

type MemberParticipationCreditsResponse []MemberParticipationCreditsResponseElement

type MemberParticipationCreditsResponseElement struct {
	CustID               int64  `json:"cust_id"`
	SeasonID             int64  `json:"season_id"`
	SeriesID             int64  `json:"series_id"`
	SeriesName           string `json:"series_name"`
	LicenseGroup         int64  `json:"license_group"`
	LicenseGroupName     string `json:"license_group_name"`
	ParticipationCredits int64  `json:"participation_credits"`
	MinWeeks             int64  `json:"min_weeks"`
	Weeks                int64  `json:"weeks"`
	EarnedCredits        int64  `json:"earned_credits"`
	TotalCredits         int64  `json:"total_credits"`
}
