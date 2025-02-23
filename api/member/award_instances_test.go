package member

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetMemberAwardInstances(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &MemberApi{
		Client: apiClient,
	}

	_, err := api.GetMemberAwardInstances(MemberAwardInstancesParams{})
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetMemberAwardInstances passed")
}
