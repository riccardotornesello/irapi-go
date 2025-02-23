package testutils

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/riccardotornesello/irapi-go/client"
)

const tokenFile = "../../.token"
const envFile = "../../.env"

func GetApiClient() *client.ApiClient {
	authToken, err := loadToken()

	if err == nil {
		return client.NewApiClientWithToken(authToken)
	} else {
		err := godotenv.Load(envFile)
		if err != nil {
			panic("Error loading .env file")
		}

		email := os.Getenv("IRACING_EMAIL")
		password := os.Getenv("IRACING_PASSWORD")

		apiClient, err := client.NewApiClient(email, password)
		if err != nil {
			panic(err)
		}

		err = saveToken(apiClient.GetAuthToken())
		if err != nil {
			panic(err)
		}

		return apiClient
	}
}

func saveToken(token string) error {
	return os.WriteFile(tokenFile, []byte(token), 0600)
}

func loadToken() (string, error) {
	data, err := os.ReadFile(tokenFile)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
