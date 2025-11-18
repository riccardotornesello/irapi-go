"""
Utility functions for formatting and code generation.

This module provides:
- String conversion utilities (snake_case to CamelCase)
- Go code formatting using 'go fmt'
"""

import subprocess
import os


def to_camel_case(snake_str):
    """
    Convert a snake_case string to CamelCase.
    
    Args:
        snake_str (str): String in snake_case format (e.g., 'get_user_info').
        
    Returns:
        str: String in CamelCase format (e.g., 'GetUserInfo').
    """
    return "".join(x.title() for x in snake_str.split("_"))


def format_go_code():
    """
    Format all generated Go code using 'go fmt'.
    
    Navigates to the repository root (two levels up from current directory)
    and runs 'go fmt ./...' to format all Go files according to Go standards.
    
    Raises:
        subprocess.CalledProcessError: If go fmt fails.
    """
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
