import os

from constants import PARAM_TYPES
from api_client import APIClient
from typing import Literal


class EndpointParameter:
    name: str
    type: str
    required: bool
    notes: list[str]

    def __init__(self, param_name: str, param_data: dict):
        self.name = param_name
        self.type = param_data["type"]
        self.required = param_data.get("required", False)
        self.notes = _parse_iracing_notes(param_data.get("notes"))

        # Check that the type is known
        if self.type not in PARAM_TYPES.keys():
            raise Exception(f"Unknown type: {self.type}")


class Endpoint:
    api_client: APIClient

    category: str
    name: str

    link: str
    notes: list[str]
    parameters: list[EndpointParameter]

    format: Literal["json", "csv"] = "json"
    s3_cache: bool = True

    def __init__(
        self,
        api_client: APIClient,
        category_name: str,
        endpoint_name: str,
        endpoint_data: dict,
    ):
        self.api_client = api_client
        self.category = category_name
        self.name = endpoint_name
        self.link = endpoint_data["link"]
        self.notes = _parse_iracing_notes(endpoint_data.get("notes"))
        self.parameters = [
            EndpointParameter(param_name, param_data)
            for param_name, param_data in endpoint_data.get("parameters", {}).items()
        ]

        # TODO: allow override self.format = "json"
        # TODO: allow override self.s3_cache = True

    def save_response(self, cached: bool = True) -> bool:
        file_path = f"output/responses/{self.category}__{self.name}.{self.format}"
        if cached and os.path.exists(file_path):
            return True

        try:
            response = self.api_client.call_endpoint(self.link)
        except Exception as e:
            print(f"Error fetching {self.category}__{self.name}: {e}")
            return False

        # Save the response to a file
        os.makedirs("output/responses", exist_ok=True)
        with open(file_path, "w") as f:
            f.write(response)

        return True


def _parse_iracing_notes(notes) -> list[str]:
    """
    Gets the notes from the endpoint or the parameter's data and returns a list of notes.
    The list is empty if there are no notes.
    """

    if not notes:
        return []

    if not isinstance(notes, list):
        notes = [notes]

    return notes


def generate_endpoints(
    api_client: APIClient, iracing_documentation: dict
) -> list[Endpoint]:
    endpoints = []

    for category_name, category_data in iracing_documentation.items():
        for endpoint_name, endpoint_data in category_data.items():
            endpoints.append(
                Endpoint(api_client, category_name, endpoint_name, endpoint_data)
            )

    return endpoints
