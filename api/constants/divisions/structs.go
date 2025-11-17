package divisions

import ()

type ConstantsDivisionsResponse []ConstantsDivisionsResponseElement

type ConstantsDivisionsResponseElement struct {
	Label string `json:"label"`
	Value int64  `json:"value"`
}
