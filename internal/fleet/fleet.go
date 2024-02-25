package fleet

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/dpnetca/gostSDK/internal/client"
)

type FleetClient struct {
	client *client.STClient
}

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

func NewFleetClient(client *client.STClient) *FleetClient {
	return &FleetClient{client: client}
}

type ListShipsResponse struct {
	Data []Ship `json:"data"`
	meta Meta   `jston:"meta"`
}

func (f *FleetClient) ListShips() ([]Ship, error) {
	var ships []Ship
	page:=1
	limit:=20
	for {
		endpoint := fmt.Sprintf("/my/ships?page=%d&limit=%d", page, limit)
		url := f.client.Base_url + endpoint
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}

		res, err := f.client.SendWithAuth(req)
		if err != nil {
			return nil, err
		}

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		if res.StatusCode > 299 {
			return nil, fmt.Errorf("%s", data)
		}

		var response ListShipsResponse
		if err = json.Unmarshal(data, &response); err != nil {
			return nil, err
		}

		ships = append(ships, response.Data...)

		page ++
		if page >=1 {
			break
		}
	}

	return ships, nil

}

type GetShipResponse struct {
	Data Ship `json:"data"`
}

func (f *FleetClient) GetShip(symbol string) (*Ship, error) {
	url := f.client.Base_url + "/my/ships/" + symbol
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, err := f.client.SendWithAuth(req)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("%s", data)
	}

	var response GetShipResponse
	if err = json.Unmarshal(data, &response); err != nil {
		return nil, err
	}
	return &response.Data, nil
}
