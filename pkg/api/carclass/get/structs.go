package get

type CarclassGetResponse []CarclassGetResponseElement

type CarclassGetResponseElement struct {
	CarClassID    int64         `json:"car_class_id"`
	CarsInClass   []CarsInClass `json:"cars_in_class"`
	CustID        int64         `json:"cust_id"`
	Name          string        `json:"name"`
	RainEnabled   bool          `json:"rain_enabled"`
	RelativeSpeed int64         `json:"relative_speed"`
	ShortName     string        `json:"short_name"`
}

type CarsInClass struct {
	CarDirpath  string `json:"car_dirpath"`
	CarID       int64  `json:"car_id"`
	RainEnabled bool   `json:"rain_enabled"`
	Retired     bool   `json:"retired"`
}
