# Development Container

This directory contains the development container configuration for the iRacing API Go Client project.

## What is a Development Container?

A development container (or dev container) is a fully featured development environment that runs in a Docker container. It provides a consistent, reproducible development setup across different machines.

## Features

This development container includes:

- **Go 1.23**: The Go programming language matching the version specified in `go.mod`
- **Python 3.11**: Required for the code generation tools in the `tools/` directory
- **VS Code Extensions**: Pre-configured extensions for Go and Python development
- **Automatic Setup**: Automatically downloads Go modules and installs Python dependencies on container creation

## Prerequisites

To use this development container, you need:

1. [Docker](https://www.docker.com/get-started) installed on your machine
2. [Visual Studio Code](https://code.visualstudio.com/) with the [Dev Containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)

Alternatively, you can use:
- [GitHub Codespaces](https://github.com/features/codespaces) (works directly in the browser)

## Usage

### With VS Code

1. Open the project folder in VS Code
2. When prompted, click "Reopen in Container" (or press F1 and select "Dev Containers: Reopen in Container")
3. VS Code will build the container and set up the environment
4. Once ready, you can start developing!

### With GitHub Codespaces

1. Navigate to the repository on GitHub
2. Click the "Code" button
3. Select the "Codespaces" tab
4. Click "Create codespace on [branch name]"
5. Wait for the environment to be set up

## What Happens on Container Creation

When the container is created, it automatically:

1. Downloads all Go module dependencies (`go mod download`)
2. Installs Python packages required by the tools (`pip install -r tools/requirements.txt`)

## Installed Tools and Extensions

### Go Tools
- Go language server
- Debugging support
- Code formatting and linting

### Python Tools
- Python language server (Pylance)
- Linting and formatting tools

### VS Code Extensions
- `golang.go` - Go language support
- `ms-python.python` - Python language support
- `ms-python.vscode-pylance` - Python IntelliSense

## Customization

You can customize the development container by editing the `devcontainer.json` file. See the [Dev Containers documentation](https://containers.dev/) for more information.
