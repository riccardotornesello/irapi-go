"""
This module contains functions to generate Go structs from JSON schemas.
"""

from constants import PARAM_TYPES


def generate_key(key):
    if key.isdigit():
        key = f"Field{key}"

    return "".join(x.capitalize() for x in key.lower().replace("-", "_").split("_"))


def convert_schema_to_struct(
    schema, json_name=None, chunks_struct_name=None, flatten_array=False
):
    data = ""

    if "anyOf" in schema:
        if len(schema["anyOf"]) != 2:
            raise Exception(f"Too many possible schemas {schema}")
        if schema["anyOf"][0]["type"] != "null":
            raise Exception(f"Unknown list schema {schema}")
        return f"*{convert_schema_to_struct(schema['anyOf'][1], json_name)}"

    if json_name == "chunk_info":
        return f'client.IRacingChunkInfo `json:"{json_name}"`\n'

    if "type" not in schema:
        raise Exception(f"Missing type in key {json_name}: {schema}")

    t = schema["type"]

    if isinstance(t, list):
        if "null" in t:
            t.remove("null")
            data += "*"
        else:
            data += ""

        if len(t) != 1:
            raise Exception(f"Unknown type {t} in key {json_name}")

        t = t[0]

    if flatten_array and not t == "array":
        raise Exception(
            f"Flatten array is only supported for array types in key {json_name}"
        )

    if t == "array":
        if not flatten_array:
            data += "[]"

        if not "items" in schema:
            print(f"Warning: array without items in key {json_name}")
            data += "interface{}"
        else:
            data += convert_schema_to_struct(schema["items"])

    elif t == "object":
        if "properties" in schema:
            data += "struct {\n"

            for k, v in schema["properties"].items():
                data += f"{generate_key(k)} {convert_schema_to_struct(v,k)}"

            if chunks_struct_name:
                data += f"Chunks []{chunks_struct_name}"

            data += "}"
        elif "patternProperties" in schema:
            keys = list(schema["patternProperties"].keys())
            if len(keys) != 1:
                raise Exception(f"Unknown patternProperties {keys} in key {json_name}")

            data += "map[string]"
            data += convert_schema_to_struct(schema["patternProperties"][keys[0]])
        else:
            print("Warning: object without properties")
            data += "interface{}"

    elif t == "integer":
        if json_name is not None and (
            json_name.endswith("_pct")
            or json_name.startswith("pct_")
            or "_pct_" in json_name
        ):
            data += "float64"
        else:
            data += "int"

    elif t == "string":
        data += "string"
    elif t == "number":
        data += "float64"
    elif t == "boolean":
        data += "bool"
    elif t == "null":
        print(f"Warning: null type in key {json_name}")
        data += "interface{}"
    else:
        raise Exception(f"Unknown type {t} in key {json_name}")

    if json_name:
        data += f' `json:"{json_name}"`\n'

    return data


def generate_parameters_struct(parameters):
    if len(parameters) == 0:
        return None

    data = "struct {\n"

    for param in parameters:
        required = param["required"] != False

        type_string = PARAM_TYPES[param["type"]]
        if not required:
            type_string = "*" + type_string

        url_param = param["key"]
        if not required:
            url_param = url_param + ",omitempty"

        data += f'{generate_key(param["key"])} {type_string} `url:"{url_param}"`'

        if "notes" in param:
            notes_string = " ".join(param["notes"]).replace("\n", " ").strip()
            if len(notes_string) > 0:
                data += f" // {notes_string}"

        data += "\n"

    data += "}\n"

    return data
