package assets

import ()

type CarAssetsResponse map[string]CarAssetsResponseValue

type CarAssetsResponseValue struct {
	CarID                  int64       `json:"car_id"`
	CarRules               []CarRule   `json:"car_rules"`
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

type CarRule struct {
	RuleCategory string `json:"rule_category"`
	Text         string `json:"text"`
}
