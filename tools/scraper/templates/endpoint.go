package {{package_name}}

import (
	"encoding/json"
)

{{response_struct_code}}

func (api *{{api_name}}) {{method_name}}() (*{{response_struct_name}}, error) {
	return api.GetJson[{{response_struct_name}}]("{{endpoint_url}}")
}
