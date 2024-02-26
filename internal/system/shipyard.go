package system

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dpnetca/gostSDK/internal/meta"
)

type ShipTypes struct {
	Type string `json:"type"`
}
type ShipyardTransactions struct {
	WaypointSymbol string    `json:"waypointSymbol"`
	ShipSymbol     string    `json:"shipSymbol"`
	ShipType       string    `json:"shipType"`
	Price          int       `json:"price"`
	AgentSymbol    string    `json:"agentSymbol"`
	Timestamp      time.Time `json:"timestamp"`
}
type Requirements struct {
	Power int `json:"power"`
	Crew  int `json:"crew"`
	Slots int `json:"slots"`
}
type Frame struct {
	Symbol         string       `json:"symbol"`
	Name           string       `json:"name"`
	Description    string       `json:"description"`
	Condition      int          `json:"condition"`
	ModuleSlots    int          `json:"moduleSlots"`
	MountingPoints int          `json:"mountingPoints"`
	FuelCapacity   int          `json:"fuelCapacity"`
	Requirements   Requirements `json:"requirements"`
}
type Reactor struct {
	Symbol       string       `json:"symbol"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	Condition    int          `json:"condition"`
	PowerOutput  int          `json:"powerOutput"`
	Requirements Requirements `json:"requirements"`
}
type Engine struct {
	Symbol       string       `json:"symbol"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	Condition    int          `json:"condition"`
	Speed        int          `json:"speed"`
	Requirements Requirements `json:"requirements"`
}
type Modules struct {
	Symbol       string       `json:"symbol"`
	Capacity     int          `json:"capacity"`
	Range        int          `json:"range"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	Requirements Requirements `json:"requirements"`
}
type Mounts struct {
	Symbol       string       `json:"symbol"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	Strength     int          `json:"strength"`
	Deposits     []string     `json:"deposits"`
	Requirements Requirements `json:"requirements"`
}
type Crew struct {
	Required int `json:"required"`
	Capacity int `json:"capacity"`
}
type Ships struct {
	Type          string    `json:"type"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Supply        string    `json:"supply"`
	Activity      string    `json:"activity"`
	PurchasePrice int       `json:"purchasePrice"`
	Frame         Frame     `json:"frame"`
	Reactor       Reactor   `json:"reactor"`
	Engine        Engine    `json:"engine"`
	Modules       []Modules `json:"modules"`
	Mounts        []Mounts  `json:"mounts"`
	Crew          Crew      `json:"crew"`
}
type Shipyard struct {
	Symbol           string                 `json:"symbol"`
	ShipTypes        []ShipTypes            `json:"shipTypes"`
	Transactions     []ShipyardTransactions `json:"transactions"`
	Ships            []Ships                `json:"ships"`
	ModificationsFee int                    `json:"modificationsFee"`
}
type GetShipyardResponse struct {
	Data Shipyard  `json:"data"`
	Meta meta.Meta `json:"meta"`
}

func (c *Client) GetShipyard(system, waypoint string) (Shipyard, error) {
	endpoint := fmt.Sprintf("/systems/%s/waypoints/%s/shipyard", system, waypoint)
	url := c.client.Base_url + endpoint

	data, err := c.sendGetRequest(url)
	if err != nil {
		return Shipyard{}, err
	}

	var response GetShipyardResponse
	if err = json.Unmarshal(data, &response); err != nil {
		return Shipyard{}, err
	}
	return response.Data, nil
}
