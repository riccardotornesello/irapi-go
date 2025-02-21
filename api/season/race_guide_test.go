package season

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"riccardotornesello.it/sharedtelemetry/iracing/irapi/client"
)

func TestGetSeasonRaceGuide(t *testing.T) {
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

	api := &SeasonApi{
		Client: apiClient,
	}

	_, err = api.GetSeasonRaceGuide()
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetSeasonRaceGuide passed")
}
