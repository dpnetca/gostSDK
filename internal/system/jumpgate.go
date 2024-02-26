package system

import (
	"encoding/json"
	"fmt"

	"github.com/dpnetca/gostSDK/internal/meta"
)

type Jumpgate struct {
	Symbol      string   `json:"symbol"`
	Connections []string `json:"connections"`
}

type GetJumpgateResponse struct {
	Data Jumpgate  `json:"data"`
	Meta meta.Meta `json:"meta"`
}

func (c *Client) GetJumpgate(system, waypoint string) (Jumpgate, error) {
	endpoint := fmt.Sprintf("/systems/%s/waypoints/%s/jump-gate", system, waypoint)
	url := c.client.Base_url + endpoint

	data, err := c.sendGetRequest(url)
	if err != nil {
		return Jumpgate{}, err
	}

	var response GetJumpgateResponse
	if err = json.Unmarshal(data, &response); err != nil {
		return Jumpgate{}, err
	}
	return response.Data, nil
}
