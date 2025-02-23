package stats

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetStatsMemberBests(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &StatsApi{
		Client: apiClient,
	}

	_, err := api.GetStatsMemberBests(StatsMemberBestsParams{})
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetStatsMemberBests passed")
}
