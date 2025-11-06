package lookup

import (
	"encoding/json"
)

type LookupGetResponse []interface{}

func (api *LookupApi) Get() (*LookupGetResponse, error) {
	return api.GetJson[LookupGetResponse]("/data/lookup/get")
}
