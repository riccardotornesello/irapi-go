package stats

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetStatsMemberSummary(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &StatsApi{
		Client: apiClient,
	}

	_, err := api.GetStatsMemberSummary(StatsMemberSummaryParams{})
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetStatsMemberSummary passed")
}
