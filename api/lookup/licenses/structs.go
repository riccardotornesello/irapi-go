package licenses

import ()

type LookupLicensesResponse []LookupLicensesResponseElement

type LookupLicensesResponseElement struct {
	LicenseGroup         int64   `json:"license_group"`
	GroupName            string  `json:"group_name"`
	MinNumRaces          *int64  `json:"min_num_races"`
	ParticipationCredits int64   `json:"participation_credits"`
	MinSrToFastTrack     *int64  `json:"min_sr_to_fast_track"`
	Levels               []Level `json:"levels"`
	MinNumTt             *int64  `json:"min_num_tt"`
}

type Level struct {
	LicenseID     int64  `json:"license_id"`
	LicenseGroup  int64  `json:"license_group"`
	License       string `json:"license"`
	ShortName     string `json:"short_name"`
	LicenseLetter string `json:"license_letter"`
	Color         string `json:"color"`
}
