package contracts

import (
	"encoding/json"
	"fmt"

	"github.com/dpnetca/gostSDK/models"
)


type ListContractsResponse struct {
	Data []models.Contract  `json:"data"`
	Meta models.Meta `json:"meta"`
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

func (c *Client) GetContract(contract string) (models.Contract, error) {
	endpoint := fmt.Sprintf("/my/contracts/%s", contract)
	url := c.client.Base_url + endpoint
	data, err := c.sendGetRequest(url)
	if err != nil {
		return models.Contract{}, err
	}

	var response models.Contract
	if err = json.Unmarshal(data, &response); err != nil {
		return models.Contract{}, err
	}
	return response, nil
}
