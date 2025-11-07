package hosted

import (
	"github.com/riccardotornesello/irapi-go/api/hosted/combined_sessions"
	"github.com/riccardotornesello/irapi-go/api/hosted/sessions"
	"github.com/riccardotornesello/irapi-go/client"
)

type HostedApi struct {
	Client *client.ApiClient
}

func (api *HostedApi) CombinedSessions(parameters *combined_sessions.HostedCombinedSessionsParams) (*combined_sessions.HostedCombinedSessionsResponse, error) {
	return client.GetJson[combined_sessions.HostedCombinedSessionsResponse](api.Client, "/data/hosted/combined_sessions", parameters)
}

func (api *HostedApi) Sessions() (*sessions.HostedSessionsResponse, error) {
	return client.GetJson[sessions.HostedSessionsResponse](api.Client, "/data/hosted/sessions", nil)
}
