import json
import os


def parse_note(endpoint_data):
    note = endpoint_data.get("note")

    # If note is a list, concatenate it
    if isinstance(note, list):
        note = " ".join(note)

    return note


def parse_format(category):
    if category == "driver_stats_by_category":
        return "csv"
    return "json"


def parse_parameters(endpoint_data):
    parameters = endpoint_data.get("parameters", {})

    return [
        {
            "key": k,
            "type": v["type"],
            "required": v.get("required", False),
            "note": v.get("note"),
        }
        for k, v in parameters.items()
    ]


def skip_s3(category, endpoint):
    if category == "constants":
        return True

    if category == "member" and endpoint == "awards":
        return True

    return False


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
            endpoints[category][endpoint] = {
                "link": data["link"],
                "note": parse_note(data),
                "parameters": parse_parameters(data),
                "format": parse_format(category),
                "skip_s3": skip_s3(category, endpoint),
            }
    return endpoints
