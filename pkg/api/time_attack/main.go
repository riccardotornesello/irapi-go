package time_attack

import (
	"github.com/riccardotornesello/irapi-go/pkg/client"
)

type TimeAttackApi struct {
	client *client.ApiClient
}

func NewTimeAttackApi(client *client.ApiClient) *TimeAttackApi {
	return &TimeAttackApi{
		client: client,
	}
}
