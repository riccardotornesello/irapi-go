import json
import os
import shutil
import subprocess

import jinja2
from dotenv import load_dotenv

from session import authentiate_iracing
from endpoints import generate_endpoints, get_api_definition
from responses import get_responses
from schema import gen_schemas
from structs import convert_schema_to_struct, generate_parameters_struct
from report import generate_report


def to_camel_case(snake_str):
    return "".join(x.title() for x in snake_str.split("_"))


def main():
    load_dotenv()

    # Authenticate
    session = authentiate_iracing(
        email=os.getenv("IRACING_EMAIL"),
        password=os.getenv("IRACING_PASSWORD"),
    )

    # Clean the output folder
    shutil.rmtree("output/api", ignore_errors=True)
    shutil.rmtree("output/schemas", ignore_errors=True)

    os.makedirs("output", exist_ok=True)
    os.makedirs("output/schemas", exist_ok=True)
    os.makedirs("output/responses", exist_ok=True)
    os.makedirs("output/api", exist_ok=True)

    # Get the API definition
    api_definition = get_api_definition(session, cached=True)

    # Format the API definition
    endpoints = generate_endpoints(api_definition)

    # Get some sample responses
    # TODO: get responses that require parameters
    get_responses(session, endpoints, cached=True)

    # Generate the schemas
    # TODO: create a schema for the csv responses
    gen_schemas()

    # Prepare the templates
    jinja2_environment = jinja2.Environment(
        loader=jinja2.FileSystemLoader("templates/")
    )
    client_template = jinja2_environment.get_template("client.j2")
    endpoint_template = jinja2_environment.get_template("endpoint.j2")
    endpoint_csv_template = jinja2_environment.get_template("endpoint_csv.j2")
    test_template = jinja2_environment.get_template("test.j2")
    api_template = jinja2_environment.get_template("api.j2")

    # Store the generated categories list to create the api.go file
    categories = []

    for category, category_endpoints in endpoints.items():
        os.makedirs(f"output/api/{category}")

        camel_category = to_camel_case(category)
        client_struct_name = camel_category + "Api"

        # Create the category client file
        with open(f"output/api/{category}/main.go", "w") as f:
            f.write(
                client_template.render(
                    package_name=category,
                    client_struct_name=client_struct_name,
                )
            )

        categories.append(
            {
                "camel_category": camel_category,
                "package_name": category,
                "client_struct_name": client_struct_name,
            }
        )

        for endpoint, endpoint_data in category_endpoints.items():
            endpoint_file = f"output/api/{category}/{endpoint}.go"
            test_file = f"output/api/{category}/{endpoint}_test.go"

            endpoint_url = endpoint_data["link"].replace(
                "https://members-ng.iracing.com", ""
            )

            camel_endpoint = to_camel_case(endpoint)
            method_name = "Get" + camel_category + camel_endpoint
            if method_name[-3:] == "Get":
                method_name = method_name[:-3]

            ##################### JSON
            if endpoint_data["format"] == "json":
                response_schema_file = f"output/schemas/{category}__{endpoint}.json"
                chunks_schema_file = (
                    f"output/schemas/{category}__{endpoint}__chunks.json"
                )

                if not os.path.exists(response_schema_file):
                    print(f"Schema file not found: {response_schema_file}")
                    continue

                chunks_struct_name = None
                chunks_struct = None
                if os.path.exists(chunks_schema_file):
                    with open(chunks_schema_file) as f:
                        chunks_struct_name = (
                            camel_category + camel_endpoint + "ResponseChunks"
                        )
                        chunks_struct = convert_schema_to_struct(json.load(f))

                with open(response_schema_file) as f:
                    struct = convert_schema_to_struct(
                        json.load(f), chunks_struct_name=chunks_struct_name
                    )

                params_struct, requires_optional = (
                    generate_parameters_struct(endpoint_data["parameters"])
                    if len(endpoint_data["parameters"]) > 0
                    else (None, False)
                )

                template_data = {
                    "package_name": category,
                    "struct_name": camel_category + camel_endpoint + "Response",
                    "struct": struct,
                    "chunks_struct_name": chunks_struct_name,
                    "chunks_struct": chunks_struct,
                    "client_struct_name": client_struct_name,
                    "method_name": method_name,
                    "endpoint": endpoint_url,
                    "params_struct": params_struct,
                    "params_struct_name": camel_category + camel_endpoint + "Params",
                    "requires_optional": requires_optional,
                    "notes": endpoint_data["notes"],
                }

                with open(endpoint_file, "w") as f:
                    f.write(endpoint_template.render(**template_data))

                with open(test_file, "w") as f:
                    f.write(test_template.render(**template_data))

            ##################### CSV
            elif endpoint_data["format"] == "csv":
                template_data = {
                    "package_name": category,
                    "struct": struct,
                    "client_struct_name": client_struct_name,
                    "method_name": method_name,
                    "endpoint": endpoint_url,
                    "notes": endpoint_data["notes"],
                }

                with open(endpoint_file, "w") as f:
                    f.write(endpoint_csv_template.render(**template_data))

                with open(test_file, "w") as f:
                    f.write(test_template.render(**template_data))

            ##################### Unknown
            else:
                print(f"Unknown format: {endpoint_data['format']}")

    with open("output/api.go", "w") as f:
        f.write(api_template.render(categories=categories))

    # Format the output
    subprocess.call(
        ["go", "fmt", "./output/api/..."],
        stdout=open(os.devnull, "w"),
        stderr=subprocess.STDOUT,
    )

    generate_report()


main()
