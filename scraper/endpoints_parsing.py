import os
import logging

import requests
from dotenv import load_dotenv
from tqdm.contrib.concurrent import thread_map

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
        # TODO: handle driver_stats_by_category
        if self.category == "driver_stats_by_category":
            return False

        file_path = f"output/responses/{self.category}__{self.name}.{self.format}"
        if cached and os.path.exists(file_path):
            return True

        try:
            response = self.api_client.call_endpoint(self.link)
        except Exception as e:
            logging.error(f"Error fetching {self.category}__{self.name}: {e}")
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


def get_available_endpoints_documentation(session: requests.Session) -> dict:
    res = session.get("https://members-ng.iracing.com/data/doc")
    if res.status_code != 200:
        raise Exception(f"Failed to fetch API definition: {res.text}")

    return res.json()


def generate_endpoints_list(api_client: APIClient) -> list[Endpoint]:
    endpoints = []

    iracing_documentation = get_available_endpoints_documentation(api_client.api_client)

    for category_name, category_data in iracing_documentation.items():
        for endpoint_name, endpoint_data in category_data.items():
            endpoints.append(
                Endpoint(api_client, category_name, endpoint_name, endpoint_data)
            )

    return endpoints


def save_sample_responses(
    endpoints: list[Endpoint], skip_cached: bool = True, workers: int = 5
):
    logging.info("Fetching sample responses...")

    thread_map(
        lambda endpoint: endpoint.save_response(cached=skip_cached),
        endpoints,
        max_workers=workers,
    )

    logging.info("Sample responses fetched.")


def run(api_client: APIClient, workers: int = 5):
    endpoints = generate_endpoints_list(api_client)
    save_sample_responses(endpoints, skip_cached=True, workers=workers)
    return endpoints


if __name__ == "__main__":
    logging.basicConfig(
        level=logging.INFO, format="%(asctime)s - %(levelname)s - %(message)s"
    )

    load_dotenv()

    api_client = APIClient(
        email=os.getenv("IRACING_EMAIL"),
        password=os.getenv("IRACING_PASSWORD"),
    )

    run(api_client)
