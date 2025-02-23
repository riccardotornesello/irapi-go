package results

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetResults(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &ResultsApi{
		Client: apiClient,
	}

	_, err := api.GetResults(ResultsGetParams{})
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetResults passed")
}
