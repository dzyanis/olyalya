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
	_, err := c.conn.Post("/db/create", map[string]interface{}{
		"name": dbName,
	})
	if err != nil {
		return err
	}
	c.dbName = dbName
	return nil
}

func (c *Client) DbList() ([]string, error) {
	req, err := c.conn.Get("/db/list", map[string]interface{}{})
	if err != nil {
		return []string{}, err
	}
	defer req.Body.Close()

	type JsonList struct {
		Count int 	`json:"count"`
		Names []string	`json:"names"`
	}
	var jl JsonList

	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&jl)
	if err != nil {
		return []string{}, err
	}

	return jl.Names, nil
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

	_, err := c.conn.Post(fmt.Sprintf("/db/%s/set", c.dbName), map[string]interface{}{
		"name": name,
		"value": value,
	})
	if err != nil {
		return err
	}
	return nil
}

// curl -X GET 'http://localhost:8080/db/dz/get/one'
func (c *Client) Get(name string) (string, error) {
	//r, err := c.conn.Get(fmt.Sprintf("/db/%s/get/%s", c.dbName, name), map[string]interface{}{})
	//if err != nil {
	//	return "", err
	//}
	//b, err := c.conn.ReadResponse(r)
	//if err != nil {
	//	return b, err
	//}
	//return b, nil
}
