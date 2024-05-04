package models

type Agent struct {
	AccountId       string `json:"accountId,omitempty"`
	Symbol          string `json:"symbol,omitempty"`
	Headquarters    string `json:"headquarters,omitempty"`
	Credits         int64  `json:"credits,omitempty"`
	StartingFaction string `json:"startingFaction,omitempty"`
	ShipCount       int32  `json:"shipCount,omitempty"`
}
