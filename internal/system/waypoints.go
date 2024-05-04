package system

import (
	"encoding/json"
	"fmt"

	"github.com/dpnetca/gostSDK/models"
)

type ListWaypointInSystemResponse struct {
	Data []models.Waypoint `json:"data"`
	Meta models.Meta       `json:"meta"`
}

// This will return automatically page through list and return all waypoints in system
// to get a single page of ships use ListWaypointsInSystemByPage
func (c *Client) ListWaypointsInSystem(systemSymbol string) ([]models.Waypoint, error) {
	var waypoints []models.Waypoint
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
