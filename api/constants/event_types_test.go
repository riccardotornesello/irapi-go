package constants

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetConstantsEventTypes(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &ConstantsApi{
		Client: apiClient,
	}

	_, err := api.GetConstantsEventTypes()
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetConstantsEventTypes passed")
}
