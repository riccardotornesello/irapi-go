package team

import (
	"github.com/riccardotornesello/irapi-go/api/team/get"
	"github.com/riccardotornesello/irapi-go/api/team/membership"
	"github.com/riccardotornesello/irapi-go/client"
)

type TeamApi struct {
	Client *client.ApiClient
}

func (api *TeamApi) Get(parameters *get.TeamGetParams) (*get.TeamGetResponse, error) {
	return client.GetJson[get.TeamGetResponse](api.Client, "/data/team/get", parameters)
}

func (api *TeamApi) Membership() (*membership.TeamMembershipResponse, error) {
	return client.GetJson[membership.TeamMembershipResponse](api.Client, "/data/team/membership", nil)
}
