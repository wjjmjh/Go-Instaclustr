package src

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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

func (c *APIClient) RetrieveActiveClusters() (*[]Cluster, error) {
	url := fmt.Sprintf("%s/provisioning/v1/", c.apiServerHostname)
	resp, err := c.MakeRequest(url, "GET", nil)
	if err != nil {
		return nil, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Status code: %d, message: %s", resp.StatusCode, bodyText))
	}
	var clusters []Cluster
	json.Unmarshal(bodyText, &clusters)
	return &clusters, nil
}

func (c *APIClient) ReadCluster(clusterID string) (*Cluster, error) {
	url := fmt.Sprintf("%s/provisioning/v1/%s", c.apiServerHostname, clusterID)
	resp, err := c.MakeRequest(url, "GET", nil)
	if err != nil {
		return nil, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 202 {
		return nil, errors.New(fmt.Sprintf("Status code: %d, message: %s", resp.StatusCode, bodyText))
	}
	var cluster Cluster
	json.Unmarshal(bodyText, &cluster)
	return &cluster, nil
}

func (c *APIClient) DeleteCluster(clusterID string) error {
	url := fmt.Sprintf("%s/provisioning/v1/%s", c.apiServerHostname, clusterID)
	resp, err := c.MakeRequest(url, "DELETE", nil)
	if err != nil {
		return err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 202 {
		return errors.New(fmt.Sprintf("Status code: %d, message: %s", resp.StatusCode, bodyText))
	}
	return nil
}

func (c *APIClient) DeleteActivateClusters() error {
	clusters, err := c.RetrieveActiveClusters()
	if err != nil {
		return err
	}
	for _, cluster := range *clusters {
		c.DeleteCluster(cluster.ID)
	}
	return nil
}

func (c *APIClient) CreateCluster(data []byte) (string, error) {
	url := fmt.Sprintf("%s/provisioning/v1/extended/", c.apiServerHostname)
	resp, err := c.MakeRequest(url, "POST", data)
	if err != nil {
		return "", err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 202 {
		return "", errors.New(fmt.Sprintf("Status code: %d, message: %s", resp.StatusCode, bodyText))
	}
	var respJson interface{}
	var id string
	err = json.Unmarshal(bodyText, &respJson)
	if err != nil {
		return "", err
	}
	respJsonData := respJson.(map[string]interface{})
	for _, value := range respJsonData {
		id = fmt.Sprintf("%v", value)
	}
	return id, nil
}

func (c *APIClient) DeleteEncryptionKey(keyID string) error {
	url := fmt.Sprintf("%s/provisioning/v1/encryption-keys/%s", c.apiServerHostname, keyID)
	resp, err := c.MakeRequest(url, "DELETE", nil)
	if err != nil {
		return err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 202 {
		return errors.New(fmt.Sprintf("Status code: %d, message: %s", resp.StatusCode, bodyText))
	}
	return nil
}
