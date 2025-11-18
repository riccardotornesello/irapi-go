package event_types

type ConstantsEventTypesResponse []ConstantsEventTypesResponseElement

type ConstantsEventTypesResponseElement struct {
	Label string `json:"label"`
	Value int64  `json:"value"`
}
