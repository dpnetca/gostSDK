package fleet

import (
	"encoding/json"
	"fmt"
	"time"
)

type Ship struct {
	Symbol       string           `json:"symbol"`
	Registration ShipRegistration `json:"registration"`
	Nav          ShipNav          `json:"nav"`
	Crew         ShipCrew         `json:"crew"`
}
type ShipRegistration struct {
	Name          string `json:"name"`
	FactionSymbol string `json:"factionSymbol"`
	Role          string `json:"role"`
}
type ShipNavRouteWaypoint struct {
	Symbol       string `json:"symbol"`
	Type         string `json:"type"`
	SystemSymbol string `json:"systemSymbol"`
	X            int    `json:"x"`
	Y            int    `json:"y"`
}
type ShipNavRoute struct {
	Destination   ShipNavRouteWaypoint `json:"destination"`
	Origin        ShipNavRouteWaypoint `json:"origin"`
	DepartureTime time.Time            `json:"departureTime"`
	Arrival       time.Time            `json:"arrival"`
}
type ShipNav struct {
	SystemSymbol   string       `json:"systemSymbol"`
	WaypointSymbol string       `json:"waypointSymbol"`
	Route          ShipNavRoute `json:"route"`
	Status         string       `json:"status"`
	FlightMode     string       `json:"flightMode"`
}
type ShipCrew struct {
	Current  int    `json:"current"`
	Required int    `json:"required"`
	Capacity int    `json:"capacity"`
	Rotation string `json:"rotation"`
	Morale   int    `json:"morale"`
	Wages    int    `json:"wages"`
}

type Meta struct {
	Total int `json:"total"`
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type ListShipsResponse struct {
	Data []Ship `json:"data"`
	Meta Meta   `jston:"meta"`
}

// This will return automatically page through list and return all ships
// to get a single page of ships use ListShipsByPage
func (c *Client) ListShips() ([]Ship, error) {
	var ships []Ship
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

type GetShipResponse struct {
	Data Ship `json:"data"`
}

func (c *Client) GetShip(symbol string) (*Ship, error) {
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
