package formula_car

type DriverStatsByCategoryFormulaCarResponse struct {
	AvgFinishPos  int64       `json:"AVG_FINISH_POS"`
	AvgInc        float64     `json:"AVG_INC"`
	AvgPoints     int64       `json:"AVG_POINTS"`
	AvgStartPos   int64       `json:"AVG_START_POS"`
	Champpoints   int64       `json:"CHAMPPOINTS"`
	Class         string      `json:"CLASS"`
	ClubName      interface{} `json:"CLUB_NAME"`
	Custid        int64       `json:"CUSTID"`
	Driver        string      `json:"DRIVER"`
	Irating       float64     `json:"IRATING"`
	Laps          int64       `json:"LAPS"`
	Lapslead      int64       `json:"LAPSLEAD"`
	Location      *string     `json:"LOCATION,omitempty"`
	Starts        int64       `json:"STARTS"`
	Top25Pcnt     int64       `json:"TOP25PCNT"`
	TotClubpoints int64       `json:"TOT_CLUBPOINTS"`
	Ttrating      float64     `json:"TTRATING"`
	WINS          int64       `json:"WINS"`
}
