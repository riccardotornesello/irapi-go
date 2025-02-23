package stats

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetStatsMemberRecentRaces(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &StatsApi{
		Client: apiClient,
	}

	_, err := api.GetStatsMemberRecentRaces(StatsMemberRecentRacesParams{})
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetStatsMemberRecentRaces passed")
}
