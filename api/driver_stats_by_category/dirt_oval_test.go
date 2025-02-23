package driver_stats_by_category

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetDriverStatsByCategoryDirtOval(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &DriverStatsByCategoryApi{
		Client: apiClient,
	}

	_, err := api.GetDriverStatsByCategoryDirtOval()
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetDriverStatsByCategoryDirtOval passed")
}
