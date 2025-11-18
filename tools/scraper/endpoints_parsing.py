"""
Module for parsing API endpoints and generating Go type definitions.

This module contains classes and functions for:
- Representing API endpoints and their parameters
- Fetching sample responses from endpoints
- Generating Go struct definitions from JSON responses
- Managing endpoint metadata and type information

The main classes are:
- EndpointParameter: Represents a single parameter for an endpoint
- Endpoint: Represents a complete API endpoint with all metadata
"""

import os
import logging
import shutil
import requests
import json
import re
import datetime

from tqdm.contrib.concurrent import thread_map

from format import to_camel_case
from api_client import APIClient
from typing import Literal
from data import OVERRIDES
from utils.quicktype import run_quicktype


# Mapping of iRacing parameter types to Go types
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
    """
    Represents a parameter for an API endpoint.
    
    Attributes:
        name (str): The parameter name as used in the API.
        type (str): The parameter type (string, number, boolean, numbers).
        required (bool): Whether this parameter is required for the API call.
        notes (list[str]): Additional notes about the parameter from documentation.
    """
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
        """
        Initialize an endpoint parameter from API documentation.
        
        Args:
            category_name (str): The category this endpoint belongs to.
            endpoint_name (str): The name of the endpoint.
            param_name (str): The name of this parameter.
            param_data (dict): Parameter metadata from API documentation.
            
        Raises:
            Exception: If the parameter type is not recognized.
        """
        self.name = param_name
        self.type = param_data["type"]
        self.required = param_data.get("required", False)
        self.notes = _parse_iracing_notes(param_data.get("notes"))

        # Check that the type is known
        if self.type not in PARAM_TYPES.keys():
            raise Exception(f"Unknown type: {self.type}")

    def get_param_struct_row(self) -> str:
        """
        Generate a Go struct field definition for this parameter.
        
        Returns:
            str: Go struct field with proper type and tags for URL encoding.
        """
        go_type = PARAM_TYPES[self.type]["type"]
        return (
            f'{to_camel_case(self.name)} {go_type} `url:"{self.name},omitempty,comma"`'
        )


class Endpoint:
    """
    Represents an API endpoint with all its metadata and generated code.
    
    This class manages the complete lifecycle of an endpoint from documentation
    to code generation, including fetching sample responses and generating Go types.
    
    Attributes:
        api_client (APIClient): Client for making API requests.
        category (str): Category name (e.g., 'league', 'member').
        name (str): Endpoint name within the category.
        link (str): Full URL of the API endpoint.
        notes (list[str]): Documentation notes for this endpoint.
        parameters (list[EndpointParameter]): List of parameters this endpoint accepts.
        format (Literal["json", "csv"]): Response format (default: "json").
        s3_cache (bool): Whether responses are cached in S3.
        sample_responses_parsed (bool): Whether sample responses have been fetched.
        chunks_sampled (bool): Whether chunked responses have been sampled.
        response_struct_name (str | None): Name of the generated Go response struct.
        chunk_struct_name (str | None): Name of the generated Go chunk struct.
        response_struct (str | None): Generated Go struct code for responses.
        parameters_struct_name (str | None): Name of the generated Go parameters struct.
        parameters_struct (str | None): Generated Go struct code for parameters.
        required_imports (set[str]): Set of Go package imports needed.
    """
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
        """
        Initialize an endpoint from API documentation data.
        
        Args:
            api_client (APIClient): Authenticated API client for requests.
            category_name (str): The category this endpoint belongs to.
            endpoint_name (str): The name of this endpoint.
            endpoint_data (dict): Endpoint metadata from API documentation.
        """
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
        """
        Generate a Go struct definition for endpoint parameters.
        
        Creates a struct with fields for each parameter, including proper
        Go naming conventions and URL encoding tags.
        
        Returns:
            str | None: The generated struct code, or None if no parameters exist.
        """
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
        """
        Add a Go package import path to the required imports set.
        
        Args:
            import_path (str): Full Go package import path.
        """
        self.required_imports.add(import_path)

    def fetch_sample_responses(self, cached: bool = True) -> bool:
        """
        Fetch sample responses from the API endpoint.
        
        Retrieves actual responses from the endpoint using sample parameters
        defined in the OVERRIDES configuration. Saves responses to disk for
        type generation.
        
        Args:
            cached (bool): If True, skip fetching if responses already exist. 
                          Defaults to True.
        
        Returns:
            bool: True if successful, False if fetching failed or was skipped.
        """
        # TODO: handle driver_stats_by_category
        if self.category == "driver_stats_by_category":
            logging.warning(
                f"Skipping sample response fetch for {self.category}__{self.name} due to complexity."
            )
            return False

        responses_folder = f"output/responses/{self.category}/{self.name}"
        os.makedirs(responses_folder, exist_ok=True)

        # If the cache is enabled and we have files, skip fetching
        if len(os.listdir(responses_folder)) > 0 and cached:
            self.sample_responses_parsed = True

            # If any of the files starts with "chunks__", we have sampled chunks
            for f in os.listdir(responses_folder):
                if f.startswith("chunks__"):
                    self.chunks_sampled = True
                    break

            return True

        sample_data_suite = (
            OVERRIDES.get(self.category, {}).get(self.name, {}).get("params", [])
        )

        # Skip if some parameters are required but have no sample value
        if len(sample_data_suite) == 0:
            has_required_params = any(
                parameter.required for parameter in self.parameters
            )
            if has_required_params:
                logging.error(
                    f"Skipping sample response fetch for {self.category}__{self.name} due to parameters."
                )
                return False

            # Use an empty dict as default sample data
            sample_data_suite = [{}]

        timestamp = datetime.datetime.now().strftime("%Y%m%d%H%M%S")

        for i, sample_data in enumerate(sample_data_suite):
            file_path = f"output/responses/{self.category}/{self.name}/response__{timestamp}__{i}.{self.format}"

            try:
                sample_response = self.api_client.call_endpoint(
                    self.link,
                    params=sample_data,
                    s3_cache=self.s3_cache,
                )
            except Exception as e:
                logging.error(f"Error fetching {self.category}/{self.name}/{i}: {e}")
                return False

            # Save the response to a file
            with open(file_path, "w") as f:
                f.write(sample_response)

            # Save the chunks if any
            response_data = json.loads(sample_response)
            if isinstance(response_data, dict) and response_data.get("chunk_info"):
                if not self.fetch_sample_chunks(response_data["chunk_info"], timestamp, i):
                    return False

        self.sample_responses_parsed = True
        return True

    def fetch_sample_chunks(
        self, chunk_info: dict, timestamp: str, sample_index: int
    ) -> bool:
        """
        Fetch chunked response data from S3.
        
        Some endpoints return large datasets split into multiple chunks.
        This method fetches a sample of those chunks for type analysis.
        
        Args:
            chunk_info (dict): Chunk metadata from the API response.
            timestamp (str): Timestamp for file naming.
            sample_index (int): Index of the sample being processed.
            
        Returns:
            bool: True if successful, False on error.
        """
        MAX_CHUNKS_PER_SAMPLE = 5

        base_url = chunk_info["base_download_url"]

        for chunk_index, chunk_file_name in enumerate(chunk_info["chunk_file_names"]):
            if chunk_index >= MAX_CHUNKS_PER_SAMPLE:
                break

            file_path = f"output/responses/{self.category}/{self.name}/chunks__{timestamp}__{sample_index}__{chunk_index}.{self.format}"

            url = f"{base_url}{chunk_file_name}"
            try:
                sample_response = requests.get(url).text
            except Exception as e:
                logging.error(
                    f"Error fetching {self.category}__{self.name}__{sample_index} chunk {chunk_index}: {e}"
                )
                return False

            # Save the response to a file
            with open(file_path, "w") as f:
                f.write(sample_response)

            self.chunks_sampled = True

        return True

    def generate_go_types(self):
        """
        Generate Go type definitions from sample responses.
        
        Uses quicktype to analyze JSON responses and generate strongly-typed
        Go structs. Handles both regular responses and chunked responses,
        properly linking chunk types to the main response struct.
        """
        if not self.sample_responses_parsed:
            logging.warning(
                f"Skipped: {self.category}__{self.name} (no sample response)"
            )
            return

        # Generate the files used as input for quicktype
        responses_folder = f"output/responses/{self.category}/{self.name}"
        inputs_folder = f"output/inputs/{self.category}/{self.name}"
        os.makedirs(inputs_folder, exist_ok=True)

        for f in os.listdir(responses_folder):
            if not f.startswith(f"response__"):
                continue

            if not f.endswith(f".{self.format}"):
                continue

            if self.chunks_sampled:
                with open(os.path.join(responses_folder, f), "r") as file:
                    data = json.load(file)

                data["chunk_info"] = "PLACEHOLDER"
                data["chunks"] = "PLACEHOLDER"

                with open(os.path.join(inputs_folder, f), "w") as file:
                    json.dump(data, file, indent=4)
            else:
                shutil.copyfile(
                    os.path.join(responses_folder, f),
                    os.path.join(inputs_folder, f),
                )

        # Generate the response struct
        self.response_struct_name = to_camel_case(
            f"{self.category}__{self.name}__Response"
        )

        self.response_struct, self.required_imports = run_quicktype(
            self.response_struct_name,
            f"output/inputs/{self.category}/{self.name}/response__*",
        )

        # Generate the chunk response struct
        if self.chunks_sampled:
            plural_chunk_struct_name = to_camel_case(
                f"{self.category}__{self.name}__Response__Chunks"
            )

            chunk_struct, chunk_imports = run_quicktype(
                plural_chunk_struct_name,
                f"output/responses/{self.category}/{self.name}/chunks__*",
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
    Parse notes from the iRacing API documentation.
    
    Gets the notes from endpoint or parameter data and returns a list of notes.
    Handles both single notes and arrays of notes.
    
    Args:
        notes: Notes from API documentation (string, list, or None).
        
    Returns:
        list[str]: List of note strings, empty if no notes exist.
    """

    if not notes:
        return []

    if not isinstance(notes, list):
        notes = [notes]

    return notes


def fetch_sample_responses(
    endpoints: list[Endpoint], skip_cached: bool = True, workers: int = 5
):
    """
    Fetch sample responses for all endpoints in parallel.
    
    Uses a thread pool to efficiently fetch responses from multiple endpoints
    simultaneously. Shows progress bar during execution.
    
    Args:
        endpoints (list[Endpoint]): List of endpoints to fetch responses for.
        skip_cached (bool): If True, skip endpoints with existing responses. 
                           Defaults to True.
        workers (int): Number of parallel workers to use. Defaults to 5.
    """
    logging.info("Fetching sample responses...")

    thread_map(
        lambda endpoint: endpoint.fetch_sample_responses(cached=skip_cached),
        endpoints,
        max_workers=workers,
    )

    logging.info("Sample responses fetched.")


def generate_go_types(endpoints: list[Endpoint], workers: int = 20):
    """
    Generate Go type definitions for all endpoints in parallel.
    
    Processes sample responses to create strongly-typed Go structs using
    quicktype. Runs in parallel for efficiency.
    
    Args:
        endpoints (list[Endpoint]): List of endpoints to generate types for.
        workers (int): Number of parallel workers to use. Defaults to 20.
    """
    logging.info("Generating Go types...")

    # Remove the inputs folder
    if os.path.exists("output/inputs"):
        shutil.rmtree("output/inputs")

    thread_map(
        lambda endpoint: endpoint.generate_go_types(),
        endpoints,
        max_workers=workers,
    )

    logging.info("Go types generated.")
