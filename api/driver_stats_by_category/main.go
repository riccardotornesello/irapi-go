package driver_stats_by_category

import (
	"github.com/riccardotornesello/irapi-go/api/driver_stats_by_category/dirt_oval"
	"github.com/riccardotornesello/irapi-go/api/driver_stats_by_category/dirt_road"
	"github.com/riccardotornesello/irapi-go/api/driver_stats_by_category/formula_car"
	"github.com/riccardotornesello/irapi-go/api/driver_stats_by_category/oval"
	"github.com/riccardotornesello/irapi-go/api/driver_stats_by_category/road"
	"github.com/riccardotornesello/irapi-go/api/driver_stats_by_category/sports_car"
	"github.com/riccardotornesello/irapi-go/client"
)

type DriverStatsByCategoryApi struct {
	Client *client.ApiClient
}

func (api *DriverStatsByCategoryApi) Oval() (*oval.None, error) {
	return client.GetJson[oval.None](api.Client, "/data/driver_stats_by_category/oval", nil)
}

func (api *DriverStatsByCategoryApi) SportsCar() (*sports_car.None, error) {
	return client.GetJson[sports_car.None](api.Client, "/data/driver_stats_by_category/sports_car", nil)
}

func (api *DriverStatsByCategoryApi) FormulaCar() (*formula_car.None, error) {
	return client.GetJson[formula_car.None](api.Client, "/data/driver_stats_by_category/formula_car", nil)
}

func (api *DriverStatsByCategoryApi) Road() (*road.None, error) {
	return client.GetJson[road.None](api.Client, "/data/driver_stats_by_category/road", nil)
}

func (api *DriverStatsByCategoryApi) DirtOval() (*dirt_oval.None, error) {
	return client.GetJson[dirt_oval.None](api.Client, "/data/driver_stats_by_category/dirt_oval", nil)
}

func (api *DriverStatsByCategoryApi) DirtRoad() (*dirt_road.None, error) {
	return client.GetJson[dirt_road.None](api.Client, "/data/driver_stats_by_category/dirt_road", nil)
}
