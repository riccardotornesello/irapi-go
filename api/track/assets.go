package track

import (
	"encoding/json"
)

type TrackAssetsResponse map[string]struct {
	Coordinates         string      `json:"coordinates"`
	DetailCopy          string      `json:"detail_copy"`
	DetailTechspecsCopy *string     `json:"detail_techspecs_copy"`
	DetailVideo         interface{} `json:"detail_video"`
	Folder              string      `json:"folder"`
	GalleryImages       *string     `json:"gallery_images"`
	GalleryPrefix       *string     `json:"gallery_prefix"`
	LargeImage          string      `json:"large_image"`
	Logo                string      `json:"logo"`
	North               *string     `json:"north"`
	NumSvgImages        int         `json:"num_svg_images"`
	SmallImage          string      `json:"small_image"`
	TrackId             int         `json:"track_id"`
	TrackMap            string      `json:"track_map"`
	TrackMapLayers      struct {
		Background  string `json:"background"`
		Inactive    string `json:"inactive"`
		Active      string `json:"active"`
		Pitroad     string `json:"pitroad"`
		StartFinish string `json:"start-finish"`
		Turns       string `json:"turns"`
	} `json:"track_map_layers"`
}

// image paths are relative to https://images-static.iracing.com/
func (api *TrackApi) GetTrackAssets() (*TrackAssetsResponse, error) {
	url := "https://members-ng.iracing.com/data/track/assets"

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &TrackAssetsResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
