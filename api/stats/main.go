package stats

import (
	"github.com/riccardotornesello/irapi-go/api/stats/member_bests"
	"github.com/riccardotornesello/irapi-go/api/stats/member_career"
	"github.com/riccardotornesello/irapi-go/api/stats/member_division"
	"github.com/riccardotornesello/irapi-go/api/stats/member_recap"
	"github.com/riccardotornesello/irapi-go/api/stats/member_recent_races"
	"github.com/riccardotornesello/irapi-go/api/stats/member_summary"
	"github.com/riccardotornesello/irapi-go/api/stats/member_yearly"
	"github.com/riccardotornesello/irapi-go/api/stats/season_driver_standings"
	"github.com/riccardotornesello/irapi-go/api/stats/season_qualify_results"
	"github.com/riccardotornesello/irapi-go/api/stats/season_supersession_standings"
	"github.com/riccardotornesello/irapi-go/api/stats/season_team_standings"
	"github.com/riccardotornesello/irapi-go/api/stats/season_tt_results"
	"github.com/riccardotornesello/irapi-go/api/stats/season_tt_standings"
	"github.com/riccardotornesello/irapi-go/api/stats/world_records"
	"github.com/riccardotornesello/irapi-go/client"
)

type StatsApi struct {
	Client *client.ApiClient
}

func (api *StatsApi) MemberBests(parameters *member_bests.StatsMemberBestsParams) (*member_bests.StatsMemberBestsResponse, error) {
	return client.GetJson[member_bests.StatsMemberBestsResponse](api.Client, "/data/stats/member_bests", parameters)
}

func (api *StatsApi) MemberCareer(parameters *member_career.StatsMemberCareerParams) (*member_career.StatsMemberCareerResponse, error) {
	return client.GetJson[member_career.StatsMemberCareerResponse](api.Client, "/data/stats/member_career", parameters)
}

func (api *StatsApi) MemberDivision(parameters *member_division.StatsMemberDivisionParams) (*member_division.None, error) {
	return client.GetJson[member_division.None](api.Client, "/data/stats/member_division", parameters)
}

func (api *StatsApi) MemberRecap(parameters *member_recap.StatsMemberRecapParams) (*member_recap.StatsMemberRecapResponse, error) {
	return client.GetJson[member_recap.StatsMemberRecapResponse](api.Client, "/data/stats/member_recap", parameters)
}

func (api *StatsApi) MemberRecentRaces(parameters *member_recent_races.StatsMemberRecentRacesParams) (*member_recent_races.StatsMemberRecentRacesResponse, error) {
	return client.GetJson[member_recent_races.StatsMemberRecentRacesResponse](api.Client, "/data/stats/member_recent_races", parameters)
}

func (api *StatsApi) MemberSummary(parameters *member_summary.StatsMemberSummaryParams) (*member_summary.StatsMemberSummaryResponse, error) {
	return client.GetJson[member_summary.StatsMemberSummaryResponse](api.Client, "/data/stats/member_summary", parameters)
}

func (api *StatsApi) MemberYearly(parameters *member_yearly.StatsMemberYearlyParams) (*member_yearly.StatsMemberYearlyResponse, error) {
	return client.GetJson[member_yearly.StatsMemberYearlyResponse](api.Client, "/data/stats/member_yearly", parameters)
}

func (api *StatsApi) SeasonDriverStandings(parameters *season_driver_standings.StatsSeasonDriverStandingsParams) (*season_driver_standings.None, error) {
	return client.GetJson[season_driver_standings.None](api.Client, "/data/stats/season_driver_standings", parameters)
}

func (api *StatsApi) SeasonSupersessionStandings(parameters *season_supersession_standings.StatsSeasonSupersessionStandingsParams) (*season_supersession_standings.None, error) {
	return client.GetJson[season_supersession_standings.None](api.Client, "/data/stats/season_supersession_standings", parameters)
}

func (api *StatsApi) SeasonTeamStandings(parameters *season_team_standings.StatsSeasonTeamStandingsParams) (*season_team_standings.None, error) {
	return client.GetJson[season_team_standings.None](api.Client, "/data/stats/season_team_standings", parameters)
}

func (api *StatsApi) SeasonTtStandings(parameters *season_tt_standings.StatsSeasonTtStandingsParams) (*season_tt_standings.None, error) {
	return client.GetJson[season_tt_standings.None](api.Client, "/data/stats/season_tt_standings", parameters)
}

func (api *StatsApi) SeasonTtResults(parameters *season_tt_results.StatsSeasonTtResultsParams) (*season_tt_results.None, error) {
	return client.GetJson[season_tt_results.None](api.Client, "/data/stats/season_tt_results", parameters)
}

func (api *StatsApi) SeasonQualifyResults(parameters *season_qualify_results.StatsSeasonQualifyResultsParams) (*season_qualify_results.None, error) {
	return client.GetJson[season_qualify_results.None](api.Client, "/data/stats/season_qualify_results", parameters)
}

func (api *StatsApi) WorldRecords(parameters *world_records.StatsWorldRecordsParams) (*world_records.None, error) {
	return client.GetJson[world_records.None](api.Client, "/data/stats/world_records", parameters)
}
