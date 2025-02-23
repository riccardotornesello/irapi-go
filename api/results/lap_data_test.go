package results

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetResultsLapData(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &ResultsApi{
		Client: apiClient,
	}

	_, err := api.GetResultsLapData(ResultsLapDataParams{})
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetResultsLapData passed")
}
