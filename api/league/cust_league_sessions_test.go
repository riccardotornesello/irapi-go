package league

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetLeagueCustLeagueSessions(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &LeagueApi{
		Client: apiClient,
	}

	_, err := api.GetLeagueCustLeagueSessions(LeagueCustLeagueSessionsParams{})
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetLeagueCustLeagueSessions passed")
}
