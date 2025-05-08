package car

import (
	"encoding/json"
)

type CarGetResponse []struct {
	AiEnabled          bool   `json:"ai_enabled"`
	AllowNumberColors  bool   `json:"allow_number_colors"`
	AllowNumberFont    bool   `json:"allow_number_font"`
	AllowSponsor1      bool   `json:"allow_sponsor1"`
	AllowSponsor2      bool   `json:"allow_sponsor2"`
	AllowWheelColor    bool   `json:"allow_wheel_color"`
	AwardExempt        bool   `json:"award_exempt"`
	CarDirpath         string `json:"car_dirpath"`
	CarId              int    `json:"car_id"`
	CarName            string `json:"car_name"`
	CarNameAbbreviated string `json:"car_name_abbreviated"`
	CarTypes           []struct {
		CarType string `json:"car_type"`
	} `json:"car_types"`
	CarWeight               int      `json:"car_weight"`
	Categories              []string `json:"categories"`
	Created                 string   `json:"created"`
	FirstSale               string   `json:"first_sale"`
	ForumUrl                string   `json:"forum_url"`
	FreeWithSubscription    bool     `json:"free_with_subscription"`
	HasHeadlights           bool     `json:"has_headlights"`
	HasMultipleDryTireTypes bool     `json:"has_multiple_dry_tire_types"`
	HasRainCapableTireTypes bool     `json:"has_rain_capable_tire_types"`
	Hp                      int      `json:"hp"`
	IsPsPurchasable         bool     `json:"is_ps_purchasable"`
	MaxPowerAdjustPct       float64  `json:"max_power_adjust_pct"`
	MaxWeightPenaltyKg      int      `json:"max_weight_penalty_kg"`
	MinPowerAdjustPct       float64  `json:"min_power_adjust_pct"`
	PackageId               int      `json:"package_id"`
	Patterns                int      `json:"patterns"`
	Price                   float64  `json:"price"`
	PriceDisplay            string   `json:"price_display"`
	RainEnabled             bool     `json:"rain_enabled"`
	Retired                 bool     `json:"retired"`
	SearchFilters           string   `json:"search_filters"`
	Sku                     int      `json:"sku"`
	CarMake                 string   `json:"car_make"`
	CarModel                string   `json:"car_model"`
	PaintRules              struct {
		Field1 struct {
			Paintcaravailable       bool   `json:"PaintCarAvailable"`
			Color1                  string `json:"Color1"`
			Color2                  string `json:"Color2"`
			Color3                  string `json:"Color3"`
			Sponsor1available       bool   `json:"Sponsor1Available"`
			Sponsor2available       bool   `json:"Sponsor2Available"`
			Sponsor1                string `json:"Sponsor1"`
			Sponsor2                string `json:"Sponsor2"`
			Rulesexplanation        string `json:"RulesExplanation"`
			Paintwheelavailable     bool   `json:"PaintWheelAvailable"`
			Wheelcolor              string `json:"WheelColor"`
			Rimtypeavailable        bool   `json:"RimTypeAvailable"`
			Rimtype                 string `json:"RimType"`
			Allownumberfontchanges  bool   `json:"AllowNumberFontChanges"`
			Numberfont              string `json:"NumberFont"`
			Allownumbercolorchanges bool   `json:"AllowNumberColorChanges"`
			Numbercolor1            string `json:"NumberColor1"`
			Numbercolor2            string `json:"NumberColor2"`
			Numbercolor3            string `json:"NumberColor3"`
		} `json:"1"`
		Restrictcustompaint bool `json:"RestrictCustomPaint"`
		Field24             struct {
			Paintcaravailable bool   `json:"PaintCarAvailable"`
			Color1            string `json:"Color1"`
			Color2            string `json:"Color2"`
			Color3            string `json:"Color3"`
			Sponsor1available bool   `json:"Sponsor1Available"`
			Sponsor2available bool   `json:"Sponsor2Available"`
			Sponsor1          string `json:"Sponsor1"`
			Sponsor2          string `json:"Sponsor2"`
			Rulesexplanation  string `json:"RulesExplanation"`
		} `json:"24"`
		Field2 struct {
			Paintcaravailable       bool   `json:"PaintCarAvailable"`
			Color1                  string `json:"Color1"`
			Color2                  string `json:"Color2"`
			Color3                  string `json:"Color3"`
			Sponsor1available       bool   `json:"Sponsor1Available"`
			Sponsor2available       bool   `json:"Sponsor2Available"`
			Sponsor1                string `json:"Sponsor1"`
			Sponsor2                string `json:"Sponsor2"`
			Rulesexplanation        string `json:"RulesExplanation"`
			Paintwheelavailable     bool   `json:"PaintWheelAvailable"`
			Wheelcolor              string `json:"WheelColor"`
			Rimtypeavailable        bool   `json:"RimTypeAvailable"`
			Rimtype                 string `json:"RimType"`
			Allownumberfontchanges  bool   `json:"AllowNumberFontChanges"`
			Numberfont              string `json:"NumberFont"`
			Allownumbercolorchanges bool   `json:"AllowNumberColorChanges"`
			Numbercolor1            string `json:"NumberColor1"`
			Numbercolor2            string `json:"NumberColor2"`
			Numbercolor3            string `json:"NumberColor3"`
		} `json:"2"`
		Field3 struct {
			Paintcaravailable       bool   `json:"PaintCarAvailable"`
			Color1                  string `json:"Color1"`
			Color2                  string `json:"Color2"`
			Color3                  string `json:"Color3"`
			Sponsor1available       bool   `json:"Sponsor1Available"`
			Sponsor2available       bool   `json:"Sponsor2Available"`
			Sponsor1                string `json:"Sponsor1"`
			Sponsor2                string `json:"Sponsor2"`
			Rulesexplanation        string `json:"RulesExplanation"`
			Paintwheelavailable     bool   `json:"PaintWheelAvailable"`
			Wheelcolor              string `json:"WheelColor"`
			Rimtypeavailable        bool   `json:"RimTypeAvailable"`
			Rimtype                 string `json:"RimType"`
			Allownumberfontchanges  bool   `json:"AllowNumberFontChanges"`
			Numberfont              string `json:"NumberFont"`
			Allownumbercolorchanges bool   `json:"AllowNumberColorChanges"`
			Numbercolor1            string `json:"NumberColor1"`
			Numbercolor2            string `json:"NumberColor2"`
			Numbercolor3            string `json:"NumberColor3"`
		} `json:"3"`
		Field4 struct {
			Paintcaravailable       bool   `json:"PaintCarAvailable"`
			Color1                  string `json:"Color1"`
			Color2                  string `json:"Color2"`
			Color3                  string `json:"Color3"`
			Sponsor1available       bool   `json:"Sponsor1Available"`
			Sponsor2available       bool   `json:"Sponsor2Available"`
			Sponsor1                string `json:"Sponsor1"`
			Sponsor2                string `json:"Sponsor2"`
			Rulesexplanation        string `json:"RulesExplanation"`
			Paintwheelavailable     bool   `json:"PaintWheelAvailable"`
			Wheelcolor              string `json:"WheelColor"`
			Rimtypeavailable        bool   `json:"RimTypeAvailable"`
			Rimtype                 string `json:"RimType"`
			Allownumberfontchanges  bool   `json:"AllowNumberFontChanges"`
			Numberfont              string `json:"NumberFont"`
			Allownumbercolorchanges bool   `json:"AllowNumberColorChanges"`
			Numbercolor1            string `json:"NumberColor1"`
			Numbercolor2            string `json:"NumberColor2"`
			Numbercolor3            string `json:"NumberColor3"`
		} `json:"4"`
		Field5 struct {
			Paintcaravailable       bool   `json:"PaintCarAvailable"`
			Color1                  string `json:"Color1"`
			Color2                  string `json:"Color2"`
			Color3                  string `json:"Color3"`
			Sponsor1available       bool   `json:"Sponsor1Available"`
			Sponsor2available       bool   `json:"Sponsor2Available"`
			Sponsor1                string `json:"Sponsor1"`
			Sponsor2                string `json:"Sponsor2"`
			Rulesexplanation        string `json:"RulesExplanation"`
			Paintwheelavailable     bool   `json:"PaintWheelAvailable"`
			Wheelcolor              string `json:"WheelColor"`
			Rimtypeavailable        bool   `json:"RimTypeAvailable"`
			Rimtype                 string `json:"RimType"`
			Allownumberfontchanges  bool   `json:"AllowNumberFontChanges"`
			Numberfont              string `json:"NumberFont"`
			Allownumbercolorchanges bool   `json:"AllowNumberColorChanges"`
			Numbercolor1            string `json:"NumberColor1"`
			Numbercolor2            string `json:"NumberColor2"`
			Numbercolor3            string `json:"NumberColor3"`
		} `json:"5"`
		Field0 struct {
			Paintcaravailable       bool   `json:"PaintCarAvailable"`
			Color1                  string `json:"Color1"`
			Color2                  string `json:"Color2"`
			Color3                  string `json:"Color3"`
			Sponsor1available       bool   `json:"Sponsor1Available"`
			Sponsor2available       bool   `json:"Sponsor2Available"`
			Sponsor1                string `json:"Sponsor1"`
			Sponsor2                string `json:"Sponsor2"`
			Paintwheelavailable     bool   `json:"PaintWheelAvailable"`
			Wheelcolor              string `json:"WheelColor"`
			Rimtypeavailable        bool   `json:"RimTypeAvailable"`
			Rimtype                 string `json:"RimType"`
			Allownumberfontchanges  bool   `json:"AllowNumberFontChanges"`
			Numberfont              string `json:"NumberFont"`
			Allownumbercolorchanges bool   `json:"AllowNumberColorChanges"`
			Numbercolor1            string `json:"NumberColor1"`
			Numbercolor2            string `json:"NumberColor2"`
			Numbercolor3            string `json:"NumberColor3"`
			Rulesexplanation        string `json:"RulesExplanation"`
		} `json:"0"`
		Field6 struct {
			Paintcaravailable       bool   `json:"PaintCarAvailable"`
			Color1                  string `json:"Color1"`
			Color2                  string `json:"Color2"`
			Color3                  string `json:"Color3"`
			Sponsor1available       bool   `json:"Sponsor1Available"`
			Sponsor2available       bool   `json:"Sponsor2Available"`
			Sponsor1                string `json:"Sponsor1"`
			Sponsor2                string `json:"Sponsor2"`
			Paintwheelavailable     bool   `json:"PaintWheelAvailable"`
			Wheelcolor              string `json:"WheelColor"`
			Rimtypeavailable        bool   `json:"RimTypeAvailable"`
			Rimtype                 string `json:"RimType"`
			Allownumberfontchanges  bool   `json:"AllowNumberFontChanges"`
			Numberfont              string `json:"NumberFont"`
			Allownumbercolorchanges bool   `json:"AllowNumberColorChanges"`
			Numbercolor1            string `json:"NumberColor1"`
			Numbercolor2            string `json:"NumberColor2"`
			Numbercolor3            string `json:"NumberColor3"`
			Rulesexplanation        string `json:"RulesExplanation"`
		} `json:"6"`
		Field7 struct {
			Paintcaravailable       bool   `json:"PaintCarAvailable"`
			Color1                  string `json:"Color1"`
			Color2                  string `json:"Color2"`
			Color3                  string `json:"Color3"`
			Sponsor1available       bool   `json:"Sponsor1Available"`
			Sponsor2available       bool   `json:"Sponsor2Available"`
			Sponsor1                string `json:"Sponsor1"`
			Sponsor2                string `json:"Sponsor2"`
			Paintwheelavailable     bool   `json:"PaintWheelAvailable"`
			Wheelcolor              string `json:"WheelColor"`
			Rimtypeavailable        bool   `json:"RimTypeAvailable"`
			Rimtype                 string `json:"RimType"`
			Allownumberfontchanges  bool   `json:"AllowNumberFontChanges"`
			Numberfont              string `json:"NumberFont"`
			Allownumbercolorchanges bool   `json:"AllowNumberColorChanges"`
			Numbercolor1            string `json:"NumberColor1"`
			Numbercolor2            string `json:"NumberColor2"`
			Numbercolor3            string `json:"NumberColor3"`
			Rulesexplanation        string `json:"RulesExplanation"`
		} `json:"7"`
		Field8 struct {
			Paintcaravailable       bool   `json:"PaintCarAvailable"`
			Color1                  string `json:"Color1"`
			Color2                  string `json:"Color2"`
			Color3                  string `json:"Color3"`
			Sponsor1available       bool   `json:"Sponsor1Available"`
			Sponsor2available       bool   `json:"Sponsor2Available"`
			Sponsor1                string `json:"Sponsor1"`
			Sponsor2                string `json:"Sponsor2"`
			Paintwheelavailable     bool   `json:"PaintWheelAvailable"`
			Wheelcolor              string `json:"WheelColor"`
			Rimtypeavailable        bool   `json:"RimTypeAvailable"`
			Rimtype                 string `json:"RimType"`
			Allownumberfontchanges  bool   `json:"AllowNumberFontChanges"`
			Numberfont              string `json:"NumberFont"`
			Allownumbercolorchanges bool   `json:"AllowNumberColorChanges"`
			Numbercolor1            string `json:"NumberColor1"`
			Numbercolor2            string `json:"NumberColor2"`
			Numbercolor3            string `json:"NumberColor3"`
			Rulesexplanation        string `json:"RulesExplanation"`
		} `json:"8"`
		Field9 struct {
			Paintcaravailable       bool   `json:"PaintCarAvailable"`
			Color1                  string `json:"Color1"`
			Color2                  string `json:"Color2"`
			Color3                  string `json:"Color3"`
			Sponsor1available       bool   `json:"Sponsor1Available"`
			Sponsor2available       bool   `json:"Sponsor2Available"`
			Sponsor1                string `json:"Sponsor1"`
			Sponsor2                string `json:"Sponsor2"`
			Paintwheelavailable     bool   `json:"PaintWheelAvailable"`
			Wheelcolor              string `json:"WheelColor"`
			Rimtypeavailable        bool   `json:"RimTypeAvailable"`
			Rimtype                 string `json:"RimType"`
			Allownumberfontchanges  bool   `json:"AllowNumberFontChanges"`
			Numberfont              string `json:"NumberFont"`
			Allownumbercolorchanges bool   `json:"AllowNumberColorChanges"`
			Numbercolor1            string `json:"NumberColor1"`
			Numbercolor2            string `json:"NumberColor2"`
			Numbercolor3            string `json:"NumberColor3"`
			Rulesexplanation        string `json:"RulesExplanation"`
		} `json:"9"`
	} `json:"paint_rules"`
	SiteUrl string `json:"site_url"`
}

func (api *CarApi) GetCar() (*CarGetResponse, error) {
	url := "/data/car/get"

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &CarGetResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
