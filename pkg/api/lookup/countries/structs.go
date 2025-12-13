package countries

type LookupCountriesResponse []LookupCountriesResponseElement

type LookupCountriesResponseElement struct {
	CountryName string `json:"country_name"`
	CountryCode string `json:"country_code"`
}
