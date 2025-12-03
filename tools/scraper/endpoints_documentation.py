"""
Module for fetching and parsing iRacing API endpoint documentation.

This module retrieves the list of available API endpoints from iRacing
and creates Endpoint objects that can be used to generate Go client code.
"""

import requests
import logging
import os
import json

from dotenv import load_dotenv

from api_client import APIClient
from endpoints_parsing import Endpoint


def _get_available_endpoints_documentation(api_client: APIClient) -> dict:
    """
    Fetch the API documentation from iRacing.

    Args:
        api_client (APIClient): Authenticated API client to use for the request.

    Returns:
        dict: Dictionary containing all available endpoints organized by category.

    Raises:
        Exception: If the API documentation cannot be fetched.
    """
    res = api_client.call("https://members-ng.iracing.com/data/doc")
    if res.status_code != 200:
        raise Exception(f"Failed to fetch API definition: {res.text}")

    return res.json()


def generate_endpoints_list(api_client: APIClient) -> list[Endpoint]:
    """
    Generate a list of Endpoint objects from the iRacing API documentation.

    Fetches the complete API documentation from iRacing, saves it locally,
    and creates Endpoint objects for each discovered endpoint.

    Args:
        api_client (APIClient): Authenticated API client to use for fetching.

    Returns:
        list[Endpoint]: List of parsed Endpoint objects ready for code generation.
    """
    endpoints = []

    logging.info("Fetching available endpoints documentation...")

    # Get the documentation of available endpoints and save it locally
    iracing_documentation = _get_available_endpoints_documentation(api_client)
    with open("../../doc.json", "w") as f:
        json.dump(iracing_documentation, f, indent=4)

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
        client_id=os.getenv("IRACING_CLIENT_ID"),
        client_secret=os.getenv("IRACING_CLIENT_SECRET"),
    )

    endpoints = generate_endpoints_list(api_client)

    logging.info(f"Total endpoints parsed: {len(endpoints)}")
