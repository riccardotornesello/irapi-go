package member

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetMemberParticipationCredits(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &MemberApi{
		Client: apiClient,
	}

	_, err := api.GetMemberParticipationCredits()
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetMemberParticipationCredits passed")
}
