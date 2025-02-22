package lookup

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"riccardotornesello.it/sharedtelemetry/iracing/irapi/client"
)

func TestGetLookupClubHistory(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Fatal("Error loading .env file")
	}

	email := os.Getenv("IRACING_EMAIL")
	password := os.Getenv("IRACING_PASSWORD")

	apiClient, err := client.NewApiClient(email, password)
	if err != nil {
		t.Fatal(err)
	}

	api := &LookupApi{
		Client: apiClient,
	}

	_, err = api.GetLookupClubHistory()
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetLookupClubHistory passed")
}
