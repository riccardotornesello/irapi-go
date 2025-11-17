import os
import logging

from dotenv import load_dotenv

from api_client import APIClient
from endpoints_documentation import generate_endpoints_list
from endpoints_parsing import (
    fetch_sample_responses,
    generate_go_types,
    fetch_sample_chunks,
)
from templating import write_category_apis, write_endpoint_apis
from format import format_go_code


logging.basicConfig(
    level=logging.INFO, format="%(asctime)s - %(levelname)s - %(message)s"
)


def main():
    # Load environment variables
    load_dotenv()

    # Authenticate and init api client
    api_client = APIClient(
        email=os.getenv("IRACING_EMAIL"),
        password=os.getenv("IRACING_PASSWORD"),
    )

    # Get the list of available endpoints and parse them
    endpoints = generate_endpoints_list(api_client)
    logging.info(f"Found {len(endpoints)} endpoints")

    # Fetch sample responses for each endpoint
    fetch_sample_responses(endpoints, skip_cached=True, workers=5)

    # Fetch sample chunks for each endpoint
    fetch_sample_chunks(endpoints, skip_cached=True, workers=10)

    # Generate Go types from JSON responses
    generate_go_types(endpoints, workers=20)

    # Generate the APIs
    write_category_apis(endpoints)
    write_endpoint_apis(endpoints)

    # Format the code
    format_go_code()


if __name__ == "__main__":
    main()
