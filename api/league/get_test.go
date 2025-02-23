package league

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetLeague(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &LeagueApi{
		Client: apiClient,
	}

	_, err := api.GetLeague(LeagueGetParams{})
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetLeague passed")
}
