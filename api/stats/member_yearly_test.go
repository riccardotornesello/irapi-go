package stats

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetStatsMemberYearly(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &StatsApi{
		Client: apiClient,
	}

	_, err := api.GetStatsMemberYearly(StatsMemberYearlyParams{})
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetStatsMemberYearly passed")
}
