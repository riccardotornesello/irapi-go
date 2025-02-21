package driver_stats_by_category

import (
	"io"
)

func (api *DriverStatsByCategoryApi) GetDriverStatsByCategoryDirtOval() (io.ReadCloser, error) {
	url := "https://members-ng.iracing.com/data/driver_stats_by_category/dirt_oval"
	body, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	return body, nil
}
