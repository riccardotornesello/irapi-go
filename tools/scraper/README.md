# iRacing API Scraper

This tool automatically generates Go client code for the iRacing API by discovering endpoints, fetching sample responses, and generating strongly-typed Go structs and API methods.

## Quick Start

1. **Install dependencies:**
   ```bash
   pip install -r requirements.txt
   npm install
   ```

2. **Configure credentials:**
   Create a `.env` file:
   ```
   IRACING_EMAIL=your_email@example.com
   IRACING_PASSWORD=your_password
   ```

3. **Run the scraper:**
   ```bash
   python main.py
   ```

## How It Works

The scraper performs these steps:

1. **Authentication** - Connects to iRacing API using credentials
2. **Discovery** - Fetches list of all available endpoints
3. **Sampling** - Retrieves sample responses from each endpoint
4. **Type Generation** - Uses quicktype to generate Go structs from JSON
5. **Code Generation** - Creates API methods using Jinja2 templates
6. **Test Generation** - Generates test suites for each endpoint
7. **Formatting** - Runs `go fmt` on generated code

## Project Structure

```
tools/scraper/
├── main.py                      # Entry point
├── api_client.py                # Authentication and API requests
├── endpoints_documentation.py   # Endpoint discovery
├── endpoints_parsing.py         # Endpoint representation and sampling
├── templating.py                # Code generation from templates
├── format.py                    # Formatting utilities
├── data.py                      # Configuration overrides
├── utils/
│   └── quicktype.py            # Quicktype wrapper
├── templates/
│   ├── category.j2             # Category API template
│   ├── endpoint_call.j2        # Endpoint method template
│   ├── endpoint_structs.j2     # Struct definitions template
│   └── endpoint_test.j2        # Test suite template
├── requirements.txt            # Python dependencies
└── package.json                # Node.js dependencies (quicktype)
```

## Configuration

Edit `data.py` to customize endpoint behavior:

```python
OVERRIDES = {
    "category_name": {
        "endpoint_name": {
            "params": [{"param_name": value}],  # Sample parameters
            "s3_cache": False,                  # Skip S3 link following
            "format": "csv"                     # Response format
        }
    }
}
```

## Output

The scraper generates files in `../../api/`:

```
api/
├── category_name/
│   ├── main.go              # Category API with all methods
│   └── endpoint_name/
│       ├── structs.go       # Request/response structs
│       └── structs_test.go  # Test suite for structs
```

## Module Documentation

### main.py
Entry point that orchestrates the complete workflow.

### api_client.py
Handles authentication and provides methods to call API endpoints with support for S3-cached responses.

### endpoints_documentation.py
Fetches the list of available endpoints from the iRacing API documentation endpoint.

### endpoints_parsing.py
Contains `Endpoint` and `EndpointParameter` classes that represent API endpoints and manage sample response fetching and type generation.

### templating.py
Generates Go code using Jinja2 templates for category APIs and endpoint structs.

### format.py
Provides string formatting utilities and runs `go fmt` on generated code.

### data.py
Contains the `OVERRIDES` dictionary for customizing endpoint behavior.

### utils/quicktype.py
Wrapper for the quicktype tool that generates Go types from JSON samples.

## Troubleshooting

### Authentication Fails
- Verify credentials in `.env` file
- Check that your iRacing account is active

### Missing Dependencies
```bash
pip install -r requirements.txt
npm install
```

### API Rate Limiting
Reduce the number of workers in `main.py`:
```python
fetch_sample_responses(endpoints, skip_cached=False, workers=1)
```

### Endpoint Requires Parameters
Add sample parameters to `OVERRIDES` in `data.py`:
```python
"category_name": {
    "endpoint_name": {
        "params": [{"required_param": sample_value}]
    }
}
```

## Development

To add support for a new endpoint that requires parameters:

1. Run the scraper once to discover the endpoint
2. Check the error logs to see which parameters are required
3. Add sample parameters to `OVERRIDES` in `data.py`
4. Run the scraper again

## More Information

See the main project README for additional details about the scraping tool and its usage.
