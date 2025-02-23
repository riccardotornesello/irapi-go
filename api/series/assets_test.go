package series

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetSeriesAssets(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &SeriesApi{
		Client: apiClient,
	}

	_, err := api.GetSeriesAssets()
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetSeriesAssets passed")
}
