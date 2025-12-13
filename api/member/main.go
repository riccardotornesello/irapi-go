package member

import (
	"github.com/riccardotornesello/irapi-go/api/member/award_instances"
	"github.com/riccardotornesello/irapi-go/api/member/awards"
	"github.com/riccardotornesello/irapi-go/api/member/chart_data"
	"github.com/riccardotornesello/irapi-go/api/member/get"
	"github.com/riccardotornesello/irapi-go/api/member/info"
	"github.com/riccardotornesello/irapi-go/api/member/participation_credits"
	"github.com/riccardotornesello/irapi-go/api/member/profile"
	"github.com/riccardotornesello/irapi-go/client"
)

type MemberApi struct {
	client *client.ApiClient
}

func NewMemberApi(client *client.ApiClient) *MemberApi {
	return &MemberApi{
		client: client,
	}
}

func (api *MemberApi) Awards(parameters *awards.MemberAwardsParams) (*awards.MemberAwardsResponse, error) {
	resp, err := client.GetJson[awards.MemberAwardsResponse](api.client, "/data/member/awards", parameters)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *MemberApi) AwardInstances(parameters *award_instances.MemberAwardInstancesParams) (*award_instances.MemberAwardInstancesResponse, error) {
	resp, err := client.GetJson[award_instances.MemberAwardInstancesResponse](api.client, "/data/member/award_instances", parameters)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *MemberApi) ChartData(parameters *chart_data.MemberChartDataParams) (*chart_data.MemberChartDataResponse, error) {
	resp, err := client.GetJson[chart_data.MemberChartDataResponse](api.client, "/data/member/chart_data", parameters)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *MemberApi) Get(parameters *get.MemberGetParams) (*get.MemberGetResponse, error) {
	resp, err := client.GetJson[get.MemberGetResponse](api.client, "/data/member/get", parameters)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *MemberApi) Info() (*info.MemberInfoResponse, error) {
	resp, err := client.GetJson[info.MemberInfoResponse](api.client, "/data/member/info", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *MemberApi) ParticipationCredits() (*participation_credits.MemberParticipationCreditsResponse, error) {
	resp, err := client.GetJson[participation_credits.MemberParticipationCreditsResponse](api.client, "/data/member/participation_credits", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *MemberApi) Profile(parameters *profile.MemberProfileParams) (*profile.MemberProfileResponse, error) {
	resp, err := client.GetJson[profile.MemberProfileResponse](api.client, "/data/member/profile", parameters)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
