package member

import (
	"encoding/json"
)

type MemberGetResponse struct {
	Success bool     `json:"success"`
	CustIDS []int64  `json:"cust_ids"`
	Members []Member `json:"members"`
}

import "time"

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
}

type Helmet struct {
	Pattern    int64  `json:"pattern"`
	Color1     string `json:"color1"`
	Color2     string `json:"color2"`
	Color3     string `json:"color3"`
	FaceType   int64  `json:"face_type"`
	HelmetType int64  `json:"helmet_type"`
}


func (api *MemberApi) Get() (*MemberGetResponse, error) {
	return api.GetJson[MemberGetResponse]("/data/member/get")
}