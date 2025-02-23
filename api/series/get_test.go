package series

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetSeries(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &SeriesApi{
		Client: apiClient,
	}

	_, err := api.GetSeries()
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetSeries passed")
}
