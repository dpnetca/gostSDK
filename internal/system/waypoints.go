package system

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dpnetca/gostSDK/internal/meta"
)

type Orbitals struct {
	Symbol string `json:"symbol"`
}
type Faction struct {
	Symbol string `json:"symbol"`
}
type Traits struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type Modifiers struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type Chart struct {
	WaypointSymbol string    `json:"waypointSymbol"`
	SubmittedBy    string    `json:"submittedBy"`
	SubmittedOn    time.Time `json:"submittedOn"`
}
type Waypoint struct {
	Symbol              string      `json:"symbol"`
	Type                string      `json:"type"`
	SystemSymbol        string      `json:"systemSymbol"`
	X                   int         `json:"x"`
	Y                   int         `json:"y"`
	Orbitals            []Orbitals  `json:"orbitals"`
	Orbits              string      `json:"orbits"`
	Faction             Faction     `json:"faction"`
	Traits              []Traits    `json:"traits"`
	Modifiers           []Modifiers `json:"modifiers"`
	Chart               Chart       `json:"chart"`
	IsUnderConstruction bool        `json:"isUnderConstruction"`
}

type ListWaypointInSystemResponse struct {
	Data []Waypoint `json:"data"`
	Meta meta.Meta  `json:"meta"`
}

// This will return automatically page through list and return all waypoints in system
// to get a single page of ships use ListWaypointsInSystemByPage
func (c *Client) ListWaypointsInSystem(systemSymbol string) ([]Waypoint, error) {
	var waypoints []Waypoint
	page := 1
	limit := 20
	for {
		response, err := c.ListWaypointsInSystemByPage(systemSymbol, page, limit)
		if err != nil {
			return nil, err
		}
		waypoints = append(waypoints, response.Data...)

		if page*limit >= response.Meta.Total {
			break
		}
		page++
	}

	return waypoints, nil
}

func (c *Client) ListWaypointsInSystemByPage(systemSymbol string, page, limit int) (ListWaypointInSystemResponse, error) {
	endpoint := fmt.Sprintf("/systems/%s/waypoints?page=%d&limit=%d", systemSymbol, page, limit)
	url := c.client.Base_url + endpoint
	data, err := c.sendGetRequest(url)
	if err != nil {
		return ListWaypointInSystemResponse{}, err
	}

	var response ListWaypointInSystemResponse
	if err = json.Unmarshal(data, &response); err != nil {
		return ListWaypointInSystemResponse{}, err
	}
	return response, nil
}
