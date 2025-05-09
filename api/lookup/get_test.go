package lookup

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetLookup(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &LookupApi{
		Client: apiClient,
	}

	_, err := api.GetLookup()
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetLookup passed")
}
