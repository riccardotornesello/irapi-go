package team

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetTeam(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &TeamApi{
		Client: apiClient,
	}

	_, err := api.GetTeam(TeamGetParams{})
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetTeam passed")
}
