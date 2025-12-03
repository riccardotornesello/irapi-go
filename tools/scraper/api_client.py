"""
This module provides authentication and API communication for the iRacing API.

The APIClient class handles:
- Authentication using email/password credentials
- Making authenticated requests to API endpoints
- Handling S3-cached responses
"""

import requests
import hashlib
import base64


class APIClient:
    """
    Client for interacting with the iRacing API.

    This client handles authentication and provides methods to call API endpoints.
    It supports both direct API responses and S3-cached responses.

    Attributes:
        api_client (requests.Session): Authenticated session for making API requests.
    """

    access_token: str

    def __init__(
        self, email: str, password: str, client_id: str, client_secret: str
    ) -> None:
        """
        Initialize the API client and authenticate with iRacing.

        Args:
            email (str): iRacing account email address.
            password (str): iRacing account password.
            client_id (str): iRacing API client ID.
            client_secret (str): iRacing API client secret.

        Raises:
            Exception: If authentication fails or credentials are invalid.
        """
        body = {
            "grant_type": "password_limited",
            "client_id": client_id,
            "client_secret": _mask(client_secret, client_id),
            "username": email,
            "password": _mask(password, email),
            "scope": "iracing.auth",
        }

        response = requests.post("https://oauth.iracing.com/oauth2/token", data=body)
        response.raise_for_status()

        self.access_token = response.json().get("access_token")

    def call(self, url: str, method="GET", params: dict | None = None):
        return requests.request(
            method,
            url,
            headers={"Authorization": f"Bearer {self.access_token}"},
            params=params,
        )

    def call_endpoint(
        self, url: str, params: dict | None = None, s3_cache: bool = True
    ) -> str:
        """
        Call an API endpoint and return the response text.

        Args:
            url (str): The API endpoint URL to call.
            params (dict, optional): Query parameters for the request. Defaults to None.
            s3_cache (bool, optional): If True, follows S3 link from response.
                                      If False, returns direct response. Defaults to True.

        Returns:
            str: The response text from the API or S3 bucket.

        Raises:
            Exception: If the API call fails or returns a non-200 status code.
        """
        response = self.call(url, params=params)
        if response.status_code != 200:
            raise Exception(f"API call failed: {response.text}")

        if not s3_cache:
            return response.text

        s3_url = response.json()["link"]
        response = requests.get(s3_url)

        return response.text


def _mask(secret: str, id: str) -> str:
    """Mask a secret using a given ID."""

    token = hashlib.sha256(f"{secret}{id}".encode()).digest()
    base64_token = base64.b64encode(token).decode()
    return base64_token
