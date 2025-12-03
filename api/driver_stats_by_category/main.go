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

func (api *DriverStatsByCategoryApi) Oval() ([]oval.DriverStatsByCategoryOvalResponse, error) {
	resp, err := client.GetCsv[oval.DriverStatsByCategoryOvalResponse](api.Client, "/data/driver_stats_by_category/oval", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *DriverStatsByCategoryApi) SportsCar() ([]sports_car.DriverStatsByCategorySportsCarResponse, error) {
	resp, err := client.GetCsv[sports_car.DriverStatsByCategorySportsCarResponse](api.Client, "/data/driver_stats_by_category/sports_car", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *DriverStatsByCategoryApi) FormulaCar() ([]formula_car.DriverStatsByCategoryFormulaCarResponse, error) {
	resp, err := client.GetCsv[formula_car.DriverStatsByCategoryFormulaCarResponse](api.Client, "/data/driver_stats_by_category/formula_car", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *DriverStatsByCategoryApi) Road() ([]road.DriverStatsByCategoryRoadResponse, error) {
	resp, err := client.GetCsv[road.DriverStatsByCategoryRoadResponse](api.Client, "/data/driver_stats_by_category/road", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *DriverStatsByCategoryApi) DirtOval() ([]dirt_oval.DriverStatsByCategoryDirtOvalResponse, error) {
	resp, err := client.GetCsv[dirt_oval.DriverStatsByCategoryDirtOvalResponse](api.Client, "/data/driver_stats_by_category/dirt_oval", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *DriverStatsByCategoryApi) DirtRoad() ([]dirt_road.DriverStatsByCategoryDirtRoadResponse, error) {
	resp, err := client.GetCsv[dirt_road.DriverStatsByCategoryDirtRoadResponse](api.Client, "/data/driver_stats_by_category/dirt_road", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
