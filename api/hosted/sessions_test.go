package hosted

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetHostedSessions(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &HostedApi{
		Client: apiClient,
	}

	_, err := api.GetHostedSessions()
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetHostedSessions passed")
}
