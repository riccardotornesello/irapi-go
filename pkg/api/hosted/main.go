package hosted

import (
	"github.com/riccardotornesello/irapi-go/pkg/api/hosted/combined_sessions"
	"github.com/riccardotornesello/irapi-go/pkg/api/hosted/sessions"
	"github.com/riccardotornesello/irapi-go/pkg/client"
)

type HostedApi struct {
	client *client.ApiClient
}

func NewHostedApi(client *client.ApiClient) *HostedApi {
	return &HostedApi{
		client: client,
	}
}

func (api *HostedApi) CombinedSessions(parameters *combined_sessions.HostedCombinedSessionsParams) (*combined_sessions.HostedCombinedSessionsResponse, error) {
	resp, err := client.GetJson[combined_sessions.HostedCombinedSessionsResponse](api.client, "/data/hosted/combined_sessions", parameters)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *HostedApi) Sessions() (*sessions.HostedSessionsResponse, error) {
	resp, err := client.GetJson[sessions.HostedSessionsResponse](api.client, "/data/hosted/sessions", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
