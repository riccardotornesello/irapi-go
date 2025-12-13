package session

import (
	"github.com/riccardotornesello/irapi-go/client"
)

type SessionApi struct {
	client *client.ApiClient
}

func NewSessionApi(client *client.ApiClient) *SessionApi {
	return &SessionApi{
		client: client,
	}
}
