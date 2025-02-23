package member

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetMember(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &MemberApi{
		Client: apiClient,
	}

	_, err := api.GetMember(MemberGetParams{})
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetMember passed")
}
