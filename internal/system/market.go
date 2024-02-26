package system

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dpnetca/gostSDK/internal/meta"
)

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

type GetMarketResponse struct {
	Data Market    `json:"data"`
	Meta meta.Meta `json:"meta"`
}

func (c *Client) GetMarket(system, waypoint string) (Market, error) {
	endpoint := fmt.Sprintf("/systems/%s/waypoints/%s/market", system, waypoint)
	url := c.client.Base_url + endpoint

	data, err := c.sendGetRequest(url)
	if err != nil {
		return Market{}, err
	}

	var response GetMarketResponse
	if err = json.Unmarshal(data, &response); err != nil {
		return Market{}, err
	}
	return response.Data, nil
}
