package client

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Client struct {
	conn *Connection
	instName string
}

var (
	ErrClientDbWasNotSelected = errors.New("DataBase wasn't selected")
)

func NewClient(addr string, port int) *Client {
	return &Client{
		conn: NewConnection(addr, port),
	}
}

func (c *Client) InstCreate(instName string) error {
	_, err := c.conn.Post("/inst/create", map[string]interface{}{
		"name": instName,
	})
	if err != nil {
		return err
	}
	c.instName = instName
	return nil
}

func (c *Client) InstList() ([]string, error) {
	req, err := c.conn.Get("/inst/list", map[string]interface{}{})
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

func (c *Client) InstSelect(instName string) error {
	req, err := c.conn.Get(fmt.Sprintf("/inst/%s", instName), map[string]interface{}{})
	if err != nil {
		return err
	}
	defer req.Body.Close()

	type JsonStruct struct {
		Error string	`json:"error"`
	}
	var j JsonStruct

	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&j)
	if err != nil {
		return err
	}

	if j.Error != "" {
		return errors.New("DB is not exist")
	}

	c.instName = instName
	return nil
}

func (c *Client) InstName() string {
	return c.instName
}

// curl -X POST -H 'Content-Type: application/json' 'http://localhost:8080/inst/dz/set' -d '{"name": "one", "value": "1"}'
func (c *Client) Set(name string, value string) error {
	if c.instName == "" {
		return ErrClientDbWasNotSelected
	}

	_, err := c.conn.Post(fmt.Sprintf("/inst/%s/set", c.instName), map[string]interface{}{
		"name": name,
		"value": value,
	})
	if err != nil {
		return err
	}
	return nil
}

// curl -X GET 'http://localhost:8080/inst/dz/get/one'
func (c *Client) Get(name string) (string, error) {
	//r, err := c.conn.Get(fmt.Sprintf("/inst/%s/get/%s", c.dbName, name), map[string]interface{}{})
	//if err != nil {
	//	return "", err
	//}
	//b, err := c.conn.ReadResponse(r)
	//if err != nil {
	//	return b, err
	//}
	//return b, nil
	return "", nil
}
