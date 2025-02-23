package lookup

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetLookupClubHistory(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &LookupApi{
		Client: apiClient,
	}

	_, err := api.GetLookupClubHistory(LookupClubHistoryParams{})
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetLookupClubHistory passed")
}
