package member

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetMemberInfo(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &MemberApi{
		Client: apiClient,
	}

	_, err := api.GetMemberInfo()
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetMemberInfo passed")
}
