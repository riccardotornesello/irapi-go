package member

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
)

type MemberProfileParams struct {
	CustId *int `url:"cust_id,omitempty"` // Defaults to the authenticated member.
}

type MemberProfileResponse struct {
	RecentAwards []struct {
		MemberAwardId      int    `json:"member_award_id"`
		AwardId            int    `json:"award_id"`
		Achievement        bool   `json:"achievement"`
		AwardCount         int    `json:"award_count"`
		AwardDate          string `json:"award_date"`
		AwardOrder         int    `json:"award_order"`
		AwardedDescription string `json:"awarded_description"`
		Description        string `json:"description"`
		GroupName          string `json:"group_name"`
		HasPdf             bool   `json:"has_pdf"`
		IconUrlLarge       string `json:"icon_url_large"`
		IconUrlSmall       string `json:"icon_url_small"`
		IconUrlUnawarded   string `json:"icon_url_unawarded"`
		Name               string `json:"name"`
		SubsessionId       int    `json:"subsession_id"`
		Viewed             bool   `json:"viewed"`
		Weight             int    `json:"weight"`
	} `json:"recent_awards"`
	Activity struct {
		Recent30daysCount    int `json:"recent_30days_count"`
		Prev30daysCount      int `json:"prev_30days_count"`
		ConsecutiveWeeks     int `json:"consecutive_weeks"`
		MostConsecutiveWeeks int `json:"most_consecutive_weeks"`
	} `json:"activity"`
	Success    bool   `json:"success"`
	ImageUrl   string `json:"image_url"`
	MemberInfo struct {
		CustId      int    `json:"cust_id"`
		DisplayName string `json:"display_name"`
		Helmet      struct {
			Pattern    int    `json:"pattern"`
			Color1     string `json:"color1"`
			Color2     string `json:"color2"`
			Color3     string `json:"color3"`
			FaceType   int    `json:"face_type"`
			HelmetType int    `json:"helmet_type"`
		} `json:"helmet"`
		LastLogin   string `json:"last_login"`
		MemberSince string `json:"member_since"`
		ClubId      int    `json:"club_id"`
		ClubName    string `json:"club_name"`
		Ai          bool   `json:"ai"`
		Licenses    []struct {
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
		} `json:"licenses"`
		Country     string `json:"country"`
		CountryCode string `json:"country_code"`
	} `json:"member_info"`
	Disabled       bool `json:"disabled"`
	LicenseHistory []struct {
		CategoryId   int     `json:"category_id"`
		Category     string  `json:"category"`
		CategoryName string  `json:"category_name"`
		LicenseLevel int     `json:"license_level"`
		SafetyRating float64 `json:"safety_rating"`
		Cpi          float64 `json:"cpi"`
		Irating      int     `json:"irating"`
		TtRating     int     `json:"tt_rating"`
		Color        string  `json:"color"`
		GroupName    string  `json:"group_name"`
		GroupId      int     `json:"group_id"`
		Seq          int     `json:"seq"`
	} `json:"license_history"`
	RecentEvents []struct {
		EventType        string  `json:"event_type"`
		SubsessionId     int     `json:"subsession_id"`
		StartTime        string  `json:"start_time"`
		EventId          int     `json:"event_id"`
		EventName        string  `json:"event_name"`
		SimsessionType   int     `json:"simsession_type"`
		StartingPosition int     `json:"starting_position"`
		FinishPosition   int     `json:"finish_position"`
		BestLapTime      int     `json:"best_lap_time"`
		PercentRank      int     `json:"percent_rank"`
		CarId            int     `json:"car_id"`
		CarName          string  `json:"car_name"`
		LogoUrl          *string `json:"logo_url"`
		Track            struct {
			ConfigName string `json:"config_name"`
			TrackId    int    `json:"track_id"`
			TrackName  string `json:"track_name"`
		} `json:"track"`
	} `json:"recent_events"`
	CustId         int  `json:"cust_id"`
	IsGenericImage bool `json:"is_generic_image"`
	FollowCounts   struct {
		Followers int `json:"followers"`
		Follows   int `json:"follows"`
	} `json:"follow_counts"`
}

func (api *MemberApi) GetMemberProfile(params MemberProfileParams) (*MemberProfileResponse, error) {
	paramsString, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	url := "/data/member/profile?" + paramsString.Encode()

	respBody, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	response := &MemberProfileResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
