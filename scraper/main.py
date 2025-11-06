import os
import logging

from dotenv import load_dotenv

from api_client import APIClient
import structs_parser
import endpoints_parsing
import categories_writing


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

    # Get the list of available endpoints, parse them and save sample responses
    endpoints = endpoints_parsing.run(api_client)
    logging.info(f"Found {len(endpoints)} endpoints")

    # Generate Go types from JSON responses
    structs_parser.run()

    # Generate category APIs
    categories_writing.run(endpoints)

    # # Generate endpoint calls
    # # TODO: specify map fields
    # # TODO: handle csv responses
    # # TODO: loader
    # print("Generating endpoint calls...")


if __name__ == "__main__":
    main()
