package system

import (
	"encoding/json"
	"fmt"

	"github.com/dpnetca/gostSDK/models"
)

	type GetJumpgateResponse struct {
	Data models.Jumpgate  `json:"data"`
	Meta models.Meta `json:"meta"`
}

func (c *Client) GetJumpgate(system, waypoint string) (models.Jumpgate, error) {
	endpoint := fmt.Sprintf("/systems/%s/waypoints/%s/jump-gate", system, waypoint)
	url := c.client.Base_url + endpoint

	data, err := c.sendGetRequest(url)
	if err != nil {
		return models.Jumpgate{}, err
	}

	var response GetJumpgateResponse
	if err = json.Unmarshal(data, &response); err != nil {
		return models.Jumpgate{}, err
	}
	return response.Data, nil
}
