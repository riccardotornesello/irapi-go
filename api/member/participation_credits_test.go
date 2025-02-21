package member

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"riccardotornesello.it/sharedtelemetry/iracing/irapi/client"
)

func TestGetMemberParticipationCredits(t *testing.T) {
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

	api := &MemberApi{
		Client: apiClient,
	}

	_, err = api.GetMemberParticipationCredits()
	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestGetMemberParticipationCredits passed")
}
