package constants

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetConstantsDivisions(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &ConstantsApi{
		Client: apiClient,
	}

	_, err := api.GetConstantsDivisions()
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetConstantsDivisions passed")
}
