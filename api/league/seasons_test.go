package league

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetLeagueSeasons(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &LeagueApi{
		Client: apiClient,
	}

	_, err := api.GetLeagueSeasons(LeagueSeasonsParams{})
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetLeagueSeasons passed")
}
