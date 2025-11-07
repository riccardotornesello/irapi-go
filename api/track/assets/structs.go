package assets

import ()

type TrackAssetsResponse map[string]TrackAssetsResponseValue

type TrackAssetsResponseValue struct {
	Coordinates         string               `json:"coordinates"`
	DetailCopy          string               `json:"detail_copy"`
	DetailTechspecsCopy *DetailTechspecsCopy `json:"detail_techspecs_copy"`
	DetailVideo         interface{}          `json:"detail_video"`
	Folder              string               `json:"folder"`
	GalleryImages       *string              `json:"gallery_images"`
	GalleryPrefix       *string              `json:"gallery_prefix"`
	LargeImage          string               `json:"large_image"`
	Logo                string               `json:"logo"`
	North               *North               `json:"north"`
	NumSVGImages        int64                `json:"num_svg_images"`
	SmallImage          string               `json:"small_image"`
	TrackID             int64                `json:"track_id"`
	TrackMap            string               `json:"track_map"`
	TrackMapLayers      TrackMapLayers       `json:"track_map_layers"`
}

type TrackMapLayers struct {
	Background  Background  `json:"background"`
	Inactive    Inactive    `json:"inactive"`
	Active      Active      `json:"active"`
	Pitroad     Pitroad     `json:"pitroad"`
	StartFinish StartFinish `json:"start-finish"`
	Turns       Turns       `json:"turns"`
}

type DetailTechspecsCopy string

const (
	Dtsc2Ta DetailTechspecsCopy = "dtsc2 ta"
	Null    DetailTechspecsCopy = "null"
)

type North string

const (
	The160Deg North = "160deg"
	The90Deg  North = "90deg"
)

type Active string

const (
	ActiveSVG Active = "active.svg"
)

type Background string

const (
	BackgroundSVG Background = "background.svg"
)

type Inactive string

const (
	InactiveSVG Inactive = "inactive.svg"
)

type Pitroad string

const (
	PitroadSVG Pitroad = "pitroad.svg"
)

type StartFinish string

const (
	StartFinishSVG StartFinish = "start-finish.svg"
)

type Turns string

const (
	TurnsSVG Turns = "turns.svg"
)
