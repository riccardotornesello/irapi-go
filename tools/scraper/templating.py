import os
import logging

import jinja2

from format import to_camel_case
from endpoints_parsing import Endpoint


jinja2_environment = jinja2.Environment(loader=jinja2.FileSystemLoader("templates/"))

endpoint_template = jinja2_environment.get_template("endpoint.go")
category_template = jinja2_environment.get_template("category.go")


def write_category_apis(endpoints: list[Endpoint]):
    logging.info("Generating category APIs...")

    # Get unique categories
    categories = set([endpoint.category for endpoint in endpoints])

    # Generate the API code for each category
    for category in categories:
        api_code = category_template.render(
            package_name=category, api_name=to_camel_case(category) + "Api"
        )
        os.makedirs(f"../api/{category}", exist_ok=True)
        with open(f"../api/{category}/main.go", "w") as f:
            f.write(api_code)

    logging.info("Category APIs generated successfully.")


def write_endpoint_api(endpoint: Endpoint):
    if not endpoint.response_struct:
        logging.warning(
            f"Skipped: {endpoint.category}__{endpoint.name} (no response struct)"
        )
        return

    package_name = endpoint.category
    api_name = to_camel_case(endpoint.category) + "Api"
    endpoint_url = endpoint.link.replace("https://members-ng.iracing.com", "")
    method_name = to_camel_case(endpoint.name)

    with open(f"../api/{endpoint.category}/{endpoint.name}.go", "w") as output_file:
        output_file.write(
            endpoint_template.render(
                package_name=package_name,
                api_name=api_name,
                method_name=method_name,
                endpoint_url=endpoint_url,
                response_struct_code=endpoint.response_struct,
                response_struct_name=endpoint.response_struct_name,
            )
        )


def write_endpoint_apis(endpoints: list[Endpoint]):
    logging.info("Generating endpoint APIs...")

    for endpoint in endpoints:
        write_endpoint_api(endpoint)

    logging.info("Endpoint APIs generated successfully.")
