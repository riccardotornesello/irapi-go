package league

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetLeagueSeasonSessions(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &LeagueApi{
		Client: apiClient,
	}

	_, err := api.GetLeagueSeasonSessions(LeagueSeasonSessionsParams{})
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetLeagueSeasonSessions passed")
}
