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
	resp, err := client.GetJson[get.TeamGetResponse](api.Client, "/data/team/get", parameters)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *TeamApi) Membership() (*membership.TeamMembershipResponse, error) {
	resp, err := client.GetJson[membership.TeamMembershipResponse](api.Client, "/data/team/membership", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
