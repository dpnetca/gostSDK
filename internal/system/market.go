package system

import (
	"encoding/json"
	"fmt"

	"github.com/dpnetca/gostSDK/models"
)

type GetMarketResponse struct {
	Data models.Market      `json:"data"`
	Meta models.Meta `json:"meta"`
}

func (c *Client) GetMarket(system, waypoint string) (models.Market, error) {
	endpoint := fmt.Sprintf("/systems/%s/waypoints/%s/market", system, waypoint)
	url := c.client.Base_url + endpoint

	data, err := c.sendGetRequest(url)
	if err != nil {
		return models.Market{}, err
	}

	var response GetMarketResponse
	if err = json.Unmarshal(data, &response); err != nil {
		return models.Market{}, err
	}
	return response.Data, nil
}
