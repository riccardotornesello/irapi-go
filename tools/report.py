import csv
import json
import os


def generate_report():
    endpoints = json.load(open("output/endpoints.json"))

    with open("output/report.csv", "w") as f:
        writer = csv.writer(f)
        writer.writerow(
            ["Category", "Endpoint", "Has params", "Body parsed", "Has chunks"]
        )
        for category, category_endpoints in endpoints.items():
            for endpoint, endpoint_data in category_endpoints.items():
                has_params = len(endpoint_data["parameters"]) > 0
                has_chunks = os.path.exists(
                    f"output/schemas/{category}__{endpoint}__chunks.json"
                )
                body_parsed = os.path.exists(
                    f"output/schemas/{category}__{endpoint}.{endpoint_data['format']}"
                )
                writer.writerow(
                    [category, endpoint, has_params, body_parsed, has_chunks]
                )
