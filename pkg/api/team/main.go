package team

import (
	"github.com/riccardotornesello/irapi-go/pkg/api/team/get"
	"github.com/riccardotornesello/irapi-go/pkg/api/team/membership"
	"github.com/riccardotornesello/irapi-go/pkg/client"
)

type TeamApi struct {
	client *client.ApiClient
}

func NewTeamApi(client *client.ApiClient) *TeamApi {
	return &TeamApi{
		client: client,
	}
}

func (api *TeamApi) Get(parameters *get.TeamGetParams) (*get.TeamGetResponse, error) {
	resp, err := client.GetJson[get.TeamGetResponse](api.client, "/data/team/get", parameters)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *TeamApi) Membership() (*membership.TeamMembershipResponse, error) {
	resp, err := client.GetJson[membership.TeamMembershipResponse](api.client, "/data/team/membership", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
