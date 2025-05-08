"""
These functions converts the iRacing API definition into a more usable format.
"""

import json
import os

from constants import PARAM_TYPES
from data import OVERRIDES


def parse_notes(endpoint_data):
    """
    Gets the notes from the endpoint or the parameter's data and returns a list of notes.
    The list is empty if there are no notes.
    """

    notes = endpoint_data.get("note")

    if not notes:
        return []

    if not isinstance(notes, list):
        notes = [notes]

    return notes


def parse_parameters(endpoint_data):
    parameters = endpoint_data.get("parameters", {})

    # Check that the types are known
    for v in parameters.values():
        if v["type"] not in PARAM_TYPES.keys():
            raise Exception(f"Unknown type: {v['type']}")

    return [
        {
            "key": k,
            "type": v["type"],
            "required": v.get("required", False),
            "notes": parse_notes(v),
        }
        for k, v in parameters.items()
    ]


def get_api_definition(session, cached=False):
    file_path = "output/api_definition.json"
    if cached and os.path.exists(file_path):
        return json.load(open(file_path))

    res = session.get("https://members-ng.iracing.com/data/doc")
    if res.status_code != 200:
        raise Exception(f"Failed to fetch API definition: {res.text}")

    api_definition = res.json()
    with open(file_path, "w") as f:
        json.dump(api_definition, f, indent=2)

    return api_definition


def generate_endpoints(api_definition):
    endpoints = {}

    for category, category_endpoints in api_definition.items():
        endpoints[category] = {}

        for endpoint, data in category_endpoints.items():
            endpoint_overrides = OVERRIDES.get(category, {}).get(endpoint, {})

            endpoints[category][endpoint] = {
                "link": data["link"],
                "notes": parse_notes(data),
                "parameters": parse_parameters(data),
                "format": endpoint_overrides.get("format", "json"),
                "s3_cache": endpoint_overrides.get("s3_cache", True),
            }

    with open("output/endpoints.json", "w") as f:
        json.dump(endpoints, f, indent=2)

    return endpoints
