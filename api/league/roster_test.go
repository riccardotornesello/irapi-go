package league

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetLeagueRoster(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &LeagueApi{
		Client: apiClient,
	}

	_, err := api.GetLeagueRoster(LeagueRosterParams{})
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetLeagueRoster passed")
}
