import os
import logging
import subprocess

from tqdm.contrib.concurrent import thread_map

from format import to_camel_case
from constants import PARAM_TYPES
from api_client import APIClient
from typing import Literal
from data import OVERRIDES


class EndpointParameter:
    name: str
    type: str
    required: bool
    notes: list[str]

    sample_value = None

    def __init__(
        self,
        category_name: str,
        endpoint_name: str,
        param_name: str,
        param_data: dict,
    ):
        self.name = param_name
        self.type = param_data["type"]
        self.required = param_data.get("required", False)
        self.notes = _parse_iracing_notes(param_data.get("notes"))

        # Check that the type is known
        if self.type not in PARAM_TYPES.keys():
            raise Exception(f"Unknown type: {self.type}")

        # If available, set the sample value
        self.sample_value = (
            OVERRIDES.get(category_name, {})
            .get(endpoint_name, {})
            .get("params", {})
            .get(param_name, None)
        )


class Endpoint:
    api_client: APIClient

    category: str
    name: str

    link: str
    notes: list[str]
    parameters: list[EndpointParameter]

    format: Literal["json", "csv"] = "json"
    s3_cache: bool = True

    sample_response: str | None = None
    response_struct_name: str | None = None
    response_struct: str | None = None

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
            EndpointParameter(category_name, endpoint_name, param_name, param_data)
            for param_name, param_data in endpoint_data.get("parameters", {}).items()
        ]

        # TODO: allow override self.format = "json"
        # TODO: allow override self.s3_cache = True

    def fetch_sample_response(self, cached: bool = True) -> bool:
        # TODO: handle driver_stats_by_category
        if self.category == "driver_stats_by_category":
            logging.warning(
                f"Skipping sample response fetch for {self.category}__{self.name} due to complexity."
            )
            return False

        # Skip if some parameters are not set
        for parameter in self.parameters:
            if parameter.required and not parameter.sample_value:
                logging.error(
                    f"Skipping sample response fetch for {self.category}__{self.name} due to parameters."
                )
                return False

        file_path = f"output/responses/{self.category}__{self.name}.{self.format}"
        if cached and os.path.exists(file_path):
            with open(file_path, "r") as f:
                self.sample_response = f.read()
            return True

        try:
            self.sample_response = self.api_client.call_endpoint(
                self.link,
                params={
                    parameter.name: parameter.sample_value
                    for parameter in self.parameters
                },
            )
        except Exception as e:
            logging.error(f"Error fetching {self.category}__{self.name}: {e}")
            return False

        # Save the response to a file
        os.makedirs("output/responses", exist_ok=True)
        with open(file_path, "w") as f:
            f.write(self.sample_response)

        return True

    def generate_go_types(self):
        """
        Run the quicktype command to generate Go types.
        """

        if not self.sample_response:
            logging.warning(
                f"Skipped: {self.category}__{self.name} (no sample response)"
            )
            return

        # Generate the struct name from the file name
        self.response_struct_name = to_camel_case(
            f"{self.category}__{self.name}__Response"
        )

        # Quicktype command
        command = [
            "npx",
            "quicktype",
            "--src-lang",
            "json",
            "--lang",
            "go",
            "--just-types",
            "-t",
            self.response_struct_name,
        ]

        try:
            # Run the command and capture the output. Pass the sample response as stdin.
            result = subprocess.run(
                command,
                check=True,  # Raise an error if the command fails
                capture_output=True,
                text=True,
                input=self.sample_response,
            )
            self.response_struct = result.stdout

        except subprocess.CalledProcessError as e:
            logging.error(f"FAILED: Quicktype for {self.category}__{self.name} failed.")
            logging.error(f"Error: {e.stderr.strip()}")

        except FileNotFoundError:
            logging.error(
                "ERROR: Make sure 'npx' and 'quicktype' are installed and available in the PATH."
            )


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


def fetch_sample_responses(
    endpoints: list[Endpoint], skip_cached: bool = True, workers: int = 5
):
    logging.info("Fetching sample responses...")

    thread_map(
        lambda endpoint: endpoint.fetch_sample_response(cached=skip_cached),
        endpoints,
        max_workers=workers,
    )

    logging.info("Sample responses fetched.")


def generate_go_types(endpoints: list[Endpoint], workers: int = 20):
    logging.info("Generating Go types...")

    thread_map(
        lambda endpoint: endpoint.generate_go_types(),
        endpoints,
        max_workers=workers,
    )

    logging.info("Go types generated.")
