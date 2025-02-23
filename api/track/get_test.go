package track

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetTrack(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &TrackApi{
		Client: apiClient,
	}

	_, err := api.GetTrack()
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetTrack passed")
}
