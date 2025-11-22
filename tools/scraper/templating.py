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
endpoint_structs_template = jinja2_environment.get_template("endpoint_structs.j2")
endpoint_test_template = jinja2_environment.get_template("endpoint_test.j2")
category_template = jinja2_environment.get_template("category.j2")


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


def write_endpoint_test(endpoint: Endpoint) -> None:
    """
    Generate test file for a single endpoint.
    
    Creates a test file containing unit tests for the endpoint's structs,
    including marshaling and unmarshaling tests.
    
    Args:
        endpoint (Endpoint): The endpoint to generate tests for.
    """
    if not endpoint.response_struct:
        logging.warning(
            f"Skipped test: {endpoint.category}__{endpoint.name} (no response struct)"
        )
        return

    os.makedirs(f"../../api/{endpoint.category}/{endpoint.name}", exist_ok=True)
    with open(
        f"../../api/{endpoint.category}/{endpoint.name}/structs_test.go", "w"
    ) as output_file:
        output_file.write(
            endpoint_test_template.render(
                package_name=endpoint.name,
                parameters_struct_name=endpoint.parameters_struct_name,
                response_struct_name=endpoint.response_struct_name,
                endpoint_name=endpoint.name,
                category=endpoint.category,
                has_parameters=endpoint.parameters_struct_name is not None,
            )
        )


def write_endpoint_tests(endpoints: list[Endpoint]) -> None:
    """
    Generate test files for all endpoints.
    
    Processes all endpoints and creates their respective test files.
    
    Args:
        endpoints (list[Endpoint]): List of all endpoints to generate tests for.
    """
    logging.info("Generating endpoint tests...")

    for endpoint in endpoints:
        write_endpoint_test(endpoint)

    logging.info("Endpoint tests generated successfully.")
