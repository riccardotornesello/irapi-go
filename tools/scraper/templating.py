"""
Module for generating Go API client code from Jinja2 templates.

This module uses Jinja2 templates to generate:
- Category-level API files (main.go for each category)
- Endpoint-level struct files (structs.go for each endpoint)
- API method implementations

Templates are loaded from the templates/ directory.
"""

import os
import logging

import jinja2

from format import to_camel_case
from endpoints_parsing import Endpoint


# Initialize Jinja2 environment with template directory
jinja2_environment = jinja2.Environment(loader=jinja2.FileSystemLoader("templates/"))

# Load templates
endpoint_call_template = jinja2_environment.get_template("endpoint_call.j2")
endpoint_call_csv_template = jinja2_environment.get_template("endpoint_call_csv.j2")
endpoint_structs_template = jinja2_environment.get_template("endpoint_structs.j2")
category_template = jinja2_environment.get_template("category.j2")
category_tests_template = jinja2_environment.get_template("category_tests.j2")
endpoint_test_template = jinja2_environment.get_template("endpoint_test.j2")


def write_category_apis(endpoints: list[Endpoint]) -> None:
    """
    Generate category-level API files.

    Creates main.go files for each category, containing an API struct with
    methods for all endpoints in that category.

    Args:
        endpoints (list[Endpoint]): List of all endpoints to process.
    """
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
        endpoint_tests = []
        required_imports = {"github.com/riccardotornesello/irapi-go/client"}
        test_packages = []
        for endpoint in endpoints:
            if not endpoint.response_struct:
                continue

            endpoint_template = (
                endpoint_call_csv_template
                if endpoint.format == "csv"
                else endpoint_call_template
            )
            endpoint_calls.append(
                endpoint_template.render(
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

            endpoint_tests.append(
                endpoint_test_template.render(
                    api_name=api_name,
                    method_name=to_camel_case(endpoint.name),
                    endpoint_name=endpoint.name,
                    parameters_struct_name=endpoint.parameters_struct_name,
                )
            )
            if endpoint.parameters_struct_name:
                test_packages.append(
                    f"github.com/riccardotornesello/irapi-go/api/{category}/{endpoint.name}"
                )

        api_code = category_template.render(
            package_name=category,
            api_name=api_name,
            endpoint_calls=endpoint_calls,
            imports=required_imports,
        )
        tests_code = category_tests_template.render(
            package_name=category,
            endpoint_tests=endpoint_tests,
            test_packages=test_packages,
        )

        os.makedirs(f"../../api/{category}", exist_ok=True)
        with open(f"../../api/{category}/main.go", "w") as f:
            f.write(api_code)
        # TODO: uncomment and write tests
        # if len(endpoint_tests) > 0:
        #     with open(f"../../api/{category}/main_test.go", "w") as f:
        #         f.write(tests_code)

    logging.info("Category APIs generated successfully.")


def write_endpoint_api(endpoint: Endpoint) -> None:
    """
    Generate structs.go file for a single endpoint.

    Creates a file containing the parameter and response struct definitions
    for the endpoint.

    Args:
        endpoint (Endpoint): The endpoint to generate structs for.
    """
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


def write_endpoint_apis(endpoints: list[Endpoint]) -> None:
    """
    Generate structs.go files for all endpoints.

    Processes all endpoints and creates their respective struct definition files.

    Args:
        endpoints (list[Endpoint]): List of all endpoints to generate structs for.
    """
    logging.info("Generating endpoint APIs...")

    for endpoint in endpoints:
        write_endpoint_api(endpoint)

    logging.info("Endpoint APIs generated successfully.")
