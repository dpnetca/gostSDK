package client

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/dpnetca/gost/logger"
	"golang.org/x/time/rate"
)

type Client struct {
	Base_url          string
	authorization     string
	Client            *http.Client
	NormalRateLimiter *rate.Limiter
	BurstRateLimiter  *rate.Limiter
}

var ReSystem = regexp.MustCompile(`(.*)/X1-[\w\d]+/(.*)`)
var ReWaypoint = regexp.MustCompile(`(.*)/X1-[\w\d]+-[\w\d]+/(.*)`)

func New(token string) *Client {
	base_url := "https://api.spacetraders.io/v2"
	rateLimit := rate.NewLimiter(rate.Every(500*time.Millisecond), 1)
	// burstRateLimit := rate.NewLimiter(rate.Every(2000*time.Millisecond), 1)
	authorization := fmt.Sprintf("Bearer %s", token)

	c := &Client{
		Base_url:          base_url,
		authorization:     authorization,
		Client:            http.DefaultClient,
		NormalRateLimiter: rateLimit,
		// BurstRateLimiter:  burstRateLimit,
	}
	return c
}

func (c *Client) Send(req *http.Request) (*http.Response, error) {
	ctx := context.Background()
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if c.NormalRateLimiter.Allow() {
	// } else if c.BurstRateLimiter.Allow() {
	} else {
		err := c.NormalRateLimiter.Wait(ctx)
		if err != nil {
			return nil, err
		}
	}
	sendTime := time.Now()
	resp, err := c.Client.Do(req)
	elapsed := time.Since(sendTime)
	if err != nil {
		return nil, err
	}
	normalizedUrl := ReSystem.ReplaceAllString(fmt.Sprintf("%s", resp.Request.URL), `${1}/X1-./${2}`)
	normalizedUrl = ReWaypoint.ReplaceAllString(normalizedUrl, `${1}/X1-.-./${2}`)

	logger.Debug(
		"http request done",
		"URL", resp.Request.URL,
		"normalizedUrl", normalizedUrl,
		"requestBody", resp.Request.Body,
		"requestMethod", resp.Request.Method,
		"statusCode", resp.StatusCode,
		"responseTime", elapsed.Milliseconds(),
	)
	if resp.StatusCode == 429 {
		logger.Warn("rate limited", "headers", resp.Header)
		// TODO: change sleep...maybe exponetial backoff...
		time.Sleep(1 * time.Second)
		resp, err = c.Send(req)
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}

func (c *Client) SendWithAuth(req *http.Request) (*http.Response, error) {

	req.Header.Add("Authorization", c.authorization)

	res, err := c.Send(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *Client) UpdateAuth(token string) {
	c.authorization = fmt.Sprintf("Bearer %s", token)
}
