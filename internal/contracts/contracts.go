package contracts

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/dpnetca/gostSDK/models"
)

type AcceptContractResponse struct {
	Data struct {
		Agent    models.Agent    `json:"agent"`
		Contract models.Contract `json:"contact"`
	} `json:"data"`
}

func (c *Client) AcceptContract(contract string) (models.Agent, models.Contract, error) {
	endpoint := fmt.Sprintf("/my/contracts/%s/accept", contract)
	url := c.client.Base_url + endpoint
	data, err := c.sendPostRequest(url, nil)
	if err != nil {
		return models.Agent{}, models.Contract{}, err
	}

	var response AcceptContractResponse
	if err = json.Unmarshal(data, &response); err != nil {
		return models.Agent{}, models.Contract{}, err
	}
	return response.Data.Agent, response.Data.Contract, nil
}

type DeliverToContractResponse struct {
	Data struct {
		Contract models.Contract  `json:"contact"`
		Cargo    models.ShipCargo `json:"cargo"`
	} `json:"data"`
}
type DeliverToContractRequest struct {
	ShipSymbol  string `json:"shipSymbol"`
	TradeSymbol string `json:"tradeSymbol"`
	Units       int    `json:"units"`
}

func (c *Client) DeliverCargoToContract(contract, shipSymbol, tradeSymbol string, units int) (models.Contract, models.ShipCargo, error) {
	endpoint := fmt.Sprintf("/my/contracts/%s/deliver", contract)
	url := c.client.Base_url + endpoint
	reqData := DeliverToContractRequest{
		ShipSymbol:  shipSymbol,
		TradeSymbol: tradeSymbol,
		Units:       units,
	}
	body, err := json.Marshal(reqData)
	if err != nil {
		return models.Contract{}, models.ShipCargo{}, err
	}
	data, err := c.sendPostRequest(url, bytes.NewBuffer(body))
	if err != nil {
		return models.Contract{}, models.ShipCargo{}, err
	}

	var response DeliverToContractResponse
	if err = json.Unmarshal(data, &response); err != nil {
		return models.Contract{}, models.ShipCargo{}, err
	}
	return response.Data.Contract, response.Data.Cargo, nil
}

type FulfillContractResponse struct {
	Data struct {
		Agent    models.Agent    `json:"agent"`
		Contract models.Contract `json:"contact"`
	} `json:"data"`
}

func (c *Client) FulfillContract(contract string) (models.Agent, models.Contract, error) {
	endpoint := fmt.Sprintf("/my/contracts/%s/fulfill", contract)
	url := c.client.Base_url + endpoint
	data, err := c.sendPostRequest(url, nil)
	if err != nil {
		return models.Agent{}, models.Contract{}, err
	}

	var response FulfillContractResponse
	if err = json.Unmarshal(data, &response); err != nil {
		return models.Agent{}, models.Contract{}, err
	}
	return response.Data.Agent, response.Data.Contract, nil
}

type GetContractResponse struct {
	Data models.Contract `json:"data"`
}

func (c *Client) GetContract(contract string) (models.Contract, error) {
	endpoint := fmt.Sprintf("/my/contracts/%s", contract)
	url := c.client.Base_url + endpoint
	data, err := c.sendGetRequest(url)
	if err != nil {
		return models.Contract{}, err
	}

	var response GetContractResponse
	if err = json.Unmarshal(data, &response); err != nil {
		return models.Contract{}, err
	}
	return response.Data, nil
}

type ListContractsResponse struct {
	Data []models.Contract `json:"data"`
	Meta models.Meta       `json:"meta"`
}

func (c *Client) ListContracts() ([]models.Contract, error) {
	var contracts []models.Contract
	page := 1
	limit := 20
	for {
		response, err := c.ListContractsByPage(page, limit)
		if err != nil {
			return nil, err
		}
		contracts = append(contracts, response.Data...)

		if page*limit >= response.Meta.Total {
			break
		}
		page++
	}

	return contracts, nil
}

func (c *Client) ListContractsByPage(page, limit int) (ListContractsResponse, error) {
	endpoint := fmt.Sprintf("/my/contracts?page=%d&limit=%d", page, limit)
	url := c.client.Base_url + endpoint
	data, err := c.sendGetRequest(url)
	if err != nil {
		return ListContractsResponse{}, err
	}

	var response ListContractsResponse
	if err = json.Unmarshal(data, &response); err != nil {
		return ListContractsResponse{}, err
	}
	return response, nil
}
