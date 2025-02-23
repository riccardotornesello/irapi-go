package hosted

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetHostedCombinedSessions(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &HostedApi{
		Client: apiClient,
	}

	_, err := api.GetHostedCombinedSessions(HostedCombinedSessionsParams{})
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetHostedCombinedSessions passed")
}
