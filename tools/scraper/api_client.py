"""
This module contains the function to authenticate with the iRacing API.
"""

import requests
import hashlib
import base64


class APIClient:
    def __init__(self, email, password):
        token = hashlib.sha256(f"{password}{email}".encode()).digest()
        base64_token = base64.b64encode(token).decode()

        self.api_client = requests.Session()

        response = self.api_client.post(
            "https://members-ng.iracing.com/auth",
            json={"email": email, "password": base64_token},
        )

        if response.status_code != 200:
            raise Exception(f"Failed to authenticate: {response.text}")

        auth_code = response.json().get("authcode", 0)
        if auth_code == 0:
            raise Exception(f"Failed to authenticate: {response.text}")

    def call_endpoint(self, url, params=None, s3_cache=True) -> str:
        response = self.api_client.get(url, params=params)
        if response.status_code != 200:
            raise Exception(f"API call failed: {response.text}")

        if not s3_cache:
            return response.text

        s3_url = response.json()["link"]
        response = self.api_client.get(s3_url)

        return response.text
