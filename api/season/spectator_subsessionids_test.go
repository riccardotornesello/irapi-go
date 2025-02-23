package season

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetSeasonSpectatorSubsessionids(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &SeasonApi{
		Client: apiClient,
	}

	_, err := api.GetSeasonSpectatorSubsessionids(SeasonSpectatorSubsessionidsParams{})
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetSeasonSpectatorSubsessionids passed")
}
