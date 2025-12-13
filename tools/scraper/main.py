"""
Main module for the iRacing API scraper tool.

This tool automates the generation of Go client code for the iRacing API by:
1. Fetching API endpoint documentation from iRacing
2. Retrieving sample responses from each endpoint
3. Generating Go type definitions from JSON responses
4. Creating Go API client methods using templates
5. Formatting the generated code

The tool requires iRacing credentials to be set in environment variables.
"""

import os
import logging

from dotenv import load_dotenv

from api_client import APIClient
from endpoints_documentation import generate_endpoints_list
from endpoints_parsing import fetch_sample_responses, generate_go_types
from templating import write_category_apis, write_endpoint_apis
from format import format_go_code


logging.basicConfig(
    level=logging.INFO, format="%(asctime)s - %(levelname)s - %(message)s"
)


def main() -> None:
    """
    Main entry point for the scraper tool.
    
    Orchestrates the complete workflow of generating Go API client code:
    1. Loads environment variables (IRACING_EMAIL, IRACING_PASSWORD)
    2. Authenticates with the iRacing API
    3. Fetches and parses available endpoints
    4. Retrieves sample responses for endpoints
    5. Generates Go type definitions from responses
    6. Generates API client code from templates
    7. Formats the generated Go code
    """
    # Load environment variables
    load_dotenv()

    # Authenticate and init api client
    api_client = APIClient(
        email=os.getenv("IRACING_EMAIL"),
        password=os.getenv("IRACING_PASSWORD"),
        client_id=os.getenv("IRACING_CLIENT_ID"),
        client_secret=os.getenv("IRACING_CLIENT_SECRET"),
    )

    # Get the list of available endpoints and parse them
    endpoints = generate_endpoints_list(api_client)
    logging.info(f"Found {len(endpoints)} endpoints")

    # Fetch sample responses for each endpoint
    fetch_sample_responses(endpoints, skip_cached=True, workers=5)

    # Generate Go types from JSON responses
    generate_go_types(endpoints, workers=20)

    # Generate the APIs
    write_category_apis(endpoints)
    write_endpoint_apis(endpoints)

    # Format the code
    format_go_code()


if __name__ == "__main__":
    main()
