import os
import logging

from format import to_camel_case
from templating import jinja2_environment
from endpoints_parsing import Endpoint


def run(endpoints: list[Endpoint]):
    logging.info("Generating category APIs...")

    # Get unique categories
    categories = set([endpoint.category for endpoint in endpoints])

    # Prepare the templates
    category_template = jinja2_environment.get_template("category.go")

    for category in categories:
        # Generate the API code for each category
        api_code = category_template.render(
            package_name=category, api_name=to_camel_case(category) + "Api"
        )
        os.makedirs(f"../api/{category}", exist_ok=True)
        with open(f"../api/{category}/main.go", "w") as f:
            f.write(api_code)

    logging.info("Category APIs generated successfully.")
