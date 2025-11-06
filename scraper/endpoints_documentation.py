import requests
import logging
import os

from dotenv import load_dotenv

from api_client import APIClient
from endpoints_parsing import Endpoint


def _get_available_endpoints_documentation(session: requests.Session) -> dict:
    res = session.get("https://members-ng.iracing.com/data/doc")
    if res.status_code != 200:
        raise Exception(f"Failed to fetch API definition: {res.text}")

    return res.json()


def generate_endpoints_list(api_client: APIClient) -> list[Endpoint]:
    endpoints = []

    logging.info("Fetching available endpoints documentation...")

    iracing_documentation = _get_available_endpoints_documentation(api_client.api_client)

    for category_name, category_data in iracing_documentation.items():
        for endpoint_name, endpoint_data in category_data.items():
            endpoints.append(
                Endpoint(api_client, category_name, endpoint_name, endpoint_data)
            )

    logging.info("Endpoints documentation fetched and parsed successfully.")

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

    endpoints = generate_endpoints_list(api_client)

    logging.info(f"Total endpoints parsed: {len(endpoints)}")
