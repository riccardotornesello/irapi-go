import os
import json

from genson import SchemaBuilder


def gen_schemas():
    for file in os.listdir("output/responses"):
        if not file.endswith(".json"):
            continue

        category = file.split("__")[0]
        endpoint = file.split("__")[1].split(".")[0]

        builder = SchemaBuilder()

        # Manual adjustements
        # TODO: move to a separate file
        if category == "car" and endpoint == "assets":
            builder.add_schema(
                {"type": "object", "patternProperties": {r"^\d+$": None}}
            )
        elif category == "series" and endpoint == "seasons":
            builder.add_schema(
                {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "allowed_season_members": {
                                "type": "object",
                                "patternProperties": {r"^\d+$": None},
                            }
                        },
                    },
                },
            )

        builder.add_object(json.load(open(f"output/responses/{file}")))

        schema = builder.to_schema()

        with open(f"output/schemas/{file}", "w") as f:
            f.write(json.dumps(schema, indent=2))
