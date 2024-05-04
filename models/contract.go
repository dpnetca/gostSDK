package models

import "time"

type Payment struct {
	OnAccepted  int `json:"onAccepted"`
	OnFulfilled int `json:"onFulfilled"`
}
type Deliver struct {
	TradeSymbol       string `json:"tradeSymbol"`
	DestinationSymbol string `json:"destinationSymbol"`
	UnitsRequired     int    `json:"unitsRequired"`
	UnitsFulfilled    int    `json:"unitsFulfilled"`
}
type Terms struct {
	Deadline time.Time `json:"deadline"`
	Payment  Payment   `json:"payment"`
	Deliver  []Deliver `json:"deliver"`
}
type Contract struct {
	ID               string    `json:"id"`
	FactionSymbol    string    `json:"factionSymbol"`
	Type             string    `json:"type"`
	Terms            Terms     `json:"terms"`
	Accepted         bool      `json:"accepted"`
	Fulfilled        bool      `json:"fulfilled"`
	Expiration       time.Time `json:"expiration"`
	DeadlineToAccept time.Time `json:"deadlineToAccept"`
}
