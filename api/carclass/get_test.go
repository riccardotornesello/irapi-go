package carclass

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetCarclass(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &CarclassApi{
		Client: apiClient,
	}

	_, err := api.GetCarclass()
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetCarclass passed")
}
