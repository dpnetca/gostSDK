package agents

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/dpnetca/gostSDK/internal/client"
)

type AgentClient struct {
	client *client.STClient
}

type AgentReponse struct {
	Agent Agent `json:"data"`
}
type Agent struct {
	AccountId       string `json:"accountId,omitempty"`
	Symbol          string `json:"symbol,omitempty"`
	Headquarters    string `json:"headquarters,omitempty"`
	Credits         int64  `json:"credits,omitempty"`
	StartingFaction string `json:"startingFaction,omitempty"`
	ShipCount       int32  `json:"shipCount,omitempty"`
}

func NewAgentClient(client *client.STClient) *AgentClient {
	return &AgentClient{client: client}
}

func (a *AgentClient) GetAgent() (*Agent, error) {
	url := a.client.Base_url + "/my/agent"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, err := a.client.SendWithAuth(req)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("%s", data)
	}

	var agentResponse AgentReponse
	if err = json.Unmarshal(data, &agentResponse); err != nil {
		return nil, err
	}
	return &agentResponse.Agent, nil
}
