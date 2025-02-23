package member

import (
	"log"
	"testing"

	"github.com/riccardotornesello/irapi-go/testutils"
)

func TestGetMemberChartData(t *testing.T) {
	apiClient := testutils.GetApiClient()

	api := &MemberApi{
		Client: apiClient,
	}

	_, err := api.GetMemberChartData(MemberChartDataParams{})
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetMemberChartData passed")
}
