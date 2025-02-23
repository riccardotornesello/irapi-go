package season

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetSeasonSpectatorSubsessionidsDetail(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &SeasonApi{
		Client: apiClient,
	}

	_, err := api.GetSeasonSpectatorSubsessionidsDetail(SeasonSpectatorSubsessionidsDetailParams{})
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetSeasonSpectatorSubsessionidsDetail passed")
}
