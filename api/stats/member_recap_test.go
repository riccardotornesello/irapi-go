package stats

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetStatsMemberRecap(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &StatsApi{
		Client: apiClient,
	}

	_, err := api.GetStatsMemberRecap(StatsMemberRecapParams{})
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetStatsMemberRecap passed")
}
