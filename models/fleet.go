package models

import "time"

type Ship struct {
	Symbol       string           `json:"symbol"`
	Registration ShipRegistration `json:"registration"`
	Nav          ShipNav          `json:"nav"`
	Crew         ShipCrew         `json:"crew"`
	Frame        ShipFrame        `json:"frame"`
	Reactor      ShipReactor      `json:"reactor"`
	Engine       ShipEngine       `json:"engine"`
	Cooldown     ShipCooldown     `json:"cooldown"`
	Modules      []ShipModules    `json:"modules"`
	Mounts       []ShipMounts     `json:"mounts"`
	Cargo        ShipCargo        `json:"cargo"`
	Fuel         ShipFuel         `json:"fuel"`
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
type ShipRequirements struct {
	Power int `json:"power"`
	Crew  int `json:"crew"`
	Slots int `json:"slots"`
}
type ShipFrame struct {
	Symbol         string           `json:"symbol"`
	Name           string           `json:"name"`
	Description    string           `json:"description"`
	Condition      int              `json:"condition"`
	Integrity      int              `json:"integrity"`
	ModuleSlots    int              `json:"moduleSlots"`
	MountingPoints int              `json:"mountingPoints"`
	FuelCapacity   int              `json:"fuelCapacity"`
	Requirements   ShipRequirements `json:"requirements"`
}
type ShipReactor struct {
	Symbol       string       `json:"symbol"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	Condition    int          `json:"condition"`
	Integrity    int          `json:"integrity"`
	PowerOutput  int          `json:"powerOutput"`
	Requirements ShipRequirements `json:"requirements"`
}
type ShipEngine struct {
	Symbol       string       `json:"symbol"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	Condition    int          `json:"condition"`
	Integrity    int          `json:"integrity"`
	Speed        int          `json:"speed"`
	Requirements ShipRequirements `json:"requirements"`
}
type ShipCooldown struct {
	ShipSymbol       string    `json:"shipSymbol"`
	TotalSeconds     int       `json:"totalSeconds"`
	RemainingSeconds int       `json:"remainingSeconds"`
	Expiration       time.Time `json:"expiration"`
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
type ShipFuel struct {
	Current  int `json:"current"`
	Capacity int `json:"capacity"`
	Consumed struct {
		Amount    int       `json:"amount"`
		Timestamp time.Time `json:"timestamp"`
	} `json:"consumed"`
}
type ShipModules struct {
	Symbol       string       `json:"symbol"`
	Capacity     int          `json:"capacity"`
	Range        int          `json:"range"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	Requirements ShipRequirements `json:"requirements"`
}
type ShipMounts struct {
	Symbol       string       `json:"symbol"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	Strength     int          `json:"strength"`
	Deposits     []string     `json:"deposits"`
	Requirements ShipRequirements `json:"requirements"`
}
type ShipConditionEvent struct {
	Symbol      string `json:"symbol"`
	Component   string `json:"component"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
