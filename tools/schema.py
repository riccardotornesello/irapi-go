"""
Generate JSON schemas from the responses.

This module processes API response JSON files and generates JSON schemas using the
genson library. It also post-processes the schemas to automatically detect objects
with numeric string keys and converts them to use patternProperties (which will
generate Go maps instead of structs).
"""

import os
import json
import re

from genson import SchemaBuilder

from data import OVERRIDES


# Regex pattern to match numeric string keys (e.g., "1", "2", "100")
NUMERIC_KEY_PATTERN = r"^\d+$"


def all_keys_are_numeric_strings(properties):
    """
    Check if all keys in a properties dict are string representations of numbers.
    
    Uses the NUMERIC_KEY_PATTERN regex to ensure consistency with patternProperties
    generation. This matches strings containing only digits (0-9).
    
    Args:
        properties (dict): The properties dictionary from a JSON schema object
        
    Returns:
        bool: True if all keys are numeric strings, False otherwise
        
    Examples:
        {"1": {...}, "2": {...}, "100": {...}} -> True
        {"name": {...}, "age": {...}} -> False
        {"1": {...}, "name": {...}} -> False
    """
    if not properties or len(properties) == 0:
        return False
    
    return all(re.match(NUMERIC_KEY_PATTERN, key) for key in properties.keys())


def convert_numeric_objects_to_maps(schema):
    """
    Recursively process a schema and convert objects with all-numeric keys to patternProperties.
    
    This function walks through the schema tree and identifies objects where all property
    keys are numeric strings. These are converted to use patternProperties with a pattern
    matching numeric strings, which will generate a Go map[string]T instead of a struct.
    
    Args:
        schema (dict): The JSON schema to process (modified in place)
        
    Returns:
        dict: The modified schema
        
    Note:
        This modifies the schema in place and also returns it for convenience.
    """
    if not isinstance(schema, dict):
        return schema
    
    # Process object types with properties
    if schema.get("type") == "object" and "properties" in schema:
        properties = schema["properties"]
        
        # Check if all keys are numeric strings
        if all_keys_are_numeric_strings(properties):
            # Get all the property schemas to merge them
            property_schemas = list(properties.values())
            
            if len(property_schemas) > 0:
                # Merge all property schemas to ensure the pattern value type
                # accommodates all possible property types
                builder = SchemaBuilder()
                for prop_schema in property_schemas:
                    # Add each property schema to merge them
                    builder.add_schema(prop_schema)
                
                value_schema = builder.to_schema()
                
                # Convert to patternProperties format with numeric string pattern
                schema["patternProperties"] = {NUMERIC_KEY_PATTERN: value_schema}
                del schema["properties"]
                
                # Recursively process the value schema
                convert_numeric_objects_to_maps(value_schema)
                
                return schema
    
    # Recursively process all nested schemas
    for key, value in list(schema.items()):
        if isinstance(value, dict):
            convert_numeric_objects_to_maps(value)
        elif isinstance(value, list):
            for item in value:
                if isinstance(item, dict):
                    convert_numeric_objects_to_maps(item)
    
    return schema


def gen_schemas():
    """
    Generate JSON schemas from API response files.
    
    This function:
    1. Reads JSON response files from output/responses/
    2. Applies manual schema overrides from data.py
    3. Generates schemas using genson
    4. Automatically converts objects with all-numeric keys to use patternProperties
    5. Writes the processed schemas to output/schemas/
    """
    for file in os.listdir("output/responses"):
        if not file.endswith(".json"):
            continue

        category = file.split("__")[0]
        endpoint = file.split("__")[1].split(".")[0]

        builder = SchemaBuilder()

        # Manual adjustments from OVERRIDES in data.py
        schema_overrides = (
            OVERRIDES.get(category, {}).get(endpoint, {}).get("schema", {})
        )
        if schema_overrides:
            builder.add_schema(schema_overrides)

        builder.add_object(json.load(open(f"output/responses/{file}")))

        schema = builder.to_schema()
        
        # Automatically detect and convert objects with numeric keys to maps
        schema = convert_numeric_objects_to_maps(schema)

        with open(f"output/schemas/{file}", "w") as f:
            f.write(json.dumps(schema, indent=2))
