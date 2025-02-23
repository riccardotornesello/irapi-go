package season

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetSeasonRaceGuide(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &SeasonApi{
		Client: apiClient,
	}

	_, err := api.GetSeasonRaceGuide(SeasonRaceGuideParams{})
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetSeasonRaceGuide passed")
}
