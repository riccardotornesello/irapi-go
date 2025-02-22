"""
This is used to generate a report of the endpoints that have been processed.
In later updates it should also generate a documentation.
"""

import json
import os


ok_emoji = "✅"
error_emoji = "❌"
null_emoji = "-"


def generate_report():
    endpoints = json.load(open("output/endpoints.json"))

    with open("output/report.md", "w") as f:

        f.write("# Report\n\n")

        f.write("## Endpoints\n\n")
        f.write("| Category | Endpoint | Has params | Body parsed | Has chunks |\n")
        f.write("|----------|----------|------------|-------------|------------|\n")
        for category, category_endpoints in endpoints.items():
            if category == "meta":
                continue
            for endpoint, endpoint_data in category_endpoints.items():
                has_params = len(endpoint_data["parameters"]) > 0
                has_chunks = os.path.exists(
                    f"output/schemas/{category}__{endpoint}__chunks.json"
                )
                body_parsed = os.path.exists(
                    f"output/schemas/{category}__{endpoint}.{endpoint_data['format']}"
                )
                f.write(
                    f"| {category} | {endpoint} | {ok_emoji if has_params else null_emoji} | {ok_emoji if body_parsed else null_emoji} | {ok_emoji if has_chunks else null_emoji} |\n"
                )
