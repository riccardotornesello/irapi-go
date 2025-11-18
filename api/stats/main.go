package stats

import (
	"github.com/riccardotornesello/irapi-go/api/stats/member_bests"
	"github.com/riccardotornesello/irapi-go/api/stats/member_career"
	"github.com/riccardotornesello/irapi-go/api/stats/member_recap"
	"github.com/riccardotornesello/irapi-go/api/stats/member_recent_races"
	"github.com/riccardotornesello/irapi-go/api/stats/member_summary"
	"github.com/riccardotornesello/irapi-go/api/stats/member_yearly"
	"github.com/riccardotornesello/irapi-go/client"
)

type StatsApi struct {
	Client *client.ApiClient
}

func (api *StatsApi) MemberBests(parameters *member_bests.StatsMemberBestsParams) (*member_bests.StatsMemberBestsResponse, error) {
	resp, err := client.GetJson[member_bests.StatsMemberBestsResponse](api.Client, "/data/stats/member_bests", parameters)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *StatsApi) MemberCareer(parameters *member_career.StatsMemberCareerParams) (*member_career.StatsMemberCareerResponse, error) {
	resp, err := client.GetJson[member_career.StatsMemberCareerResponse](api.Client, "/data/stats/member_career", parameters)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *StatsApi) MemberRecap(parameters *member_recap.StatsMemberRecapParams) (*member_recap.StatsMemberRecapResponse, error) {
	resp, err := client.GetJson[member_recap.StatsMemberRecapResponse](api.Client, "/data/stats/member_recap", parameters)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *StatsApi) MemberRecentRaces(parameters *member_recent_races.StatsMemberRecentRacesParams) (*member_recent_races.StatsMemberRecentRacesResponse, error) {
	resp, err := client.GetJson[member_recent_races.StatsMemberRecentRacesResponse](api.Client, "/data/stats/member_recent_races", parameters)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *StatsApi) MemberSummary(parameters *member_summary.StatsMemberSummaryParams) (*member_summary.StatsMemberSummaryResponse, error) {
	resp, err := client.GetJson[member_summary.StatsMemberSummaryResponse](api.Client, "/data/stats/member_summary", parameters)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *StatsApi) MemberYearly(parameters *member_yearly.StatsMemberYearlyParams) (*member_yearly.StatsMemberYearlyResponse, error) {
	resp, err := client.GetJson[member_yearly.StatsMemberYearlyResponse](api.Client, "/data/stats/member_yearly", parameters)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
