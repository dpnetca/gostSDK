package client

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

type STClient struct {
	Base_url          string
	authorization     string
	Client            *http.Client
	NormalRateLimiter *rate.Limiter
	BurstRateLimiter  *rate.Limiter
}

func (c *STClient) Send(req *http.Request) (*http.Response, error) {
	ctx := context.Background()
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if c.NormalRateLimiter.Allow() {
	} else if c.BurstRateLimiter.Allow() {
	} else {
		err := c.NormalRateLimiter.Wait(ctx)
		if err != nil {
			return nil, err
		}
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *STClient) SendWithAuth(req *http.Request) (*http.Response, error) {

	req.Header.Add("Authorization", c.authorization)
	res, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *STClient) UpdateAuth(token string) {
	c.authorization = fmt.Sprintf("Bearer %s", token)
}

func NewClient(token string) *STClient {
	base_url := "https://api.spacetraders.io/v2"
	rateLimit := rate.NewLimiter(rate.Every(500*time.Millisecond), 2)
	burstRateLimit := rate.NewLimiter(rate.Every(2*time.Second), 30)
	authorization := fmt.Sprintf("Bearer %s", token)

	c := &STClient{
		Base_url:          base_url,
		authorization:     authorization,
		Client:            http.DefaultClient,
		NormalRateLimiter: rateLimit,
		BurstRateLimiter:  burstRateLimit,
	}
	return c
}
