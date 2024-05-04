package agents

import (
	"fmt"
	"io"
	"net/http"

	"github.com/dpnetca/gost/pkg/gostSDK/internal/client"
)

type Client struct {
	client *client.Client
}


func New(client *client.Client) *Client {
	return &Client{client: client}
}

func (c Client) sendGetRequest(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.client.SendWithAuth(req)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("%s", data)
	}
	return data, nil

}
