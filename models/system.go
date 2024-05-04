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
