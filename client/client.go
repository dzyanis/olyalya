package client

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Client struct {
	conn *Connection
	dbName string
}

var (
	ErrClientDbWasNotSelected = errors.New("DataBase wasn't selected")
)

func NewClient(addr string, port int) *Client {
	return &Client{
		conn: NewConnection(addr, port),
	}
}

// curl -X POST -H 'Content-Type: application/json' 'http://localhost:8080/db/create' -d '{"name": "dz"}'
func (c *Client) Create(dbName string) error {
	body, err := json.Marshal(map[string]string{
		"name": dbName,
	})
	if err != nil {
		return err
	}
	_, err = c.conn.methodPost("/db/create", body)
	if err != nil {
		return err
	}
	c.dbName = dbName
	return nil
}

func (c *Client) Select(dbName string) error {
	c.dbName = dbName
	return nil
}

// curl -X POST -H 'Content-Type: application/json' 'http://localhost:8080/db/dz/set' -d '{"name": "one", "value": "1"}'
func (c *Client) Set(name string, value string) error {
	if c.dbName == "" {
		return ErrClientDbWasNotSelected
	}

	body, err := json.Marshal(map[string]string{
		"name": name,
		"value": value,
	})
	if err != nil {
		return err
	}
	_, err = c.conn.methodPost(fmt.Sprintf("/db/%s/set", c.dbName), body)
	if err != nil {
		return err
	}
	return nil
}

// curl -X GET 'http://localhost:8080/db/dz/get/one'
func (c *Client) Get(name string) (string, error) {
	if c.dbName == "" {
		return "", ErrClientDbWasNotSelected
	}

	body := []byte{}
	r, err := c.conn.methodGet(fmt.Sprintf("/db/%s/get/%s", c.dbName, name), body)
	if err != nil {
		return "", err
	}
	b, err := c.conn.ReadResponse(r)
	if err != nil {
		return b, err
	}
	return b, nil
}
