package system

import (
	"encoding/json"
	"fmt"

	"github.com/dpnetca/gostSDK/models"
)

type GetShipyardResponse struct {
	Data models.Shipyard `json:"data"`
	Meta models.Meta     `json:"meta"`
}

func (c *Client) GetShipyard(system, waypoint string) (models.Shipyard, error) {
	endpoint := fmt.Sprintf("/systems/%s/waypoints/%s/shipyard", system, waypoint)
	url := c.client.Base_url + endpoint

	data, err := c.sendGetRequest(url)
	if err != nil {
		return models.Shipyard{}, err
	}

	var response GetShipyardResponse
	if err = json.Unmarshal(data, &response); err != nil {
		return models.Shipyard{}, err
	}
	return response.Data, nil
}
