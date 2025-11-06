import os
import subprocess
import logging

from tqdm.contrib.concurrent import thread_map

from format import to_camel_case


INPUT_DIR = "output/responses"
OUTPUT_DIR = "output/types"


def generate_go_types(input_path: str):
    """
    Run the quicktype command to generate Go types from a single JSON file.
    """
    # Get the base name of the file (e.g. "car__assets.json")
    filename = os.path.basename(input_path)

    # Check if the file has a .json extension
    if not filename.lower().endswith(".json"):
        logging.warning(f"Skipped: {filename} (not a JSON file)")
        return

    # Create the output path by replacing the file extension with .go
    # The quicktype command wants the output file name (which will be the type name)
    # In this case, we keep the original file name
    output_path = os.path.join(OUTPUT_DIR, filename.replace(".json", ".go"))

    # Generate the struct name from the file name
    struct_name = to_camel_case(filename.replace(".json", "")) + "Response"

    # Quicktype command
    command = [
        "npx",
        "quicktype",
        "--src",
        input_path,
        "--out",
        output_path,
        "--src-lang",
        "json",
        "--lang",
        "go",
        "--just-types",
        "-t",
        struct_name,
    ]

    logging.debug(f"Generating type for: {filename}")

    try:
        # Run the command and capture the output
        subprocess.run(
            command,
            check=True,  # Raise an error if the command fails
            capture_output=True,
            text=True,
        )
        logging.debug(f"SUCCESS: Types generated in {output_path}")

    except subprocess.CalledProcessError as e:
        logging.error(f"FAILED: Quicktype for {filename} failed.")
        logging.error(f"Error: {e.stderr.strip()}")
    except FileNotFoundError:
        logging.error(
            "ERROR: Make sure 'npx' and 'quicktype' are installed and available in the PATH."
        )


def run(workers: int = 20):
    # Clean the output directory
    if not os.path.exists(OUTPUT_DIR):
        os.makedirs(OUTPUT_DIR)
    else:
        for f in os.listdir(OUTPUT_DIR):
            os.remove(os.path.join(OUTPUT_DIR, f))

    if not os.path.isdir(INPUT_DIR):
        raise FileNotFoundError(f"Input directory '{INPUT_DIR}' does not exist.")

    logging.info(f"Starting processing of JSON files in '{INPUT_DIR}'...")

    # Iterate over all items in the input directory
    thread_map(
        generate_go_types,
        [os.path.join(INPUT_DIR, filename) for filename in os.listdir(INPUT_DIR)],
        max_workers=workers,
    )

    logging.info("Processing complete.")


if __name__ == "__main__":
    logging.basicConfig(
        level=logging.INFO, format="%(asctime)s - %(levelname)s - %(message)s"
    )

    run()
