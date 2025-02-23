package track

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetTrackAssets(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &TrackApi{
		Client: apiClient,
	}

	_, err := api.GetTrackAssets()
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetTrackAssets passed")
}
