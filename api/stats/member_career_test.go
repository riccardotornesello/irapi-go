package stats

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetStatsMemberCareer(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &StatsApi{
		Client: apiClient,
	}

	_, err := api.GetStatsMemberCareer(StatsMemberCareerParams{})
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetStatsMemberCareer passed")
}
