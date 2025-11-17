package assets

import ()

type TrackAssetsResponse map[string]TrackAssetsResponseValue

type TrackAssetsResponseValue struct {
	Coordinates         string         `json:"coordinates"`
	DetailCopy          string         `json:"detail_copy"`
	DetailTechspecsCopy *string        `json:"detail_techspecs_copy"`
	DetailVideo         interface{}    `json:"detail_video"`
	Folder              string         `json:"folder"`
	GalleryImages       *string        `json:"gallery_images"`
	GalleryPrefix       *string        `json:"gallery_prefix"`
	LargeImage          string         `json:"large_image"`
	Logo                string         `json:"logo"`
	North               *string        `json:"north"`
	NumSVGImages        int64          `json:"num_svg_images"`
	SmallImage          string         `json:"small_image"`
	TrackID             int64          `json:"track_id"`
	TrackMap            string         `json:"track_map"`
	TrackMapLayers      TrackMapLayers `json:"track_map_layers"`
}

type TrackMapLayers struct {
	Background  string `json:"background"`
	Inactive    string `json:"inactive"`
	Active      string `json:"active"`
	Pitroad     string `json:"pitroad"`
	StartFinish string `json:"start-finish"`
	Turns       string `json:"turns"`
}
