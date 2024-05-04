package models

import "time"

type Jumpgate struct {
	Symbol      string   `json:"symbol"`
	Connections []string `json:"connections"`
}

type Exports struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type Imports struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type Exchange struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type MarketTransactions struct {
	WaypointSymbol string    `json:"waypointSymbol"`
	ShipSymbol     string    `json:"shipSymbol"`
	TradeSymbol    string    `json:"tradeSymbol"`
	Type           string    `json:"type"`
	Units          int       `json:"units"`
	PricePerUnit   int       `json:"pricePerUnit"`
	TotalPrice     int       `json:"totalPrice"`
	Timestamp      time.Time `json:"timestamp"`
}
type TradeGoods struct {
	Symbol        string `json:"symbol"`
	Type          string `json:"type"`
	TradeVolume   int    `json:"tradeVolume"`
	Supply        string `json:"supply"`
	Activity      string `json:"activity"`
	PurchasePrice int    `json:"purchasePrice"`
	SellPrice     int    `json:"sellPrice"`
}
type Market struct {
	Symbol       string               `json:"symbol"`
	Exports      []Exports            `json:"exports"`
	Imports      []Imports            `json:"imports"`
	Exchange     []Exchange           `json:"exchange"`
	Transactions []MarketTransactions `json:"transactions"`
	TradeGoods   []TradeGoods         `json:"tradeGoods"`
}
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
type ShipyardShip struct {
	Type          string        `json:"type"`
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	Supply        string        `json:"supply"`
	Activity      string        `json:"activity"`
	PurchasePrice int           `json:"purchasePrice"`
	Frame         ShipFrame     `json:"frame"`
	Reactor       ShipReactor   `json:"reactor"`
	Engine        ShipEngine    `json:"engine"`
	Modules       []ShipModules `json:"modules"`
	Mounts        []ShipMounts  `json:"mounts"`
	Crew          ShipCrew      `json:"crew"`
}
type Shipyard struct {
	Symbol           string                 `json:"symbol"`
	ShipTypes        []ShipTypes            `json:"shipTypes"`
	Transactions     []ShipyardTransactions `json:"transactions"`
	Ships            []ShipyardShip         `json:"ships"`
	ModificationsFee int                    `json:"modificationsFee"`
}

type System struct {
	Symbol       string     `json:"symbol"`
	SectorSymbol string     `json:"sectorSymbol"`
	X            int        `json:"x"`
	Y            int        `json:"y"`
	Waypoints    []Waypoint `json:"waypoints"`
}

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
