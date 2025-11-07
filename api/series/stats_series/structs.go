package stats_series

import ()

type SeriesStatsSeriesResponse []SeriesStatsSeriesResponseElement

type SeriesStatsSeriesResponseElement struct {
	SeriesID          int64              `json:"series_id"`
	SeriesName        string             `json:"series_name"`
	SeriesShortName   string             `json:"series_short_name"`
	CategoryID        int64              `json:"category_id"`
	Category          Category           `json:"category"`
	Active            bool               `json:"active"`
	Official          bool               `json:"official"`
	FixedSetup        bool               `json:"fixed_setup"`
	Logo              *string            `json:"logo"`
	LicenseGroup      int64              `json:"license_group"`
	LicenseGroupTypes []LicenseGroupType `json:"license_group_types"`
	AllowedLicenses   []AllowedLicense   `json:"allowed_licenses"`
	Seasons           []Season           `json:"seasons"`
	SearchFilters     *string            `json:"search_filters,omitempty"`
}

type AllowedLicense struct {
	GroupName       GroupName `json:"group_name"`
	LicenseGroup    int64     `json:"license_group"`
	MaxLicenseLevel int64     `json:"max_license_level"`
	MinLicenseLevel int64     `json:"min_license_level"`
	ParentID        int64     `json:"parent_id"`
}

type LicenseGroupType struct {
	LicenseGroupType int64 `json:"license_group_type"`
}

type Season struct {
	SeasonID          int64              `json:"season_id"`
	SeriesID          int64              `json:"series_id"`
	SeasonName        string             `json:"season_name"`
	SeasonShortName   string             `json:"season_short_name"`
	SeasonYear        int64              `json:"season_year"`
	SeasonQuarter     int64              `json:"season_quarter"`
	Active            bool               `json:"active"`
	Official          bool               `json:"official"`
	DriverChanges     bool               `json:"driver_changes"`
	FixedSetup        bool               `json:"fixed_setup"`
	LicenseGroup      int64              `json:"license_group"`
	HasSupersessions  bool               `json:"has_supersessions"`
	CarSwitching      bool               `json:"car_switching"`
	LicenseGroupTypes []LicenseGroupType `json:"license_group_types"`
	CarClasses        []CarClass         `json:"car_classes"`
	RaceWeeks         []RaceWeek         `json:"race_weeks"`
}

type CarClass struct {
	CarClassID    int64  `json:"car_class_id"`
	ShortName     string `json:"short_name"`
	Name          string `json:"name"`
	RelativeSpeed int64  `json:"relative_speed"`
}

type RaceWeek struct {
	SeasonID    int64 `json:"season_id"`
	RaceWeekNum int64 `json:"race_week_num"`
	Track       Track `json:"track"`
}

type Track struct {
	ConfigName *string   `json:"config_name,omitempty"`
	TrackID    int64     `json:"track_id"`
	TrackName  TrackName `json:"track_name"`
}

type GroupName string

const (
	ClassA GroupName = "Class A"
	ClassB GroupName = "Class B"
	ClassC GroupName = "Class C"
	ClassD GroupName = "Class D"
	Pro    GroupName = "Pro"
	ProWC  GroupName = "Pro/WC"
	Rookie GroupName = "Rookie"
)

type Category string

const (
	DirtOval   Category = "dirt_oval"
	DirtRoad   Category = "dirt_road"
	FormulaCar Category = "formula_car"
	Oval       Category = "oval"
	Road       Category = "road"
	SportsCar  Category = "sports_car"
)

type TrackName string

const (
	AlgarveInternationalCircuit             TrackName = "Algarve International Circuit"
	AutoClubSpeedway                        TrackName = "Auto Club Speedway"
	AutodromoInternazionaleDelMugello       TrackName = "Autodromo Internazionale del Mugello"
	AutodromoInternazionaleEnzoEDinoFerrari TrackName = "Autodromo Internazionale Enzo e Dino Ferrari"
	AutodromoNazionaleMonza                 TrackName = "Autodromo Nazionale Monza"
	AutódromoHermanosRodríguez              TrackName = "Autódromo Hermanos Rodríguez"
	AutódromoJoséCarlosPace                 TrackName = "Autódromo José Carlos Pace"
	BarberMotorsportsPark                   TrackName = "Barber Motorsports Park"
	BarkRiverInternationalRaceway           TrackName = "Bark River International Raceway"
	BrandsHatchCircuit                      TrackName = "Brands Hatch Circuit"
	BristolMotorSpeedway                    TrackName = "Bristol Motor Speedway"
	CadwellParkCircuit                      TrackName = "Cadwell Park Circuit"
	CanadianTireMotorsportsPark             TrackName = "Canadian Tire Motorsports Park"
	CedarLakeSpeedway                       TrackName = "Cedar Lake Speedway"
	CentripetalCircuit                      TrackName = "Centripetal Circuit"
	CharlotteMotorSpeedway                  TrackName = "Charlotte Motor Speedway"
	ChicagoStreetCourse                     TrackName = "Chicago Street Course"
	ChicagolandSpeedway                     TrackName = "Chicagoland Speedway"
	ChiliBowl                               TrackName = "Chili Bowl"
	CircuitDES24HeuresDuMans                TrackName = "Circuit des 24 Heures du Mans"
	CircuitDeBarcelonaCatalunya             TrackName = "Circuit de Barcelona Catalunya"
	CircuitDeLédenon                        TrackName = "Circuit de Lédenon"
	CircuitDeNeversMagnyCours               TrackName = "Circuit de Nevers Magny-Cours"
	CircuitDeSPAFrancorchamps               TrackName = "Circuit de Spa-Francorchamps"
	CircuitGillesVilleneuve                 TrackName = "Circuit Gilles Villeneuve"
	CircuitOfTheAmericas                    TrackName = "Circuit of the Americas"
	CircuitParkZandvoort                    TrackName = "Circuit Park Zandvoort"
	CircuitZandvoort                        TrackName = "Circuit Zandvoort"
	CircuitZolder                           TrackName = "Circuit Zolder"
	CircuitoDeJerezÁngelNieto               TrackName = "Circuito de Jerez - Ángel Nieto"
	CircuitoDeNavarra                       TrackName = "Circuito de Navarra"
	ConcordSpeedway                         TrackName = "Concord Speedway"
	CrandonInternationalRaceway             TrackName = "Crandon International Raceway"
	DarlingtonRaceway                       TrackName = "Darlington Raceway"
	DaytonaInternationalSpeedway            TrackName = "Daytona International Speedway"
	DaytonaRallycrossAndDirtRoad            TrackName = "Daytona Rallycross and Dirt Road"
	DetroitGrandPrixAtBelleIsle             TrackName = "Detroit Grand Prix at Belle Isle"
	DoningtonParkRacingCircuit              TrackName = "Donington Park Racing Circuit"
	DoverMotorSpeedway                      TrackName = "Dover Motor Speedway"
	EchoParkSpeedwayAtlanta                 TrackName = "EchoPark Speedway (Atlanta)"
	EldoraSpeedway                          TrackName = "Eldora Speedway"
	FairburySpeedway                        TrackName = "Fairbury Speedway"
	FederatedAutoPartsRacewayAtI55          TrackName = "Federated Auto Parts Raceway at I-55"
	FirebirdMotorsportsPark                 TrackName = "Firebird Motorsports Park"
	FiveFlagsSpeedway                       TrackName = "Five Flags Speedway"
	FujiInternationalSpeedway               TrackName = "Fuji International Speedway"
	HickoryMotorSpeedway                    TrackName = "Hickory Motor Speedway"
	HockenheimringBadenWürttemberg          TrackName = "Hockenheimring Baden-Württemberg"
	HomesteadMiamiSpeedway                  TrackName = "Homestead Miami Speedway"
	Hungaroring                             TrackName = "Hungaroring"
	HusetSSpeedway                          TrackName = "Huset's Speedway"
	IRacingSuperspeedway                    TrackName = "iRacing Superspeedway"
	IndianapolisMotorSpeedway               TrackName = "Indianapolis Motor Speedway"
	IowaSpeedway                            TrackName = "Iowa Speedway"
	IowaSpeedwayOval2011                    TrackName = "Iowa Speedway - Oval - 2011"
	IrwindaleSpeedway                       TrackName = "Irwindale Speedway"
	KansasSpeedway                          TrackName = "Kansas Speedway"
	KentuckySpeedway                        TrackName = "Kentucky Speedway"
	KevinHarvickSKernRaceway                TrackName = "Kevin Harvick's Kern Raceway"
	KnockhillRacingCircuit                  TrackName = "Knockhill Racing Circuit"
	KnoxvilleRaceway                        TrackName = "Knoxville Raceway"
	KokomoSpeedway                          TrackName = "Kokomo Speedway"
	LAColiseumRaceway                       TrackName = "LA Coliseum Raceway"
	LangleySpeedway                         TrackName = "Langley Speedway"
	LanierNationalSpeedway                  TrackName = "Lanier National Speedway"
	LasVegasMotorSpeedway                   TrackName = "Las Vegas Motor Speedway"
	LegacyCharlotteMotorSpeedway2008        TrackName = "[Legacy] Charlotte Motor Speedway - 2008"
	LegacyKentuckySpeedway2011              TrackName = "[Legacy] Kentucky Speedway - 2011"
	LegacyMichiganInternationalSpeedway2009 TrackName = "[Legacy] Michigan International Speedway - 2009"
	LegacyPhoenixRaceway2008                TrackName = "[Legacy] Phoenix Raceway - 2008"
	LegacyPoconoRaceway2009                 TrackName = "[Legacy] Pocono Raceway - 2009"
	LegacySilverstoneCircuit2008            TrackName = "[Legacy] Silverstone Circuit - 2008"
	LegacyTexasMotorSpeedway2009            TrackName = "[Legacy] Texas Motor Speedway - 2009"
	LernervilleSpeedway                     TrackName = "Lernerville Speedway"
	LimalandMotorsportsPark                 TrackName = "Limaland Motorsports Park"
	LimeRockPark                            TrackName = "Lime Rock Park"
	LincolnSpeedway                         TrackName = "Lincoln Speedway"
	LongBeachStreetCircuit                  TrackName = "Long Beach Street Circuit"
	LucasOilIndianapolisRacewayPark         TrackName = "Lucas Oil Indianapolis Raceway Park"
	LucasOilSpeedway                        TrackName = "Lucas Oil Speedway"
	LånkebanenHellRX                        TrackName = "Lånkebanen (Hell RX)"
	MartinsvilleSpeedway                    TrackName = "Martinsville Speedway"
	MichiganInternationalSpeedway           TrackName = "Michigan International Speedway"
	MidOhioSportsCarCourse                  TrackName = "Mid-Ohio Sports Car Course"
	MillbridgeSpeedway                      TrackName = "Millbridge Speedway"
	MisanoWorldCircuitMarcoSimoncelli       TrackName = "Misano World Circuit Marco Simoncelli"
	MobilityResortMotegi                    TrackName = "Mobility Resort Motegi"
	MotorLandAragón                         TrackName = "MotorLand Aragón"
	MotorsportArenaOschersleben             TrackName = "Motorsport Arena Oschersleben"
	MountPanoramaCircuit                    TrackName = "Mount Panorama Circuit"
	MyrtleBeachSpeedway                     TrackName = "Myrtle Beach Speedway"
	NashvilleFairgroundsSpeedway            TrackName = "Nashville Fairgrounds Speedway"
	NashvilleSuperspeedway                  TrackName = "Nashville Superspeedway"
	NewHampshireMotorSpeedway               TrackName = "New Hampshire Motor Speedway"
	NewSmyrnaSpeedway                       TrackName = "New Smyrna Speedway"
	NorthWilkesboroSpeedway                 TrackName = "North Wilkesboro Speedway"
	NürburgringCombined                     TrackName = "Nürburgring Combined"
	NürburgringGrandPrixStrecke             TrackName = "Nürburgring Grand-Prix-Strecke"
	NürburgringNordschleife                 TrackName = "Nürburgring Nordschleife"
	OkayamaInternationalCircuit             TrackName = "Okayama International Circuit"
	OranParkRaceway                         TrackName = "Oran Park Raceway"
	OswegoSpeedway                          TrackName = "Oswego Speedway"
	OultonParkCircuit                       TrackName = "Oulton Park Circuit"
	OxfordPlainsSpeedway                    TrackName = "Oxford Plains Speedway"
	PhillipIslandCircuit                    TrackName = "Phillip Island Circuit"
	PhoenixRaceway                          TrackName = "Phoenix Raceway"
	PoconoRaceway                           TrackName = "Pocono Raceway"
	PortRoyalSpeedway                       TrackName = "Port Royal Speedway"
	PortlandInternationalRaceway            TrackName = "Portland International Raceway"
	RedBullRing                             TrackName = "Red Bull Ring"
	RetiredCharlotteMotorSpeedway           TrackName = "[Retired] Charlotte Motor Speedway"
	RetiredLimeRockPark2008                 TrackName = "[Retired] Lime Rock Park - 2008"
	RetiredOxfordPlainsSpeedway             TrackName = "[Retired] Oxford Plains Speedway"
	RetiredPhoenixRaceway                   TrackName = "[Retired] Phoenix Raceway"
	RetiredVirginiaInternationalRaceway     TrackName = "[Retired] Virginia International Raceway"
	RetiredWatkinsGlenInternational         TrackName = "[Retired] Watkins Glen International"
	RichmondRaceway                         TrackName = "Richmond Raceway"
	RoadAmerica                             TrackName = "Road America"
	RoadAtlanta                             TrackName = "Road Atlanta"
	RockinghamSpeedway                      TrackName = "Rockingham Speedway"
	RudskogenMotorsenter                    TrackName = "Rudskogen Motorsenter"
	Sachsenring                             TrackName = "Sachsenring"
	SandownInternationalMotorRaceway        TrackName = "Sandown International Motor Raceway"
	SebringInternationalRaceway             TrackName = "Sebring International Raceway"
	ShellVPowerMotorsportParkAtTheBend      TrackName = "Shell V-Power Motorsport Park at The Bend"
	SilverstoneCircuit                      TrackName = "Silverstone Circuit"
	SlingerSpeedway                         TrackName = "Slinger Speedway"
	SnettertonCircuit                       TrackName = "Snetterton Circuit"
	SonomaRaceway                           TrackName = "Sonoma Raceway"
	SouthBostonSpeedway                     TrackName = "South Boston Speedway"
	SouthernNationalMotorsportsPark         TrackName = "Southern National Motorsports Park"
	StaffordMotorSpeedway                   TrackName = "Stafford Motor Speedway"
	SummitPointRaceway                      TrackName = "Summit Point Raceway"
	SuzukaInternationalRacingCourse         TrackName = "Suzuka International Racing Course"
	TalladegaSuperspeedway                  TrackName = "Talladega Superspeedway"
	TexasMotorSpeedway                      TrackName = "Texas Motor Speedway"
	TheBullring                             TrackName = "The Bullring"
	TheDirtTrackAtCharlotte                 TrackName = "The Dirt Track at Charlotte"
	TheMilwaukeeMile                        TrackName = "The Milwaukee Mile"
	ThompsonSpeedwayMotorsportsPark         TrackName = "Thompson Speedway Motorsports Park"
	ThruxtonCircuit                         TrackName = "Thruxton Circuit"
	TrackNameCharlotteMotorSpeedway         TrackName = "Charlotte Motor Speedway "
	TsukubaCircuit                          TrackName = "Tsukuba Circuit"
	USAInternationalSpeedway                TrackName = "USA International Speedway"
	VirginiaInternationalRaceway            TrackName = "Virginia International Raceway"
	VolusiaSpeedwayPark                     TrackName = "Volusia Speedway Park"
	WatkinsGlenInternational                TrackName = "Watkins Glen International"
	WeatherTechRacewayAtLagunaSeca          TrackName = "WeatherTech Raceway at Laguna Seca"
	WeedsportSpeedway                       TrackName = "Weedsport Speedway"
	WildWestMotorsportsPark                 TrackName = "Wild West Motorsports Park"
	WilliamsGroveSpeedway                   TrackName = "Williams Grove Speedway"
	WillowSpringsInternationalRaceway       TrackName = "Willow Springs International Raceway"
	WintonMotorRaceway                      TrackName = "Winton Motor Raceway"
	WorldWideTechnologyRacewayGateway       TrackName = "World Wide Technology Raceway (Gateway)"
)
