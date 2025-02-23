package constants

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetConstantsCategories(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &ConstantsApi{
		Client: apiClient,
	}

	_, err := api.GetConstantsCategories()
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetConstantsCategories passed")
}
