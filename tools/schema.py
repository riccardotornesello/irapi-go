import os
import json

from genson import SchemaBuilder

from data import SCHEMA_ADJUSTMENTS


def gen_schemas():
    for file in os.listdir("output/responses"):
        if not file.endswith(".json"):
            continue

        category = file.split("__")[0]
        endpoint = file.split("__")[1].split(".")[0]

        builder = SchemaBuilder()

        # Manual adjustements
        if category in SCHEMA_ADJUSTMENTS and endpoint in SCHEMA_ADJUSTMENTS[category]:
            builder.add_schema(SCHEMA_ADJUSTMENTS[category][endpoint])

        builder.add_object(json.load(open(f"output/responses/{file}")))

        schema = builder.to_schema()

        with open(f"output/schemas/{file}", "w") as f:
            f.write(json.dumps(schema, indent=2))
