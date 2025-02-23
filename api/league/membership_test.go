package league

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetLeagueMembership(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &LeagueApi{
		Client: apiClient,
	}

	_, err := api.GetLeagueMembership(LeagueMembershipParams{})
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetLeagueMembership passed")
}
