package drivers

type LookupDriversParams struct {
	SearchTerm string `url:"search_term,comma"`
	LeagueId   *int   `url:"league_id,omitempty,comma"`
}

type LookupDriversResponse []LookupDriversResponseElement

type LookupDriversResponseElement struct {
	CustID          int64  `json:"cust_id"`
	DisplayName     string `json:"display_name"`
	Helmet          Helmet `json:"helmet"`
	ProfileDisabled bool   `json:"profile_disabled"`
}

type Helmet struct {
	Pattern    int64  `json:"pattern"`
	Color1     string `json:"color1"`
	Color2     string `json:"color2"`
	Color3     string `json:"color3"`
	FaceType   int64  `json:"face_type"`
	HelmetType int64  `json:"helmet_type"`
}
