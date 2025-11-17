import subprocess
import os


def to_camel_case(snake_str):
    return "".join(x.title() for x in snake_str.split("_"))


def format_go_code():
    # Get the current working directory
    target_dir = os.getcwd()

    # Go up two levels
    for _ in range(2):
        target_dir = os.path.dirname(target_dir)

    # Run the 'go fmt' command in the target directory
    command = ["go", "fmt", "./..."]

    subprocess.run(
        command,
        cwd=target_dir,
        capture_output=True,
        text=True,
        check=True,
    )
