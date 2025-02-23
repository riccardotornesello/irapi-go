package league

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetLeagueGetPointsSystems(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &LeagueApi{
		Client: apiClient,
	}

	_, err := api.GetLeagueGetPointsSystems(LeagueGetPointsSystemsParams{})
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetLeagueGetPointsSystems passed")
}
