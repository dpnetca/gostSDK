package system

import (
	"encoding/json"
	"fmt"

	"github.com/dpnetca/gostSDK/internal/meta"
)

type System struct {
	Symbol       string `json:"symbol"`
	SectorSymbol string `json:"sectorSymbol"`
	X            int    `json:"x"`
	Y            int    `json:"y"`
}

type GetSystemResponse struct {
	Data System    `json:"data"`
	Meta meta.Meta `jston:"meta"`
}

func (c *Client) GetSystem(systemSymbol string) (System, error) {
	endpoint := fmt.Sprintf("/systems/%s", systemSymbol)
	url := c.client.Base_url + endpoint
	data, err := c.sendGetRequest(url)
	if err != nil {
		return System{}, err
	}
	var response GetSystemResponse
	if err = json.Unmarshal(data, &response); err != nil {
		return System{}, err
	}
	return response.Data, nil
}
