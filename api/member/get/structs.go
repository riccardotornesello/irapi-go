package get

import (
	"time"
)

type MemberGetParams struct {
	CustIds         []int `url:"cust_ids,omitempty,comma"`
	IncludeLicenses bool  `url:"include_licenses,omitempty,comma"`
}

type MemberGetResponse struct {
	Success bool     `json:"success"`
	CustIDS []int64  `json:"cust_ids"`
	Members []Member `json:"members"`
}

type Member struct {
	CustID         int64     `json:"cust_id"`
	DisplayName    string    `json:"display_name"`
	Helmet         Helmet    `json:"helmet"`
	LastLogin      time.Time `json:"last_login"`
	MemberSince    string    `json:"member_since"`
	FlairID        int64     `json:"flair_id"`
	FlairName      string    `json:"flair_name"`
	FlairShortname string    `json:"flair_shortname"`
	AI             bool      `json:"ai"`
	Licenses       []License `json:"licenses,omitempty"`
}

type Helmet struct {
	Pattern    int64  `json:"pattern"`
	Color1     string `json:"color1"`
	Color2     string `json:"color2"`
	Color3     string `json:"color3"`
	FaceType   int64  `json:"face_type"`
	HelmetType int64  `json:"helmet_type"`
}

type License struct {
	CategoryID    int64   `json:"category_id"`
	Category      string  `json:"category"`
	CategoryName  string  `json:"category_name"`
	LicenseLevel  int64   `json:"license_level"`
	SafetyRating  float64 `json:"safety_rating"`
	Cpi           float64 `json:"cpi"`
	TtRating      int64   `json:"tt_rating"`
	MprNumRaces   int64   `json:"mpr_num_races"`
	Color         string  `json:"color"`
	GroupName     string  `json:"group_name"`
	GroupID       int64   `json:"group_id"`
	ProPromotable bool    `json:"pro_promotable"`
	Seq           int64   `json:"seq"`
	MprNumTTS     int64   `json:"mpr_num_tts"`
	Irating       *int64  `json:"irating,omitempty"`
}
