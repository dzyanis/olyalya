package client

import (
	"fmt"
	"net/http"
	"bytes"
	"bufio"
)

type Connection struct {
	Addr string
	Port int
}

func NewConnection(addr string, port int) *Connection {
	return &Connection{
		Addr: addr,
		Port: port,
	}
}

func (c *Connection) GetFullURL(uri string) string {
	return fmt.Sprintf("http://%s:%d%s", c.Addr, c.Port, uri)
}

func (c *Connection) request(method string, uri string, body []byte) (*http.Response, error) {
	url := c.GetFullURL(uri)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()
	return resp, nil
}

func (c *Connection) methodGet(uri string, body []byte) (*http.Response, error)  {
	return c.request("GET", uri, body)
}

func (c *Connection) methodPost(uri string, body []byte) (*http.Response, error) {
	return c.request("POST", uri, body)
}

func (c *Connection) methodDelete(uri string, body []byte) (*http.Response, error)  {
	return c.request("DELETE", uri, body)
}

func (c *Connection) ReadResponse(response *http.Response) (string, error)  {
	scanner := bufio.NewScanner(response.Body)
	scanner.Split(bufio.ScanRunes)
	var buf bytes.Buffer
	for scanner.Scan() {
		return scanner.Text(), nil;
		buf.WriteString(scanner.Text())
	}
	return buf.String(), nil
}