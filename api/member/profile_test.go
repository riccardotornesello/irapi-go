package member

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetMemberProfile(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &MemberApi{
		Client: apiClient,
	}

	_, err := api.GetMemberProfile(MemberProfileParams{})
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetMemberProfile passed")
}
