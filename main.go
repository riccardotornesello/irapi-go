package irapi

import (
	"github.com/riccardotornesello/irapi-go/pkg/api/car"
	"github.com/riccardotornesello/irapi-go/pkg/api/carclass"
	"github.com/riccardotornesello/irapi-go/pkg/api/constants"
	"github.com/riccardotornesello/irapi-go/pkg/api/driver_stats_by_category"
	"github.com/riccardotornesello/irapi-go/pkg/api/hosted"
	"github.com/riccardotornesello/irapi-go/pkg/api/league"
	"github.com/riccardotornesello/irapi-go/pkg/api/lookup"
	"github.com/riccardotornesello/irapi-go/pkg/api/member"
	"github.com/riccardotornesello/irapi-go/pkg/api/results"
	"github.com/riccardotornesello/irapi-go/pkg/api/season"
	"github.com/riccardotornesello/irapi-go/pkg/api/series"
	"github.com/riccardotornesello/irapi-go/pkg/api/session"
	"github.com/riccardotornesello/irapi-go/pkg/api/stats"
	"github.com/riccardotornesello/irapi-go/pkg/api/team"
	"github.com/riccardotornesello/irapi-go/pkg/api/time_attack"
	"github.com/riccardotornesello/irapi-go/pkg/api/track"
	"github.com/riccardotornesello/irapi-go/pkg/client"
)

type IRacingApiClient struct {
	Client *client.ApiClient

	Car                   *car.CarApi
	Carclass              *carclass.CarclassApi
	Constants             *constants.ConstantsApi
	DriverStatsByCategory *driver_stats_by_category.DriverStatsByCategoryApi
	Hosted                *hosted.HostedApi
	League                *league.LeagueApi
	Lookup                *lookup.LookupApi
	Member                *member.MemberApi
	Results               *results.ResultsApi
	Season                *season.SeasonApi
	Series                *series.SeriesApi
	Session               *session.SessionApi
	Stats                 *stats.StatsApi
	Team                  *team.TeamApi
	TimeAttack            *time_attack.TimeAttackApi
	Track                 *track.TrackApi
}

func genClient(apiClient *client.ApiClient) *IRacingApiClient {
	return &IRacingApiClient{
		Client: apiClient,

		Car:                   car.NewCarApi(apiClient),
		Carclass:              carclass.NewCarclassApi(apiClient),
		Constants:             constants.NewConstantsApi(apiClient),
		DriverStatsByCategory: driver_stats_by_category.NewDriverStatsByCategoryApi(apiClient),
		Hosted:                hosted.NewHostedApi(apiClient),
		League:                league.NewLeagueApi(apiClient),
		Lookup:                lookup.NewLookupApi(apiClient),
		Member:                member.NewMemberApi(apiClient),
		Results:               results.NewResultsApi(apiClient),
		Season:                season.NewSeasonApi(apiClient),
		Series:                series.NewSeriesApi(apiClient),
		Session:               session.NewSessionApi(apiClient),
		Stats:                 stats.NewStatsApi(apiClient),
		Team:                  team.NewTeamApi(apiClient),
		TimeAttack:            time_attack.NewTimeAttackApi(apiClient),
		Track:                 track.NewTrackApi(apiClient),
	}
}

func NewIRacingPasswordLimitedApiClient(clientId, clientSecret, username, password string) (*IRacingApiClient, error) {
	apiClient, err := client.NewPasswordLimitedApiClient(username, password, &client.Options{
		ClientId:     clientId,
		ClientSecret: clientSecret,
	})
	if err != nil {
		return nil, err
	}

	return genClient(apiClient), nil
}
