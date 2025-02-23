package stats

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetStatsMemberDivision(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &StatsApi{
		Client: apiClient,
	}

	_, err := api.GetStatsMemberDivision(StatsMemberDivisionParams{})
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetStatsMemberDivision passed")
}
