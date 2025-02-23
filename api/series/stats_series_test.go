package series

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetSeriesStatsSeries(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &SeriesApi{
		Client: apiClient,
	}

	_, err := api.GetSeriesStatsSeries()
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetSeriesStatsSeries passed")
}
