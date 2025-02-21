package irapi

import (
	"riccardotornesello.it/sharedtelemetry/iracing/irapi/api/car"
	"riccardotornesello.it/sharedtelemetry/iracing/irapi/api/carclass"
	"riccardotornesello.it/sharedtelemetry/iracing/irapi/api/constants"
	"riccardotornesello.it/sharedtelemetry/iracing/irapi/api/driver_stats_by_category"
	"riccardotornesello.it/sharedtelemetry/iracing/irapi/api/hosted"
	"riccardotornesello.it/sharedtelemetry/iracing/irapi/api/league"
	"riccardotornesello.it/sharedtelemetry/iracing/irapi/api/lookup"
	"riccardotornesello.it/sharedtelemetry/iracing/irapi/api/member"
	"riccardotornesello.it/sharedtelemetry/iracing/irapi/api/results"
	"riccardotornesello.it/sharedtelemetry/iracing/irapi/api/season"
	"riccardotornesello.it/sharedtelemetry/iracing/irapi/api/series"
	"riccardotornesello.it/sharedtelemetry/iracing/irapi/api/stats"
	"riccardotornesello.it/sharedtelemetry/iracing/irapi/api/team"
	"riccardotornesello.it/sharedtelemetry/iracing/irapi/api/time_attack"
	"riccardotornesello.it/sharedtelemetry/iracing/irapi/api/track"
	"riccardotornesello.it/sharedtelemetry/iracing/irapi/client"
)

type IRacingApiClient struct {
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
	Stats                 *stats.StatsApi
	Team                  *team.TeamApi
	TimeAttack            *time_attack.TimeAttackApi
	Track                 *track.TrackApi
}

func NewIRacingApiClient(email string, password string) (*IRacingApiClient, error) {
	apiClient, err := client.NewApiClient(email, password)
	if err != nil {
		return nil, err
	}

	return &IRacingApiClient{
		Car:                   &car.CarApi{Client: apiClient},
		Carclass:              &carclass.CarclassApi{Client: apiClient},
		Constants:             &constants.ConstantsApi{Client: apiClient},
		DriverStatsByCategory: &driver_stats_by_category.DriverStatsByCategoryApi{Client: apiClient},
		Hosted:                &hosted.HostedApi{Client: apiClient},
		League:                &league.LeagueApi{Client: apiClient},
		Lookup:                &lookup.LookupApi{Client: apiClient},
		Member:                &member.MemberApi{Client: apiClient},
		Results:               &results.ResultsApi{Client: apiClient},
		Season:                &season.SeasonApi{Client: apiClient},
		Series:                &series.SeriesApi{Client: apiClient},
		Stats:                 &stats.StatsApi{Client: apiClient},
		Team:                  &team.TeamApi{Client: apiClient},
		TimeAttack:            &time_attack.TimeAttackApi{Client: apiClient},
		Track:                 &track.TrackApi{Client: apiClient},
	}, nil
}
