package car

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetCar(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &CarApi{
		Client: apiClient,
	}

	_, err := api.GetCar()
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetCar passed")
}
