import os

import jinja2

from format import to_camel_case

jinja2_environment = jinja2.Environment(loader=jinja2.FileSystemLoader("templates/"))


def generate_categories_api(endpoints):
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
