package league

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetLeagueDirectory(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &LeagueApi{
		Client: apiClient,
	}

	_, err := api.GetLeagueDirectory(LeagueDirectoryParams{})
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetLeagueDirectory passed")
}
