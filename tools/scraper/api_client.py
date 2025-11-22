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
    def __init__(self, email: str, password: str) -> None:
        """
        Initialize the API client and authenticate with iRacing.
        
        Args:
            email (str): iRacing account email address.
            password (str): iRacing account password.
            
        Raises:
            Exception: If authentication fails or credentials are invalid.
        """
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

    def call_endpoint(self, url: str, params: dict | None = None, s3_cache: bool = True) -> str:
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
        response = self.api_client.get(url, params=params)
        if response.status_code != 200:
            raise Exception(f"API call failed: {response.text}")

        if not s3_cache:
            return response.text

        s3_url = response.json()["link"]
        response = self.api_client.get(s3_url)

        return response.text
