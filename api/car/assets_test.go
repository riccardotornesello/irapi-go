package car

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetCarAssets(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &CarApi{
		Client: apiClient,
	}

	_, err := api.GetCarAssets()
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetCarAssets passed")
}
