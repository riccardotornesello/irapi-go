import os

import jinja2
from dotenv import load_dotenv
from tqdm import tqdm

from api_client import APIClient
from endpoints import generate_endpoints
from parser import JsonToGo
from format import to_camel_case
from categories import generate_categories_api

jinja2_environment = jinja2.Environment(loader=jinja2.FileSystemLoader("templates/"))


def main():
    # Load environment variables
    load_dotenv()

    # Prepare the templates
    endpoint_template = jinja2_environment.get_template("endpoint.go")

    # Authenticate and init api client
    api_client = APIClient(
        email=os.getenv("IRACING_EMAIL"),
        password=os.getenv("IRACING_PASSWORD"),
    )

    # Get the list of available endpoints and parse them
    available_endpoints = api_client.get_available_endpoints()
    endpoints = generate_endpoints(api_client, available_endpoints)
    print(f"Found {len(endpoints)} endpoints")

    # Save sample responses
    print("Fetching sample responses...")
    for endpoint in tqdm(endpoints):
        # TODO: handle this category
        if endpoint.category == "driver_stats_by_category":
            continue

        endpoint.save_response()

    # Generate category APIs
    print("Generating category APIs...")
    generate_categories_api(endpoints)

    # Generate endpoint calls
    # TODO: specify map fields
    # TODO: handle csv responses
    # TODO: loader
    print("Generating endpoint calls...")
    for endpoint in endpoints:
        response_file_name = f"{endpoint.category}__{endpoint.name}.json"
        response_file_path = os.path.join("output/responses", response_file_name)
        if not os.path.exists(response_file_path):
            print(f"Skipping {response_file_name}, no response file found")
            continue

        response_struct_name = (
            to_camel_case(endpoint.category) + to_camel_case(endpoint.name) + "Response"
        )

        with open(response_file_path, "r") as input_file:
            json_string = input_file.read()
            converter = JsonToGo(
                response_struct_name,
                flatten=False,
                example=True,
                all_omitempty=False,
            )
            struct_result = converter.convert(json_string)

        if struct_result.get("error"):
            print(f"Error processing {response_file_name}: {struct_result['error']}")
            continue

        response_struct_code = struct_result["go"]
        package_name = endpoint.category
        api_name = to_camel_case(endpoint.category) + "Api"
        endpoint_url = endpoint.link  # TODO: remove base URL if present
        method_name = to_camel_case(endpoint.name)

        with open(
            f"../api/{endpoint.category}/{endpoint.name}.go", "w"
        ) as output_file:
            output_file.write(
                endpoint_template.render(
                    package_name=package_name,
                    api_name=api_name,
                    method_name=method_name,
                    endpoint_url=endpoint_url,
                    response_struct_code=response_struct_code,
                    response_struct_name=response_struct_name,
                )
            )


if __name__ == "__main__":
    main()
