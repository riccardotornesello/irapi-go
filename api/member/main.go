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
	Client *client.ApiClient
}

func (api *MemberApi) Awards(parameters *awards.MemberAwardsParams) (*awards.MemberAwardsResponse, error) {
	return client.GetJson[awards.MemberAwardsResponse](api.Client, "/data/member/awards", parameters)
}

func (api *MemberApi) AwardInstances(parameters *award_instances.MemberAwardInstancesParams) (*award_instances.MemberAwardInstancesResponse, error) {
	return client.GetJson[award_instances.MemberAwardInstancesResponse](api.Client, "/data/member/award_instances", parameters)
}

func (api *MemberApi) ChartData(parameters *chart_data.MemberChartDataParams) (*chart_data.MemberChartDataResponse, error) {
	return client.GetJson[chart_data.MemberChartDataResponse](api.Client, "/data/member/chart_data", parameters)
}

func (api *MemberApi) Get(parameters *get.MemberGetParams) (*get.MemberGetResponse, error) {
	return client.GetJson[get.MemberGetResponse](api.Client, "/data/member/get", parameters)
}

func (api *MemberApi) Info() (*info.MemberInfoResponse, error) {
	return client.GetJson[info.MemberInfoResponse](api.Client, "/data/member/info", nil)
}

func (api *MemberApi) ParticipationCredits() (*participation_credits.MemberParticipationCreditsResponse, error) {
	return client.GetJson[participation_credits.MemberParticipationCreditsResponse](api.Client, "/data/member/participation_credits", nil)
}

func (api *MemberApi) Profile(parameters *profile.MemberProfileParams) (*profile.MemberProfileResponse, error) {
	return client.GetJson[profile.MemberProfileResponse](api.Client, "/data/member/profile", parameters)
}
