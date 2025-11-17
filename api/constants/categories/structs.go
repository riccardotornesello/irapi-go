package categories

import ()

type ConstantsCategoriesResponse []ConstantsCategoriesResponseElement

type ConstantsCategoriesResponseElement struct {
	Label string `json:"label"`
	Value int64  `json:"value"`
}
