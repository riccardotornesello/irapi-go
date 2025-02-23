package driver_stats_by_category

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetDriverStatsByCategorySportsCar(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &DriverStatsByCategoryApi{
		Client: apiClient,
	}

	_, err := api.GetDriverStatsByCategorySportsCar()
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetDriverStatsByCategorySportsCar passed")
}
