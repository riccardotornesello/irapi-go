package testutil

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/riccardotornesello/irapi-go/client"
)

func GetApiClient(t *testing.T) *client.ApiClient {
	t.Helper()

	os.Clearenv()

	err := godotenv.Load()
	if err != nil {
		t.Fatal("Error loading .env file")
	}

	client, err := client.NewPasswordLimitedApiClient("", "", "", "") // TODO
	if err != nil {
		t.Fatalf("Error creating API client: %v", err)
	}

	t.Cleanup(func() {
		os.Clearenv()
	})

	return client
}
