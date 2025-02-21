package member

import (
	"encoding/json"
)

type MemberInfoResponse struct {
	CustId         int    `json:"cust_id"`
	DisplayName    string `json:"display_name"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	OnCarName      string `json:"on_car_name"`
	MemberSince    string `json:"member_since"`
	ClubId         int    `json:"club_id"`
	ClubName       string `json:"club_name"`
	LastLogin      string `json:"last_login"`
	ReadTc         string `json:"read_tc"`
	ReadPp         string `json:"read_pp"`
	ReadCompRules  string `json:"read_comp_rules"`
	Flags          int    `json:"flags"`
	ConnectionType string `json:"connection_type"`
	DownloadServer string `json:"download_server"`
	Account        struct {
		IrDollars    float64     `json:"ir_dollars"`
		IrCredits    int         `json:"ir_credits"`
		Status       string      `json:"status"`
		CountryRules interface{} `json:"country_rules"`
	} `json:"account"`
	Helmet struct {
		Pattern    int    `json:"pattern"`
		Color1     string `json:"color1"`
		Color2     string `json:"color2"`
		Color3     string `json:"color3"`
		FaceType   int    `json:"face_type"`
		HelmetType int    `json:"helmet_type"`
	} `json:"helmet"`
	Suit struct {
		Pattern  int    `json:"pattern"`
		Color1   string `json:"color1"`
		Color2   string `json:"color2"`
		Color3   string `json:"color3"`
		BodyType int    `json:"body_type"`
	} `json:"suit"`
	Licenses struct {
		Oval struct {
			CategoryId    int     `json:"category_id"`
			Category      string  `json:"category"`
			CategoryName  string  `json:"category_name"`
			LicenseLevel  int     `json:"license_level"`
			SafetyRating  float64 `json:"safety_rating"`
			Cpi           float64 `json:"cpi"`
			Irating       int     `json:"irating"`
			TtRating      int     `json:"tt_rating"`
			MprNumRaces   int     `json:"mpr_num_races"`
			Color         string  `json:"color"`
			GroupName     string  `json:"group_name"`
			GroupId       int     `json:"group_id"`
			ProPromotable bool    `json:"pro_promotable"`
			Seq           int     `json:"seq"`
			MprNumTts     int     `json:"mpr_num_tts"`
		} `json:"oval"`
		SportsCar struct {
			CategoryId    int     `json:"category_id"`
			Category      string  `json:"category"`
			CategoryName  string  `json:"category_name"`
			LicenseLevel  int     `json:"license_level"`
			SafetyRating  float64 `json:"safety_rating"`
			Cpi           float64 `json:"cpi"`
			Irating       int     `json:"irating"`
			TtRating      int     `json:"tt_rating"`
			MprNumRaces   int     `json:"mpr_num_races"`
			Color         string  `json:"color"`
			GroupName     string  `json:"group_name"`
			GroupId       int     `json:"group_id"`
			ProPromotable bool    `json:"pro_promotable"`
			Seq           int     `json:"seq"`
			MprNumTts     int     `json:"mpr_num_tts"`
		} `json:"sports_car"`
		FormulaCar struct {
			CategoryId    int     `json:"category_id"`
			Category      string  `json:"category"`
			CategoryName  string  `json:"category_name"`
			LicenseLevel  int     `json:"license_level"`
			SafetyRating  float64 `json:"safety_rating"`
			Cpi           float64 `json:"cpi"`
			Irating       int     `json:"irating"`
			TtRating      int     `json:"tt_rating"`
			MprNumRaces   int     `json:"mpr_num_races"`
			Color         string  `json:"color"`
			GroupName     string  `json:"group_name"`
			GroupId       int     `json:"group_id"`
			ProPromotable bool    `json:"pro_promotable"`
			Seq           int     `json:"seq"`
			MprNumTts     int     `json:"mpr_num_tts"`
		} `json:"formula_car"`
		DirtOval struct {
			CategoryId    int     `json:"category_id"`
			Category      string  `json:"category"`
			CategoryName  string  `json:"category_name"`
			LicenseLevel  int     `json:"license_level"`
			SafetyRating  float64 `json:"safety_rating"`
			Cpi           float64 `json:"cpi"`
			Irating       int     `json:"irating"`
			TtRating      int     `json:"tt_rating"`
			MprNumRaces   int     `json:"mpr_num_races"`
			Color         string  `json:"color"`
			GroupName     string  `json:"group_name"`
			GroupId       int     `json:"group_id"`
			ProPromotable bool    `json:"pro_promotable"`
			Seq           int     `json:"seq"`
			MprNumTts     int     `json:"mpr_num_tts"`
		} `json:"dirt_oval"`
		DirtRoad struct {
			CategoryId    int     `json:"category_id"`
			Category      string  `json:"category"`
			CategoryName  string  `json:"category_name"`
			LicenseLevel  int     `json:"license_level"`
			SafetyRating  float64 `json:"safety_rating"`
			Cpi           float64 `json:"cpi"`
			Irating       int     `json:"irating"`
			TtRating      int     `json:"tt_rating"`
			MprNumRaces   int     `json:"mpr_num_races"`
			Color         string  `json:"color"`
			GroupName     string  `json:"group_name"`
			GroupId       int     `json:"group_id"`
			ProPromotable bool    `json:"pro_promotable"`
			Seq           int     `json:"seq"`
			MprNumTts     int     `json:"mpr_num_tts"`
		} `json:"dirt_road"`
	} `json:"licenses"`
	CarPackages []struct {
		PackageId  int   `json:"package_id"`
		ContentIds []int `json:"content_ids"`
	} `json:"car_packages"`
	TrackPackages []struct {
		PackageId  int   `json:"package_id"`
		ContentIds []int `json:"content_ids"`
	} `json:"track_packages"`
	OtherOwnedPackages   []int       `json:"other_owned_packages"`
	Dev                  bool        `json:"dev"`
	AlphaTester          bool        `json:"alpha_tester"`
	RainTester           bool        `json:"rain_tester"`
	Broadcaster          bool        `json:"broadcaster"`
	Restrictions         interface{} `json:"restrictions"`
	HasReadCompRules     bool        `json:"has_read_comp_rules"`
	HasReadNda           bool        `json:"has_read_nda"`
	FlagsHex             string      `json:"flags_hex"`
	HundredPctClub       bool        `json:"hundred_pct_club"`
	TwentyPctDiscount    bool        `json:"twenty_pct_discount"`
	LastSeason           int         `json:"last_season"`
	HasAdditionalContent bool        `json:"has_additional_content"`
	HasReadTc            bool        `json:"has_read_tc"`
	HasReadPp            bool        `json:"has_read_pp"`
}

func (api *MemberApi) GetMemberInfo() (*MemberInfoResponse, error) {
	url := "https://members-ng.iracing.com/data/member/info"

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &MemberInfoResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
