"""
These functions perform requests to the iRacing API and save the responses to files.
These responses are used to generate the schemas for the API calls.
"""

import os
import json

from data import OVERRIDES


def get_responses(session, endpoints, cached=False):
    # Save responses
    for category, category_endpoints in endpoints.items():
        for endpoint, data in category_endpoints.items():
            file_path = f"output/responses/{category}__{endpoint}.{data['format']}"
            if cached and os.path.exists(file_path):
                continue

            param_values = (
                OVERRIDES.get(category, {}).get(endpoint, {}).get("params", {})
            )

            if any(p["required"] for p in data["parameters"]) and not param_values:
                print(f"Skipping {category}__{endpoint} because it requires parameters")
                continue

            if category == "results" and endpoint in ["search_hosted", "search_series"]:
                # TODO: manage this
                continue

            print(f"Fetching response for {category}__{endpoint}")

            url = data["link"]
            response = session.get(url, params=param_values)

            if response.status_code != 200:
                raise Exception(
                    f"Failed to fetch response for {category}__{endpoint}: {response.text}"
                )

            if data["s3_cache"]:
                s3_url = response.json()["link"]
                response = session.get(s3_url)

            if data["format"] == "json":
                if "chunk_info" in response.json():
                    chunk_data = []
                    for chunk_name in response.json()["chunk_info"]["chunk_file_names"]:
                        chunk_response = session.get(
                            response.json()["chunk_info"]["base_download_url"]
                            + chunk_name
                        )
                        chunk_data += chunk_response.json()
                    with open(
                        f"output/responses/{category}__{endpoint}__chunks.{data['format']}",
                        "w",
                    ) as f:
                        f.write(json.dumps(chunk_data))

            with open(file_path, "w") as f:
                f.write(response.text)
