import logging
import subprocess


def run_quicktype(
    struct_name: str, source_files: str
) -> tuple[set[str], str] | tuple[None, None]:
    """
    Run the quicktype command to generate Go types.
    """

    # Quicktype command
    command = [
        "npx",
        "quicktype",
        "--src",
        source_files,
        "--src-lang",
        "json",
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
