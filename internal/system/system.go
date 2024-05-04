package system

import (
	"encoding/json"
	"fmt"

	"github.com/dpnetca/gostSDK/models"
)
type GetSystemResponse struct {
	Data models.System `json:"data"`
}
type ListSystemsResponse struct {
	Data []models.System  `json:"data"`
	Meta models.Meta `jston:"meta"`
}

func (c *Client) GetSystem(systemSymbol string) (models.System, error) {
	endpoint := fmt.Sprintf("/systems/%s", systemSymbol)
	url := c.client.Base_url + endpoint
	data, err := c.sendGetRequest(url)
	if err != nil {
		return models.System{}, err
	}
	var response GetSystemResponse
	if err = json.Unmarshal(data, &response); err != nil {
		return models.System{}, err
	}
	return response.Data, nil
}

func (c *Client) ListSystems() ([]models.System, error) {
	var systems []models.System
	page := 1
	limit := 20
	for {
		response, err := c.ListSystemsByPage(page, limit)
		if err != nil {
			return nil, err
		}
		systems = append(systems, response.Data...)

		if page*limit >= response.Meta.Total {
			break
		}
		page++
	}

	return systems, nil
}

func (c *Client) ListSystemsByPage(page, limit int) (ListSystemsResponse, error) {
	endpoint := fmt.Sprintf("/systems?page=%d&limit=%d", page, limit)
	url := c.client.Base_url + endpoint
	data, err := c.sendGetRequest(url)
	if err != nil {
		return ListSystemsResponse{}, err
	}

	var response ListSystemsResponse
	if err = json.Unmarshal(data, &response); err != nil {
		return ListSystemsResponse{}, err
	}
	return response, nil
}
