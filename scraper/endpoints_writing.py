import os

from tqdm.contrib.concurrent import thread_map

from endpoints_parsing import Endpoint
from format import to_camel_case


def run(endpoints: list[Endpoint]):
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

        with open(f"../api/{endpoint.category}/{endpoint.name}.go", "w") as output_file:
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
