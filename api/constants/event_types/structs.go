package event_types

import ()

type ConstantsEventTypesResponse []ConstantsEventTypesResponseElement

type ConstantsEventTypesResponseElement struct {
	Label string `json:"label"`
	Value int64  `json:"value"`
}
