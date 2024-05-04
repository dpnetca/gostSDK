package agents

import (
	"encoding/json"
	"fmt"

	"github.com/dpnetca/gost/pkg/gostSDK/models"
)

type AgentReponse struct {
	Data models.Agent `json:"data"`
}

func (c *Client) GetAgent() (models.Agent, error) {
	url := c.client.Base_url + "/my/agent"
	data, err := c.sendGetRequest(url)
	if err != nil {
		return models.Agent{}, err
	}

	var agentResponse AgentReponse
	if err = json.Unmarshal(data, &agentResponse); err != nil {
		return models.Agent{}, err
	}
	return agentResponse.Data, nil
}

func (c *Client) GetPublicAgent(agent string) (models.Agent, error) {
	url := c.client.Base_url + "/agents/" + agent
	data, err := c.sendGetRequest(url)
	if err != nil {
		return models.Agent{}, err
	}

	var agentResponse AgentReponse
	if err = json.Unmarshal(data, &agentResponse); err != nil {
		return models.Agent{}, err
	}
	return agentResponse.Data, nil
}

type ListAgentsResponse struct {
	Data []models.Agent     `json:"data"`
	Meta models.Meta `json:"meta"`
}

func (c *Client) ListAgents() ([]models.Agent, error) {

	var agents []models.Agent
	page := 1
	limit := 20
	for {
		response, err := c.ListAgentsByPage(page, limit)
		if err != nil {
			return nil, err
		}
		agents = append(agents, response.Data...)

		if page*limit >= response.Meta.Total {
			break
		}
		page++
	}

	return agents, nil
}

func (c *Client) ListAgentsByPage(page, limit int) (ListAgentsResponse, error) {
	endpoint := fmt.Sprintf("/agents?page=%d&limit=%d", page, limit)
	url := c.client.Base_url + endpoint
	data, err := c.sendGetRequest(url)
	if err != nil {
		return ListAgentsResponse{}, err
	}

	var response ListAgentsResponse
	if err = json.Unmarshal(data, &response); err != nil {
		return ListAgentsResponse{}, err
	}
	return response, nil
}
