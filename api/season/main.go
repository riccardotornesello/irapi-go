package season

import (
	"github.com/riccardotornesello/irapi-go/api/season/race_guide"
	"github.com/riccardotornesello/irapi-go/api/season/spectator_subsessionids"
	"github.com/riccardotornesello/irapi-go/api/season/spectator_subsessionids_detail"
	"github.com/riccardotornesello/irapi-go/client"
)

type SeasonApi struct {
	Client *client.ApiClient
}

func (api *SeasonApi) RaceGuide(parameters *race_guide.SeasonRaceGuideParams) (*race_guide.SeasonRaceGuideResponse, error) {
	resp, err := client.GetJson[race_guide.SeasonRaceGuideResponse](api.Client, "/data/season/race_guide", parameters)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *SeasonApi) SpectatorSubsessionids(parameters *spectator_subsessionids.SeasonSpectatorSubsessionidsParams) (*spectator_subsessionids.SeasonSpectatorSubsessionidsResponse, error) {
	resp, err := client.GetJson[spectator_subsessionids.SeasonSpectatorSubsessionidsResponse](api.Client, "/data/season/spectator_subsessionids", parameters)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *SeasonApi) SpectatorSubsessionidsDetail(parameters *spectator_subsessionids_detail.SeasonSpectatorSubsessionidsDetailParams) (*spectator_subsessionids_detail.SeasonSpectatorSubsessionidsDetailResponse, error) {
	resp, err := client.GetJson[spectator_subsessionids_detail.SeasonSpectatorSubsessionidsDetailResponse](api.Client, "/data/season/spectator_subsessionids_detail", parameters)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
