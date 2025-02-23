package member

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetMemberAwards(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &MemberApi{
		Client: apiClient,
	}

	_, err := api.GetMemberAwards(MemberAwardsParams{})
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetMemberAwards passed")
}
