import os
import logging

import jinja2

from format import to_camel_case
from endpoints_parsing import Endpoint


jinja2_environment = jinja2.Environment(loader=jinja2.FileSystemLoader("templates/"))

endpoint_call_template = jinja2_environment.get_template("endpoint_call.j2")
endpoint_structs_template = jinja2_environment.get_template("endpoint_structs.j2")
category_template = jinja2_environment.get_template("category.j2")


def write_category_apis(endpoints: list[Endpoint]):
    logging.info("Generating category APIs...")

    # Group endpoints by category
    grouped_endpoints = {}
    for endpoint in endpoints:
        if endpoint.category not in grouped_endpoints:
            grouped_endpoints[endpoint.category] = []
        grouped_endpoints[endpoint.category].append(endpoint)

    # Generate the API code for each category
    for category, endpoints in grouped_endpoints.items():
        api_name = to_camel_case(category) + "Api"

        endpoint_calls = []
        required_imports = {"github.com/riccardotornesello/irapi-go/client"}
        for endpoint in endpoints:
            if not endpoint.response_struct:
                continue

            endpoint_calls.append(
                endpoint_call_template.render(
                    api_name=api_name,
                    method_name=to_camel_case(endpoint.name),
                    endpoint_name=endpoint.name,
                    parameters_struct_name=endpoint.parameters_struct_name,
                    response_struct_name=endpoint.response_struct_name,
                    endpoint_url=endpoint.link.replace(
                        "https://members-ng.iracing.com", ""
                    ),
                    response_chunk_type=endpoint.chunk_struct_name,
                )
            )
            required_imports.add(
                f"github.com/riccardotornesello/irapi-go/api/{category}/{endpoint.name}"
            )

        api_code = category_template.render(
            package_name=category,
            api_name=api_name,
            endpoint_calls=endpoint_calls,
            imports=required_imports,
        )
        os.makedirs(f"../../api/{category}", exist_ok=True)
        with open(f"../../api/{category}/main.go", "w") as f:
            f.write(api_code)

    logging.info("Category APIs generated successfully.")


def write_endpoint_api(endpoint: Endpoint):
    if not endpoint.response_struct:
        logging.warning(
            f"Skipped: {endpoint.category}__{endpoint.name} (no response struct)"
        )
        return

    os.makedirs(f"../../api/{endpoint.category}/{endpoint.name}", exist_ok=True)
    with open(
        f"../../api/{endpoint.category}/{endpoint.name}/structs.go", "w"
    ) as output_file:
        output_file.write(
            endpoint_structs_template.render(
                package_name=endpoint.name,
                response_struct_code=endpoint.response_struct,
                parameters_struct_code=endpoint.parameters_struct,
                imports=endpoint.required_imports,
            )
        )


def write_endpoint_apis(endpoints: list[Endpoint]):
    logging.info("Generating endpoint APIs...")

    for endpoint in endpoints:
        write_endpoint_api(endpoint)

    logging.info("Endpoint APIs generated successfully.")
