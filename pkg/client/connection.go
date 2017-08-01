package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"
)

type Connection struct {
	url        string
	httpClient *http.Client
}

func NewConnection(url string) *Connection {
	var netTransport = &http.Transport{
		Dial:                (&net.Dialer{Timeout: 5 * time.Second}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	var http = &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}

	return &Connection{
		url:        url,
		httpClient: http,
	}
}

func (c *Connection) Url(uri string) string {
	return fmt.Sprintf("%s%s", c.url, uri)
}

func (c *Connection) request(method string, url string, data map[string]interface{}) (*http.Response, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return c.httpClient.Do(req)
}

func (c *Connection) Put(uri string, data map[string]interface{}) (*http.Response, error) {
	return c.request("PUT", c.Url(uri), data)
}

func (c *Connection) Get(uri string, data map[string]interface{}) (*http.Response, error) {
	return c.request("GET", c.Url(uri), data)
}

func (c *Connection) Post(uri string, data map[string]interface{}) (*http.Response, error) {
	return c.request("POST", c.Url(uri), data)
}

func (c *Connection) Delete(uri string, data map[string]interface{}) (*http.Response, error) {
	return c.request("DELETE", c.Url(uri), data)
}
