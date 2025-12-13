package testutil

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/riccardotornesello/irapi-go/pkg/client"
)

func GetApiClient(t *testing.T) *client.ApiClient {
	t.Helper()

	os.Clearenv()

	err := godotenv.Load()
	if err != nil {
		t.Fatal("Error loading .env file")
	}

	username := os.Getenv("IRACING_USERNAME")
	password := os.Getenv("IRACING_PASSWORD")
	clientId := os.Getenv("IRACING_CLIENT_ID")
	clientSecret := os.Getenv("IRACING_CLIENT_SECRET")

	if username == "" || password == "" || clientId == "" || clientSecret == "" {
		t.Fatal("Missing required environment variables for API client")
	}

	client, err := client.NewPasswordLimitedApiClient(username, password, &client.Options{
		ClientId:     clientId,
		ClientSecret: clientSecret,
	})
	if err != nil {
		t.Fatalf("Error creating API client: %v", err)
	}

	t.Cleanup(func() {
		os.Clearenv()
	})

	return client
}
