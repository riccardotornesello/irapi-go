package time_attack

import (
	"github.com/riccardotornesello/irapi-go/api/time_attack/member_season_results"
	"github.com/riccardotornesello/irapi-go/client"
)

type TimeAttackApi struct {
	Client *client.ApiClient
}

func (api *TimeAttackApi) MemberSeasonResults(parameters *member_season_results.TimeAttackMemberSeasonResultsParams) (*member_season_results.None, error) {
	return client.GetJson[member_season_results.None](api.Client, "/data/time_attack/member_season_results", parameters)
}
