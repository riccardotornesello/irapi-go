package {{package_name}}

import (
	{% for imp in imports %}"{{imp}}"
	{% endfor %}
)

{% if parameters_struct_code %}
{{parameters_struct_code}}
{% endif %}

{{response_struct_code}}
