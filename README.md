# iRacing API - Go Client

This is a Go client for the iRacing API.

It is capable of connecting to most of the API endpoints and parsing the responses into Go structs.

## Usage

```go
package main

import (
	irapi "github.com/riccardotornesello/irapi-go"
	"github.com/riccardotornesello/irapi-go/api/member/get"
)

func main() {
	// Create a client using OAuth credentials
	irClient, err := irapi.NewIRacingPasswordLimitedApiClient(
		"your-client-id",
		"your-client-secret",
		"your-email@example.com",
		"your-password",
	)
	if err != nil {
		panic(err)
	}

	// Get member info
	params := &get.MemberGetParams{
		CustIds: []int{911231},
	}
	userInfo, err := irClient.Member.Get(params)
	if err != nil {
		panic(err)
	}

	// Print the member info
	for _, member := range userInfo.Members {
		println(member.DisplayName)
	}
}
```

## Logging

The client uses structured logging with support for multiple log levels and output formats. By default, the logger outputs INFO level messages in text format to stderr.

### Configuration Options

```go
// Enable debug logging
irapi.ConfigureLogger(&irapi.LoggerOptions{
    Level: irapi.LogLevelDebug,
})

// Enable JSON logging for production
irapi.ConfigureLogger(&irapi.LoggerOptions{
    Level: irapi.LogLevelInfo,
    JSON:  true,
})

// Custom output writer
irapi.ConfigureLogger(&irapi.LoggerOptions{
    Level:  irapi.LogLevelInfo,
    Output: myCustomWriter,
})

// Disable all logging
irapi.DisableLogging()
```

### Log Levels

- `LogLevelDebug`: Detailed debugging information (e.g., rate limit pauses)
- `LogLevelInfo`: General operational information (e.g., token refresh events)
- `LogLevelWarn`: Warning messages (e.g., rate limit exceeded)
- `LogLevelError`: Error messages

### Example Log Output

Text format:
```
time=2024-01-15T10:30:00.000Z level=INFO msg="refreshing access token"
time=2024-01-15T10:30:01.000Z level=INFO msg="token refreshed successfully" expiry=2024-01-15T11:30:01.000Z
```

JSON format:
```json
{"time":"2024-01-15T10:30:00.000Z","level":"INFO","msg":"refreshing access token"}
{"time":"2024-01-15T10:30:01.000Z","level":"INFO","msg":"token refreshed successfully","expiry":"2024-01-15T11:30:01.000Z"}
```

## Next steps

This project is in its early stages and will grow a lot in the coming period.

Here are the highest priority tasks:

- document the scraping tool and update it to work with the new oauth flow
- generate more endpoints with the scraping tool
- add tests to endpoints and client generation
- parse csv endpoints and add types
- automatically refresh tokens

## Scraping Tool

The scraping tool in `tools/scraper` is used to automatically generate Go client code for the iRacing API. It works by discovering available endpoints, fetching sample responses, and generating strongly-typed Go structs and API methods.

### Overview

The scraper performs the following steps:

1. **Authentication**: Connects to the iRacing API using your credentials
2. **Endpoint Discovery**: Fetches the list of all available API endpoints
3. **Sample Collection**: Retrieves sample responses from each endpoint
4. **Type Generation**: Uses quicktype to generate Go structs from JSON responses
5. **Code Generation**: Creates API client methods using Jinja2 templates
6. **Formatting**: Runs `go fmt` to format the generated code

### Prerequisites

Before running the scraper, you need:

- **Python 3.x** with pip
- **Node.js** with npm/npx (for quicktype)
- **Go** (for code formatting)
- **iRacing Account** with valid credentials

### Setup

1. Navigate to the scraper directory:
   ```bash
   cd tools/scraper
   ```

2. Install Python dependencies:
   ```bash
   pip install -r requirements.txt
   ```

3. Install Node.js dependencies:
   ```bash
   npm install
   ```

4. Create a `.env` file with your iRacing credentials:
   ```
   IRACING_EMAIL=your_email@example.com
   IRACING_PASSWORD=your_password
   ```

### Usage

Run the scraper with:

```bash
python main.py
```

The tool will:
- Fetch the API documentation from iRacing
- Download sample responses for each endpoint
- Generate Go type definitions in `../../api/`
- Create API client methods for each category

### Configuration

The scraper can be customized through the `data.py` file, which contains the `OVERRIDES` dictionary. This allows you to:

- Specify sample parameters for endpoints that require them
- Set response format (json/csv)
- Control S3 caching behavior
- Provide test data for endpoints with required parameters

Example override:
```python
"league": {
    "get": {
        "params": [
            {
                "league_id": 4403,
            }
        ],
    },
}
```

### Output Structure

The scraper generates files in the following structure:

```
api/
├── category_name/
│   ├── main.go              # Category API with all endpoint methods
│   └── endpoint_name/
│       └── structs.go       # Parameter and response structs
```

### Architecture

The scraper consists of several modules:

- **main.py**: Entry point and workflow orchestration
- **api_client.py**: Handles authentication and API requests
- **endpoints_documentation.py**: Fetches endpoint metadata from iRacing
- **endpoints_parsing.py**: Represents endpoints and fetches sample responses
- **templating.py**: Generates Go code from Jinja2 templates
- **format.py**: Formatting utilities and Go code formatting
- **data.py**: Configuration overrides for specific endpoints
- **utils/quicktype.py**: Wrapper for the quicktype code generator

### Troubleshooting

- **Authentication Errors**: Verify your credentials in the `.env` file
- **Missing Dependencies**: Ensure all Python and Node.js packages are installed
- **API Rate Limiting**: The tool uses parallel workers; reduce the `workers` parameter if you encounter rate limits
- **Endpoint Failures**: Some endpoints require specific parameters; add them to `OVERRIDES` in `data.py`

