package lookup

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetLookupLicenses(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &LookupApi{
		Client: apiClient,
	}

	_, err := api.GetLookupLicenses()
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetLookupLicenses passed")
}
