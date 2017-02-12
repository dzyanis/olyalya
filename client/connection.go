package client

import (
	"fmt"
	"net"
	"net/http"
	"bytes"
	"bufio"
	"time"
	"encoding/json"
)

type Connection struct {
	Addr string
	Port int
	Http *http.Client
}

func NewConnection(addr string, port int) *Connection {
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	var http = &http.Client{
		Timeout: time.Second * 10,
		Transport: netTransport,
	}
	return &Connection{
		Addr: addr,
		Port: port,
		Http: http,
	}
}

func (c *Connection) Url(uri string) string {
	return fmt.Sprintf("http://%s:%d%s", c.Addr, c.Port, uri)
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
	return c.Http.Do(req)
}

func (c *Connection) Put(uri string, data map[string]interface{}) (*http.Response, error)  {
	return c.request("PUT", c.Url(uri), data)
}

func (c *Connection) Get(uri string, data map[string]interface{}) (*http.Response, error)  {
	return c.request("GET", c.Url(uri), data)
}

func (c *Connection) Post(uri string, data map[string]interface{}) (*http.Response, error) {
	return c.request("POST", c.Url(uri), data)
}

func (c *Connection) Delete(uri string, data map[string]interface{}) (*http.Response, error)  {
	return c.request("DELETE", c.Url(uri), data)
}