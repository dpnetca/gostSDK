package models

import "time"

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
type ShipCargo struct {
	Capacity  int `json:"capacity"`
	Units     int `json:"units"`
	Inventory []struct {
		Symbol      string `json:"symbol"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Units       int    `json:"units"`
	} `json:"inventory"`
}
