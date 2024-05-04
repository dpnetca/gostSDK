package fleet

import (
	"encoding/json"
	"fmt"

	"github.com/dpnetca/gostSDK/models"
)

type DockShipResponse struct {
	Data struct {
		Nav models.ShipNav `json:"nav"`
	} `json:"data"`
}

func (c *Client) DockShip(symbol string) (models.ShipNav, error) {
	url := c.client.Base_url + "/my/ships/" + symbol + "/dock"
	data, err := c.sendPostRequest(url, nil)
	if err != nil {
		return models.ShipNav{}, err
	}

	var response DockShipResponse
	if err = json.Unmarshal(data, &response); err != nil {
		return models.ShipNav{}, err
	}
	return response.Data.Nav, nil

}

type GetShipResponse struct {
	Data models.Ship `json:"data"`
}

func (c *Client) GetShip(symbol string) (*models.Ship, error) {
	url := c.client.Base_url + "/my/ships/" + symbol
	data, err := c.sendGetRequest(url)
	if err != nil {
		return nil, err
	}
	var response GetShipResponse
	if err = json.Unmarshal(data, &response); err != nil {
		return nil, err
	}
	return &response.Data, nil
}

type ListShipsResponse struct {
	Data []models.Ship `json:"data"`
	Meta models.Meta   `jston:"meta"`
}

// This will return automatically page through list and return all ships
// to get a single page of ships use ListShipsByPage
func (c *Client) ListShips() ([]models.Ship, error) {
	var ships []models.Ship
	page := 1
	limit := 20
	for {
		response, err := c.ListShipsByPage(page, limit)
		if err != nil {
			return nil, err
		}
		ships = append(ships, response.Data...)

		if page*limit >= response.Meta.Total {
			break
		}
		page++
	}

	return ships, nil
}

func (c *Client) ListShipsByPage(page int, limit int) (ListShipsResponse, error) {
	endpoint := fmt.Sprintf("/my/ships?page=%d&limit=%d", page, limit)
	url := c.client.Base_url + endpoint
	data, err := c.sendGetRequest(url)
	if err != nil {
		return ListShipsResponse{}, err
	}

	var response ListShipsResponse
	if err = json.Unmarshal(data, &response); err != nil {
		return ListShipsResponse{}, err
	}
	return response, nil
}

type OrbitShipResponse struct {
	Data struct {
		Nav models.ShipNav `json:"nav"`
	} `json:"data"`
}

func (c *Client) OrbitShip(symbol string) (models.ShipNav, error) {
	url := c.client.Base_url + "/my/ships/" + symbol + "/orbit"
	data, err := c.sendPostRequest(url, nil)
	if err != nil {
		return models.ShipNav{}, err
	}

	var response OrbitShipResponse
	if err = json.Unmarshal(data, &response); err != nil {
		return models.ShipNav{}, err
	}
	return response.Data.Nav, nil

}
