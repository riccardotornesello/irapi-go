package irapi

import (
	"github.com/riccardotornesello/irapi-go/api/car"
	"github.com/riccardotornesello/irapi-go/api/carclass"
	"github.com/riccardotornesello/irapi-go/api/constants"
	"github.com/riccardotornesello/irapi-go/api/driver_stats_by_category"
	"github.com/riccardotornesello/irapi-go/api/hosted"
	"github.com/riccardotornesello/irapi-go/api/league"
	"github.com/riccardotornesello/irapi-go/api/lookup"
	"github.com/riccardotornesello/irapi-go/api/member"
	"github.com/riccardotornesello/irapi-go/api/results"
	"github.com/riccardotornesello/irapi-go/api/season"
	"github.com/riccardotornesello/irapi-go/api/series"
	"github.com/riccardotornesello/irapi-go/api/stats"
	"github.com/riccardotornesello/irapi-go/api/team"
	"github.com/riccardotornesello/irapi-go/api/time_attack"
	"github.com/riccardotornesello/irapi-go/api/track"
	"github.com/riccardotornesello/irapi-go/client"
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
