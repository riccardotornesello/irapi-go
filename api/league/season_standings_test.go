package league

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetLeagueSeasonStandings(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &LeagueApi{
		Client: apiClient,
	}

	_, err := api.GetLeagueSeasonStandings(LeagueSeasonStandingsParams{})
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetLeagueSeasonStandings passed")
}
