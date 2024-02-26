package agents

import (
	"encoding/json"
)

type Agent struct {
	AccountId       string `json:"accountId,omitempty"`
	Symbol          string `json:"symbol,omitempty"`
	Headquarters    string `json:"headquarters,omitempty"`
	Credits         int64  `json:"credits,omitempty"`
	StartingFaction string `json:"startingFaction,omitempty"`
	ShipCount       int32  `json:"shipCount,omitempty"`
}

type AgentReponse struct {
	Agent Agent `json:"data"`
}

func (c *Client) GetAgent() (*Agent, error) {
	url := c.client.Base_url + "/my/agent"
	data, err := c.sendGetRequest(url)
	if err != nil {
		return nil, err
	}

	var agentResponse AgentReponse
	if err = json.Unmarshal(data, &agentResponse); err != nil {
		return nil, err
	}
	return &agentResponse.Agent, nil
}
