package src

import (
	"bytes"
	"net/http"
	"time"
)

type APIClient struct {
	username          string
	apiKey            string
	apiServerHostname string
	client            *http.Client
}

func (c *APIClient) InitClient(hostname string, username string, apiKey string) {
	c.apiServerHostname = hostname
	c.username = username
	c.apiKey = apiKey
	c.client = &http.Client{
		Timeout:   time.Second * 60,
		Transport: &http.Transport{},
	}
}

func (c *APIClient) MakeRequest(url string, method string, data []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.username, c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
