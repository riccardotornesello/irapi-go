package info

import (
	"time"
)

type MemberInfoResponse struct {
	CustID               int64        `json:"cust_id"`
	DisplayName          string       `json:"display_name"`
	FirstName            string       `json:"first_name"`
	LastName             string       `json:"last_name"`
	OnCarName            string       `json:"on_car_name"`
	MemberSince          string       `json:"member_since"`
	FlairID              int64        `json:"flair_id"`
	FlairName            string       `json:"flair_name"`
	FlairShortname       string       `json:"flair_shortname"`
	FlairCountryCode     string       `json:"flair_country_code"`
	LastLogin            time.Time    `json:"last_login"`
	ReadTc               time.Time    `json:"read_tc"`
	ReadPp               time.Time    `json:"read_pp"`
	ReadCompRules        time.Time    `json:"read_comp_rules"`
	Flags                int64        `json:"flags"`
	ConnectionType       string       `json:"connection_type"`
	DownloadServer       string       `json:"download_server"`
	Account              Account      `json:"account"`
	Helmet               Helmet       `json:"helmet"`
	Suit                 Suit         `json:"suit"`
	Licenses             Licenses     `json:"licenses"`
	CarPackages          []Package    `json:"car_packages"`
	TrackPackages        []Package    `json:"track_packages"`
	OtherOwnedPackages   []int64      `json:"other_owned_packages"`
	Dev                  bool         `json:"dev"`
	AlphaTester          bool         `json:"alpha_tester"`
	RainTester           bool         `json:"rain_tester"`
	Broadcaster          bool         `json:"broadcaster"`
	Restrictions         Restrictions `json:"restrictions"`
	HasReadCompRules     bool         `json:"has_read_comp_rules"`
	HasReadNda           bool         `json:"has_read_nda"`
	FlagsHex             string       `json:"flags_hex"`
	HundredPctClub       bool         `json:"hundred_pct_club"`
	TwentyPctDiscount    bool         `json:"twenty_pct_discount"`
	LastSeason           int64        `json:"last_season"`
	HasAdditionalContent bool         `json:"has_additional_content"`
	HasReadTc            bool         `json:"has_read_tc"`
	HasReadPp            bool         `json:"has_read_pp"`
}

type Account struct {
	IRDollars    int64       `json:"ir_dollars"`
	IRCredits    int64       `json:"ir_credits"`
	Status       string      `json:"status"`
	CountryRules interface{} `json:"country_rules"`
}

type Package struct {
	PackageID  int64   `json:"package_id"`
	ContentIDS []int64 `json:"content_ids"`
}

type Helmet struct {
	Pattern    int64  `json:"pattern"`
	Color1     string `json:"color1"`
	Color2     string `json:"color2"`
	Color3     string `json:"color3"`
	FaceType   int64  `json:"face_type"`
	HelmetType int64  `json:"helmet_type"`
}

type Licenses struct {
	Oval       DirtOval `json:"oval"`
	SportsCar  DirtOval `json:"sports_car"`
	FormulaCar DirtOval `json:"formula_car"`
	DirtOval   DirtOval `json:"dirt_oval"`
	DirtRoad   DirtOval `json:"dirt_road"`
}

type DirtOval struct {
	CategoryID    int64   `json:"category_id"`
	Category      string  `json:"category"`
	CategoryName  string  `json:"category_name"`
	LicenseLevel  int64   `json:"license_level"`
	SafetyRating  float64 `json:"safety_rating"`
	Cpi           float64 `json:"cpi"`
	Irating       int64   `json:"irating"`
	TtRating      int64   `json:"tt_rating"`
	MprNumRaces   int64   `json:"mpr_num_races"`
	Color         string  `json:"color"`
	GroupName     string  `json:"group_name"`
	GroupID       int64   `json:"group_id"`
	ProPromotable bool    `json:"pro_promotable"`
	Seq           int64   `json:"seq"`
	MprNumTTS     int64   `json:"mpr_num_tts"`
}

type Restrictions struct {
}

type Suit struct {
	Pattern  int64  `json:"pattern"`
	Color1   string `json:"color1"`
	Color2   string `json:"color2"`
	Color3   string `json:"color3"`
	BodyType int64  `json:"body_type"`
}
