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
	client *client.ApiClient
}

func NewLeagueApi(client *client.ApiClient) *LeagueApi {
	return &LeagueApi{
		client: client,
	}
}

func (api *LeagueApi) CustLeagueSessions(parameters *cust_league_sessions.LeagueCustLeagueSessionsParams) (*cust_league_sessions.LeagueCustLeagueSessionsResponse, error) {
	resp, err := client.GetJson[cust_league_sessions.LeagueCustLeagueSessionsResponse](api.client, "/data/league/cust_league_sessions", parameters)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *LeagueApi) Directory(parameters *directory.LeagueDirectoryParams) (*directory.LeagueDirectoryResponse, error) {
	resp, err := client.GetJson[directory.LeagueDirectoryResponse](api.client, "/data/league/directory", parameters)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *LeagueApi) Get(parameters *get.LeagueGetParams) (*get.LeagueGetResponse, error) {
	resp, err := client.GetJson[get.LeagueGetResponse](api.client, "/data/league/get", parameters)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *LeagueApi) GetPointsSystems(parameters *get_points_systems.LeagueGetPointsSystemsParams) (*get_points_systems.LeagueGetPointsSystemsResponse, error) {
	resp, err := client.GetJson[get_points_systems.LeagueGetPointsSystemsResponse](api.client, "/data/league/get_points_systems", parameters)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *LeagueApi) Membership(parameters *membership.LeagueMembershipParams) (*membership.LeagueMembershipResponse, error) {
	resp, err := client.GetJson[membership.LeagueMembershipResponse](api.client, "/data/league/membership", parameters)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *LeagueApi) Roster(parameters *roster.LeagueRosterParams) (*roster.LeagueRosterResponse, error) {
	resp, err := client.GetJson[roster.LeagueRosterResponse](api.client, "/data/league/roster", parameters)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *LeagueApi) Seasons(parameters *seasons.LeagueSeasonsParams) (*seasons.LeagueSeasonsResponse, error) {
	resp, err := client.GetJson[seasons.LeagueSeasonsResponse](api.client, "/data/league/seasons", parameters)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *LeagueApi) SeasonStandings(parameters *season_standings.LeagueSeasonStandingsParams) (*season_standings.LeagueSeasonStandingsResponse, error) {
	resp, err := client.GetJson[season_standings.LeagueSeasonStandingsResponse](api.client, "/data/league/season_standings", parameters)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *LeagueApi) SeasonSessions(parameters *season_sessions.LeagueSeasonSessionsParams) (*season_sessions.LeagueSeasonSessionsResponse, error) {
	resp, err := client.GetJson[season_sessions.LeagueSeasonSessionsResponse](api.client, "/data/league/season_sessions", parameters)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
