import os
import logging
import shutil
import requests
import json
import re

from tqdm.contrib.concurrent import thread_map

from format import to_camel_case
from api_client import APIClient
from typing import Literal
from data import OVERRIDES
from utils.quicktype import run_quicktype


PARAM_TYPES = {
    "string": {
        "type": "string",
    },
    "number": {
        "type": "int",
    },
    "boolean": {
        "type": "bool",
    },
    "numbers": {
        "type": "[]int",
    },
}


class EndpointParameter:
    name: str
    type: str
    required: bool
    notes: list[str]

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

    def get_param_struct_row(self) -> str:
        go_type = PARAM_TYPES[self.type]["type"]
        return (
            f'{to_camel_case(self.name)} {go_type} `url:"{self.name},omitempty,comma"`'
        )


class Endpoint:
    api_client: APIClient

    category: str
    name: str

    link: str
    notes: list[str]
    parameters: list[EndpointParameter]

    format: Literal["json", "csv"] = "json"
    s3_cache: bool

    sample_responses_parsed: bool = False
    chunks_sampled: bool = False

    response_struct_name: str | None = None
    chunk_struct_name: str | None = None
    response_struct: str | None = None

    parameters_struct_name: str | None = None
    parameters_struct: str | None = None

    required_imports: set[str]

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

        self.required_imports = set()
        self.s3_cache = (
            OVERRIDES.get(self.category, {}).get(self.name, {}).get("s3_cache", True)
        )

        self.generate_params_struct()

    def generate_params_struct(self) -> str | None:
        if not self.parameters or len(self.parameters) == 0:
            return None

        # Generate the struct name
        self.parameters_struct_name = to_camel_case(
            f"{self.category}__{self.name}__Params"
        )

        # Generate the struct body
        struct_rows = [
            parameter.get_param_struct_row() for parameter in self.parameters
        ]
        struct_body = "\n    ".join(struct_rows)
        self.parameters_struct = (
            f"type {self.parameters_struct_name} struct {{\n    {struct_body}\n}}"
        )

    def add_required_import(self, import_path: str):
        self.required_imports.add(import_path)

    def fetch_sample_responses(self, cached: bool = True) -> bool:
        # TODO: handle driver_stats_by_category
        if self.category == "driver_stats_by_category":
            logging.warning(
                f"Skipping sample response fetch for {self.category}__{self.name} due to complexity."
            )
            return False

        sample_data_suite = (
            OVERRIDES.get(self.category, {}).get(self.name, {}).get("params", [])
        )

        if len(sample_data_suite) == 0:
            # Skip if some parameters are required but have no sample value
            has_required_params = any(
                parameter.required for parameter in self.parameters
            )
            if has_required_params:
                logging.error(
                    f"Skipping sample response fetch for {self.category}__{self.name} due to parameters."
                )
                return False

            sample_data_suite = [{}]

        for i, sample_data in enumerate(sample_data_suite):
            file_path = (
                f"output/responses/{self.category}__{self.name}__{i}.{self.format}"
            )
            if cached and os.path.exists(file_path):
                continue

            try:
                sample_response = self.api_client.call_endpoint(
                    self.link,
                    params=sample_data,
                    s3_cache=self.s3_cache,
                )
            except Exception as e:
                logging.error(f"Error fetching {self.category}__{self.name}__{i}: {e}")
                return False

            # Save the response to a file
            os.makedirs("output/responses", exist_ok=True)
            with open(file_path, "w") as f:
                f.write(sample_response)

        self.sample_responses_parsed = True
        return True

    def fetch_sample_chunks(self, cached: bool = True) -> bool:
        MAX_CHUNKS_PER_SAMPLE = 5

        if not self.sample_responses_parsed:
            logging.warning(
                f"Skipped: {self.category}__{self.name} (no sample response)"
            )
            return False

        sample_files = [
            f
            for f in os.listdir("output/responses")
            if f.startswith(f"{self.category}__{self.name}__")
            and f.endswith(f".{self.format}")
        ]

        for sample_index, sample_file in enumerate(sample_files):
            with open(os.path.join("output/responses", sample_file), "r") as f:
                sample_data = json.loads(f.read())

            # Check if data is a list
            if isinstance(sample_data, list) or not sample_data.get("chunk_info"):
                continue

            for chunk_index, chunk_file_name in enumerate(
                sample_data["chunk_info"]["chunk_file_names"]
            ):
                if chunk_index >= MAX_CHUNKS_PER_SAMPLE:
                    break

                file_path = f"output/responses/chunks__{self.category}__{self.name}__{sample_index}__{chunk_index}.json"
                if cached and os.path.exists(file_path):
                    self.chunks_sampled = True
                    continue

                url = (
                    f"{sample_data['chunk_info']['base_download_url']}{chunk_file_name}"
                )
                try:
                    sample_response = requests.get(url).text
                except Exception as e:
                    logging.error(
                        f"Error fetching {self.category}__{self.name}__{sample_index} chunk {chunk_index}: {e}"
                    )
                    continue

                # Save the response to a file
                os.makedirs("output/responses", exist_ok=True)
                with open(file_path, "w") as f:
                    f.write(sample_response)

                self.chunks_sampled = True

        return True

    def generate_go_types(self):
        if not self.sample_responses_parsed:
            logging.warning(
                f"Skipped: {self.category}__{self.name} (no sample response)"
            )
            return

        # Generate the files used as input for quicktype
        for f in os.listdir("output/responses"):
            if not f.startswith(f"{self.category}__{self.name}__"):
                continue

            if not f.endswith(f".{self.format}"):
                continue

            if self.chunks_sampled:
                with open(os.path.join("output/responses", f), "r") as file:
                    data = json.load(file)

                data["chunk_info"] = "PLACEHOLDER"
                data["chunks"] = "PLACEHOLDER"

                with open(os.path.join("output/inputs", f), "w") as file:
                    json.dump(data, file, indent=4)
            else:
                shutil.copyfile(
                    os.path.join("output/responses", f),
                    os.path.join("output/inputs", f),
                )

        # Generate the response struct
        self.response_struct_name = to_camel_case(
            f"{self.category}__{self.name}__Response"
        )

        self.response_struct, self.required_imports = run_quicktype(
            self.response_struct_name,
            f"output/inputs/{self.category}__{self.name}__*",
        )

        # Generate the chunk response struct
        if self.chunks_sampled:
            plural_chunk_struct_name = to_camel_case(
                f"{self.category}__{self.name}__Response__Chunks"
            )

            chunk_struct, chunk_imports = run_quicktype(
                plural_chunk_struct_name,
                f"output/responses/chunks__{self.category}__{self.name}__*",
            )

            self.chunk_struct_name = plural_chunk_struct_name.rstrip("s")

            # Find the Chunks and ChunkInfo fields and replace their types
            pattern = r"Chunks [^\n]+string"
            replacement = f"Chunks []{self.chunk_struct_name}"
            self.response_struct = re.sub(pattern, replacement, self.response_struct)

            pattern = r"ChunkInfo [^\n]+string"
            replacement = f"ChunkInfo client.IRacingChunkInfo"
            self.response_struct = re.sub(pattern, replacement, self.response_struct)

            # Remove the plural chunk struct
            chunk_struct = "\n".join(
                [
                    x
                    for x in chunk_struct.split("\n")
                    if plural_chunk_struct_name not in x
                ]
            )

            # Append the chunk struct to the response struct
            self.response_struct += f"\n\n{chunk_struct}"
            self.required_imports.update(chunk_imports)
            self.required_imports.add("github.com/riccardotornesello/irapi-go/client")


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
        lambda endpoint: endpoint.fetch_sample_responses(cached=skip_cached),
        endpoints,
        max_workers=workers,
    )

    logging.info("Sample responses fetched.")


def fetch_sample_chunks(
    endpoints: list[Endpoint], skip_cached: bool = True, workers: int = 5
):
    logging.info("Fetching sample chunks...")

    thread_map(
        lambda endpoint: endpoint.fetch_sample_chunks(cached=skip_cached),
        endpoints,
        max_workers=workers,
    )

    logging.info("Sample chunks fetched.")


def generate_go_types(endpoints: list[Endpoint], workers: int = 20):
    logging.info("Generating Go types...")

    # Remove and recreate the inputs folder
    if os.path.exists("output/inputs"):
        shutil.rmtree("output/inputs")
    os.makedirs("output/inputs", exist_ok=True)

    thread_map(
        lambda endpoint: endpoint.generate_go_types(),
        endpoints,
        max_workers=workers,
    )

    logging.info("Go types generated.")
