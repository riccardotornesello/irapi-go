package profile

import (
	"github.com/riccardotornesello/irapi-go/pkg/types"
)

type MemberProfileParams struct {
	CustId *int `url:"cust_id,omitempty,comma"`
}

type MemberProfileResponse struct {
	RecentAwards   []RecentAward `json:"recent_awards"`
	Activity       Activity      `json:"activity"`
	Success        bool          `json:"success"`
	ImageURL       string        `json:"image_url"`
	MemberInfo     MemberInfo    `json:"member_info"`
	Disabled       bool          `json:"disabled"`
	LicenseHistory []License     `json:"license_history"`
	RecentEvents   []RecentEvent `json:"recent_events"`
	CustID         int64         `json:"cust_id"`
	IsGenericImage bool          `json:"is_generic_image"`
	FollowCounts   FollowCounts  `json:"follow_counts"`
}

type Activity struct {
	Recent30DaysCount    int64 `json:"recent_30days_count"`
	Prev30DaysCount      int64 `json:"prev_30days_count"`
	ConsecutiveWeeks     int64 `json:"consecutive_weeks"`
	MostConsecutiveWeeks int64 `json:"most_consecutive_weeks"`
}

type FollowCounts struct {
	Followers int64 `json:"followers"`
	Follows   int64 `json:"follows"`
}

type License struct {
	CategoryID    int64   `json:"category_id"`
	Category      string  `json:"category"`
	CategoryName  string  `json:"category_name"`
	LicenseLevel  int64   `json:"license_level"`
	SafetyRating  float64 `json:"safety_rating"`
	Cpi           float64 `json:"cpi"`
	Irating       int64   `json:"irating"`
	TtRating      int64   `json:"tt_rating"`
	Color         string  `json:"color"`
	GroupName     string  `json:"group_name"`
	GroupID       int64   `json:"group_id"`
	Seq           int64   `json:"seq"`
	MprNumRaces   *int64  `json:"mpr_num_races,omitempty"`
	ProPromotable *bool   `json:"pro_promotable,omitempty"`
	MprNumTTS     *int64  `json:"mpr_num_tts,omitempty"`
}

type MemberInfo struct {
	AI             bool           `json:"ai"`
	Country        string         `json:"country"`
	CountryCode    string         `json:"country_code"`
	CustID         int64          `json:"cust_id"`
	DisplayName    string         `json:"display_name"`
	FlairID        int64          `json:"flair_id"`
	FlairName      string         `json:"flair_name"`
	FlairShortname string         `json:"flair_shortname"`
	Helmet         Helmet         `json:"helmet"`
	LastLogin      types.DateTime `json:"last_login"`
	Licenses       []License      `json:"licenses"`
	MemberSince    string         `json:"member_since"`
}

type Helmet struct {
	Pattern    int64  `json:"pattern"`
	Color1     string `json:"color1"`
	Color2     string `json:"color2"`
	Color3     string `json:"color3"`
	FaceType   int64  `json:"face_type"`
	HelmetType int64  `json:"helmet_type"`
}

type RecentAward struct {
	MemberAwardID      int64  `json:"member_award_id"`
	AwardID            int64  `json:"award_id"`
	Achievement        bool   `json:"achievement"`
	AwardCount         int64  `json:"award_count"`
	AwardDate          string `json:"award_date"`
	AwardOrder         int64  `json:"award_order"`
	AwardedDescription string `json:"awarded_description"`
	CustID             int64  `json:"cust_id"`
	Description        string `json:"description"`
	GroupName          string `json:"group_name"`
	HasPDF             bool   `json:"has_pdf"`
	IconURLLarge       string `json:"icon_url_large"`
	IconURLSmall       string `json:"icon_url_small"`
	IconURLUnawarded   string `json:"icon_url_unawarded"`
	Name               string `json:"name"`
	SubsessionID       int64  `json:"subsession_id"`
	Viewed             bool   `json:"viewed"`
	Weight             int64  `json:"weight"`
}

type RecentEvent struct {
	EventType        string         `json:"event_type"`
	SubsessionID     int64          `json:"subsession_id"`
	StartTime        types.DateTime `json:"start_time"`
	EventID          int64          `json:"event_id"`
	EventName        string         `json:"event_name"`
	SimsessionType   int64          `json:"simsession_type"`
	StartingPosition int64          `json:"starting_position"`
	FinishPosition   int64          `json:"finish_position"`
	BestLapTime      int64          `json:"best_lap_time"`
	PercentRank      int64          `json:"percent_rank"`
	CarID            int64          `json:"car_id"`
	CarName          string         `json:"car_name"`
	LogoURL          interface{}    `json:"logo_url"`
	Track            Track          `json:"track"`
}

type Track struct {
	ConfigName string `json:"config_name"`
	TrackID    int64  `json:"track_id"`
	TrackName  string `json:"track_name"`
}
