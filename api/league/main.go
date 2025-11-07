package league

import (
	"github.com/riccardotornesello/irapi-go/api/league/cust_league_sessions"
	"github.com/riccardotornesello/irapi-go/api/league/directory"
	"github.com/riccardotornesello/irapi-go/api/league/get"
	"github.com/riccardotornesello/irapi-go/api/league/get_points_systems"
	"github.com/riccardotornesello/irapi-go/api/league/membership"
	"github.com/riccardotornesello/irapi-go/api/league/roster"
	"github.com/riccardotornesello/irapi-go/api/league/season_sessions"
	"github.com/riccardotornesello/irapi-go/api/league/season_standings"
	"github.com/riccardotornesello/irapi-go/api/league/seasons"
	"github.com/riccardotornesello/irapi-go/client"
)

type LeagueApi struct {
	Client *client.ApiClient
}

func (api *LeagueApi) CustLeagueSessions(parameters *cust_league_sessions.LeagueCustLeagueSessionsParams) (*cust_league_sessions.LeagueCustLeagueSessionsResponse, error) {
	return client.GetJson[cust_league_sessions.LeagueCustLeagueSessionsResponse](api.Client, "/data/league/cust_league_sessions", parameters)
}

func (api *LeagueApi) Directory(parameters *directory.LeagueDirectoryParams) (*directory.LeagueDirectoryResponse, error) {
	return client.GetJson[directory.LeagueDirectoryResponse](api.Client, "/data/league/directory", parameters)
}

func (api *LeagueApi) Get(parameters *get.LeagueGetParams) (*get.LeagueGetResponse, error) {
	return client.GetJson[get.LeagueGetResponse](api.Client, "/data/league/get", parameters)
}

func (api *LeagueApi) GetPointsSystems(parameters *get_points_systems.LeagueGetPointsSystemsParams) (*get_points_systems.LeagueGetPointsSystemsResponse, error) {
	return client.GetJson[get_points_systems.LeagueGetPointsSystemsResponse](api.Client, "/data/league/get_points_systems", parameters)
}

func (api *LeagueApi) Membership(parameters *membership.LeagueMembershipParams) (*membership.LeagueMembershipResponse, error) {
	return client.GetJson[membership.LeagueMembershipResponse](api.Client, "/data/league/membership", parameters)
}

func (api *LeagueApi) Roster(parameters *roster.LeagueRosterParams) (*roster.None, error) {
	return client.GetJson[roster.None](api.Client, "/data/league/roster", parameters)
}

func (api *LeagueApi) Seasons(parameters *seasons.LeagueSeasonsParams) (*seasons.LeagueSeasonsResponse, error) {
	return client.GetJson[seasons.LeagueSeasonsResponse](api.Client, "/data/league/seasons", parameters)
}

func (api *LeagueApi) SeasonStandings(parameters *season_standings.LeagueSeasonStandingsParams) (*season_standings.LeagueSeasonStandingsResponse, error) {
	return client.GetJson[season_standings.LeagueSeasonStandingsResponse](api.Client, "/data/league/season_standings", parameters)
}

func (api *LeagueApi) SeasonSessions(parameters *season_sessions.LeagueSeasonSessionsParams) (*season_sessions.LeagueSeasonSessionsResponse, error) {
	return client.GetJson[season_sessions.LeagueSeasonSessionsResponse](api.Client, "/data/league/season_sessions", parameters)
}
