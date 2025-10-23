"""
This module contains functions to generate Go structs from JSON schemas.

The main functionality includes:
- Converting JSON schema objects to Go struct definitions
- Generating Go struct field names from JSON property names
- Handling various JSON schema types (object, array, string, number, etc.)
- Converting objects with pattern properties to Go maps
"""

from constants import PARAM_TYPES


def generate_key(key):
    """
    Convert a JSON property name to a Go struct field name.
    
    This function transforms a snake_case or kebab-case JSON key into a PascalCase
    Go field name. If the key is purely numeric, it prefixes it with "Field".
    
    Args:
        key (str): The JSON property name to convert
        
    Returns:
        str: The Go struct field name in PascalCase
        
    Examples:
        "user_name" -> "UserName"
        "first-name" -> "FirstName"
        "123" -> "Field123"
    """
    if key.isdigit():
        key = f"Field{key}"

    return "".join(x.capitalize() for x in key.lower().replace("-", "_").split("_"))


def convert_schema_to_struct(
    schema, json_name=None, chunks_struct_name=None, flatten_array=False
):
    """
    Convert a JSON schema to a Go struct definition.
    
    This function recursively processes a JSON schema and generates the corresponding
    Go type definition. It handles various JSON schema constructs including objects,
    arrays, primitive types, and nullable types.
    
    Args:
        schema (dict): The JSON schema to convert
        json_name (str, optional): The JSON property name for this field (used for tags)
        chunks_struct_name (str, optional): Name of the chunks struct to append to objects
        flatten_array (bool): If True, flattens array types (doesn't add [] prefix)
        
    Returns:
        str: The Go type definition as a string
        
    Raises:
        Exception: If the schema contains unsupported or invalid constructs
    """
    data = ""

    # Handle nullable types using anyOf schema construct
    if "anyOf" in schema:
        if len(schema["anyOf"]) != 2:
            raise Exception(f"Too many possible schemas {schema}")
        if schema["anyOf"][0]["type"] != "null":
            raise Exception(f"Unknown list schema {schema}")
        return f"*{convert_schema_to_struct(schema['anyOf'][1], json_name)}"

    # Special handling for chunk_info field
    if json_name == "chunk_info":
        return f'client.IRacingChunkInfo `json:"{json_name}"`\n'

    # Validate that schema has a type field
    if "type" not in schema:
        raise Exception(f"Missing type in key {json_name}: {schema}")

    t = schema["type"]

    # Handle types specified as arrays (e.g., ["string", "null"])
    if isinstance(t, list):
        if "null" in t:
            t.remove("null")
            data += "*"  # Add pointer for nullable types
        else:
            data += ""

        if len(t) != 1:
            raise Exception(f"Unknown type {t} in key {json_name}")

        t = t[0]

    # Validate flatten_array is only used with array types
    if flatten_array and not t == "array":
        raise Exception(
            f"Flatten array is only supported for array types in key {json_name}"
        )

    # Convert array types
    if t == "array":
        if not flatten_array:
            data += "[]"

        if not "items" in schema:
            print(f"Warning: array without items in key {json_name}")
            data += "interface{}"
        else:
            data += convert_schema_to_struct(schema["items"])

    # Convert object types
    elif t == "object":
        if "properties" in schema:
            # Object with explicit properties -> Go struct
            data += "struct {\n"

            for k, v in schema["properties"].items():
                data += f"{generate_key(k)} {convert_schema_to_struct(v,k)}"

            if chunks_struct_name:
                data += f"Chunks []{chunks_struct_name}"

            data += "}"
        elif "patternProperties" in schema:
            # Object with pattern properties -> Go map
            keys = list(schema["patternProperties"].keys())
            if len(keys) != 1:
                raise Exception(f"Unknown patternProperties {keys} in key {json_name}")

            data += "map[string]"
            data += convert_schema_to_struct(schema["patternProperties"][keys[0]])
        else:
            print("Warning: object without properties")
            data += "interface{}"

    # Convert integer types
    elif t == "integer":
        # Special case: fields containing "pct" should be float64 for percentages
        if json_name is not None and (
            json_name.endswith("_pct")
            or json_name.startswith("pct_")
            or "_pct_" in json_name
        ):
            data += "float64"
        else:
            data += "int"

    # Convert primitive types
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

    # Add JSON struct tag if this is a named field
    if json_name:
        data += f' `json:"{json_name}"`\n'

    return data


def generate_parameters_struct(parameters):
    """
    Generate a Go struct definition for API request parameters.
    
    This function creates a struct that can be used with the go-querystring library
    to convert struct fields to URL query parameters.
    
    Args:
        parameters (list): List of parameter dictionaries with keys:
            - key: The parameter name
            - type: The parameter type
            - required: Whether the parameter is required (default True)
            - notes: Optional documentation notes
            
    Returns:
        str: The Go struct definition, or None if no parameters
        
    Example output:
        struct {
            CustId int `url:"cust_id"` // Customer ID
            IncludeLicenses *bool `url:"include_licenses,omitempty"`
        }
    """
    if len(parameters) == 0:
        return None

    data = "struct {\n"

    for param in parameters:
        required = param["required"] != False

        type_string = PARAM_TYPES[param["type"]]
        if not required:
            type_string = "*" + type_string  # Use pointer for optional parameters

        url_param = param["key"]
        if not required:
            url_param = url_param + ",omitempty"  # Omit empty optional parameters

        data += f'{generate_key(param["key"])} {type_string} `url:"{url_param}"`'

        # Add inline documentation from notes
        if "notes" in param:
            notes_string = " ".join(param["notes"]).replace("\n", " ").strip()
            if len(notes_string) > 0:
                data += f" // {notes_string}"

        data += "\n"

    data += "}\n"

    return data
