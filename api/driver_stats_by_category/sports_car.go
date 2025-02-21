package driver_stats_by_category

import (
	"io"
)

func (api *DriverStatsByCategoryApi) GetDriverStatsByCategorySportsCar() (io.ReadCloser, error) {
	url := "https://members-ng.iracing.com/data/driver_stats_by_category/sports_car"
	body, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	return body, nil
}
