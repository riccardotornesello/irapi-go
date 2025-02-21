package car

import (
	"encoding/json"
)

type CarAssetsResponse map[string]struct {
	CarId    int `json:"car_id"`
	CarRules []struct {
		RuleCategory string `json:"rule_category"`
		Text         string `json:"text"`
	} `json:"car_rules"`
	DetailCopy             string      `json:"detail_copy"`
	DetailScreenShotImages string      `json:"detail_screen_shot_images"`
	DetailTechspecsCopy    string      `json:"detail_techspecs_copy"`
	Folder                 string      `json:"folder"`
	GalleryImages          *string     `json:"gallery_images"`
	GalleryPrefix          interface{} `json:"gallery_prefix"`
	GroupImage             *string     `json:"group_image"`
	GroupName              interface{} `json:"group_name"`
	LargeImage             string      `json:"large_image"`
	Logo                   *string     `json:"logo"`
	SmallImage             string      `json:"small_image"`
	SponsorLogo            *string     `json:"sponsor_logo"`
	TemplatePath           *string     `json:"template_path"`
}

func (api *CarApi) GetCarAssets() (*CarAssetsResponse, error) {
	url := "https://members-ng.iracing.com/data/car/assets"

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &CarAssetsResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
