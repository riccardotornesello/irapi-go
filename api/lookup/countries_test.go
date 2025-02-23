package lookup

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetLookupCountries(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &LookupApi{
		Client: apiClient,
	}

	_, err := api.GetLookupCountries()
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetLookupCountries passed")
}
