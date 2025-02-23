package series

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetSeriesSeasons(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &SeriesApi{
		Client: apiClient,
	}

	_, err := api.GetSeriesSeasons(SeriesSeasonsParams{})
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetSeriesSeasons passed")
}
