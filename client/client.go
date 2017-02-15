package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type Json struct {
	Status string	`json:"status"`
	Error string	`json:"error"`

	Value string	`json:"value"`

	Count int 	`json:"count"`
	Names []string	`json:"names"`
}

func JsonReader(body *io.ReadCloser) (Json, error) {
	var j Json
	decoder := json.NewDecoder(*body)
	err := decoder.Decode(&j)
	if err != nil {
		return j, err
	}
	return j, nil
}

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
	req, err := c.conn.Post("/inst/create", map[string]interface{}{
		"name": instName,
	})
	if err != nil {
		return err
	}
	defer req.Body.Close()

	j, err := JsonReader(&req.Body)
	if err != nil {
		return err
	}

	if j.Error!="" {
		return errors.New(j.Error)
	}

	return nil
}

func (c *Client) InstList() ([]string, error) {
	req, err := c.conn.Get("/inst/list", map[string]interface{}{})
	if err != nil {
		return []string{}, err
	}
	defer req.Body.Close()

	j, err := JsonReader(&req.Body)
	if err != nil {
		return []string{}, err
	}

	return j.Names, nil
}

func (c *Client) InstSelect(instName string) error {
	req, err := c.conn.Get(fmt.Sprintf("/inst/%s", instName), map[string]interface{}{})
	if err != nil {
		return err
	}
	defer req.Body.Close()

	j, err := JsonReader(&req.Body)
	if err != nil {
		return err
	}

	if j.Error != "" {
		return errors.New("Instance is not exist")
	}

	c.instName = instName
	return nil
}

func (c *Client) InstName() string {
	return c.instName
}

// curl -X POST -H 'Content-Type: application/json' 'http://localhost:8080/inst/dz/set' -d '{"name": "one", "value": "1"}'
func (c *Client) Set(name string, value string, ttl int) error {
	if c.instName == "" {
		return ErrClientDbWasNotSelected
	}

	req, err := c.conn.Post(fmt.Sprintf("/inst/%s/set", c.instName), map[string]interface{}{
		"name": name,
		"value": value,
		"ttl": ttl,
	})
	if err != nil {
		return err
	}
	defer req.Body.Close()

	j, err := JsonReader(&req.Body)
	if err != nil {
		return err
	}

	if j.Error != "" {
		return errors.New(j.Error)
	}

	return nil
}


// curl -X POST -H 'Content-Type: application/json' 'http://localhost:8080/inst/dz/ttl/set' -d '{"name": "one", "ttl": 2}'
func (c *Client) SetTTL(name string, ttl int) error {
	if c.instName == "" {
		return ErrClientDbWasNotSelected
	}

	req, err := c.conn.Post(fmt.Sprintf("/inst/%s/ttl/set", c.instName), map[string]interface{}{
		"name": name,
		"ttl": ttl,
	})
	if err != nil {
		return err
	}
	defer req.Body.Close()

	j, err := JsonReader(&req.Body)
	if err != nil {
		return err
	}

	if j.Error != "" {
		return errors.New(j.Error)
	}

	return nil
}

// curl -X DELETE -H 'Content-Type: application/json' 'http://localhost:8080/inst/dz/ttl/del' -d '{"name": "one"}'
func (c *Client) DelTTL(name string) error {
	if c.instName == "" {
		return ErrClientDbWasNotSelected
	}

	req, err := c.conn.Delete(fmt.Sprintf("/inst/%s/ttl/del", c.instName), map[string]interface{}{
		"name": name,
	})
	if err != nil {
		return err
	}
	defer req.Body.Close()

	j, err := JsonReader(&req.Body)
	if err != nil {
		return err
	}

	if j.Error != "" {
		return errors.New(j.Error)
	}

	return nil
}

// curl -X GET 'http://localhost:8080/inst/dz/get/one'
func (c *Client) Get(name string) (string, error) {
	if c.instName == "" {
		return "", ErrClientDbWasNotSelected
	}

	req, err := c.conn.Get(fmt.Sprintf("/inst/%s/get/%s", c.instName, name), map[string]interface{}{})
	if err != nil {
		return "", err
	}
	defer req.Body.Close()

	j, err := JsonReader(&req.Body)
	if err != nil {
		return "", err
	}

	if j.Error != "" {
		return "", errors.New(j.Error)
	}

	return j.Value, nil
}
