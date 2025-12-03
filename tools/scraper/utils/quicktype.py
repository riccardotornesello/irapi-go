"""
Wrapper module for the quicktype code generation tool.

This module provides a Python interface to quicktype, a tool that generates
strongly-typed Go structs from JSON samples. It handles:
- Running quicktype as a subprocess via npx
- Parsing the output to extract imports and struct definitions
- Error handling and logging
"""

import logging
import subprocess


def run_quicktype(
    struct_name: str,
    source_files: str,
    src_lang: str = "json",
) -> tuple[set[str], str] | tuple[None, None]:
    """
    Run the quicktype command to generate Go types from JSON files.
    
    Uses quicktype to analyze JSON files and generate corresponding Go struct
    definitions with proper types inferred from the data.
    
    Args:
        struct_name (str): Name for the root struct to generate.
        source_files (str): Glob pattern for JSON source files to analyze.
        src_lang (str): Source language of the input files (default: "json").
        
    Returns:
        tuple[str, set[str]] | tuple[None, None]: A tuple containing:
            - The generated Go struct code (without imports)
            - A set of required Go import paths
            Returns (None, None) if generation fails.
    """

    # Quicktype command
    command = [
        "npx",
        "quicktype",
        "--src",
        source_files,
        "--src-lang",
        src_lang,
        "--lang",
        "go",
        "--just-types",
        "-t",
        struct_name,
        "--no-enums",
    ]

    try:
        # Run the command and capture the output. Pass the sample response as stdin.
        result = subprocess.run(
            " ".join(command),
            capture_output=True,
            text=True,
            check=True,
            shell=True,
        )

    except subprocess.CalledProcessError as e:
        logging.error(f"FAILED: Quicktype for {struct_name} failed.")
        logging.error(f"Error: {e.stderr.strip()}")
        return (None, None)

    except FileNotFoundError:
        logging.error(
            "ERROR: Make sure 'npx' and 'quicktype' are installed and available in the PATH."
        )
        return (None, None)

    # Parse the output to extract imports and the struct definition
    required_imports = set()
    for line in result.stdout.splitlines():
        if line.startswith("import"):
            new_import = line.replace("import", "").replace('"', "").strip()
            required_imports.add(new_import)
            continue

    response_struct = "\n".join(
        [line for line in result.stdout.splitlines() if not line.startswith("import")]
    )

    return response_struct, required_imports
