// The library works with O(lya-lya) DataBase throw the HTTP protocol
package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// Basic JSON answer
type Json struct {
	Status string	`json:"status"`
	Error string	`json:"error"`

	Value string	`json:"value"`

	Count int 	`json:"count"`
	Names []string	`json:"names"`

	Length int 	`json:"length"`
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

// JSON answer for an array
type JsonArr struct {
	Status string	`json:"status"`
	Error string	`json:"error"`
	Value []string	`json:"value"`
}

func JsonArrayReader(body *io.ReadCloser) (JsonArr, error) {
	var ja JsonArr
	decoder := json.NewDecoder(*body)
	err := decoder.Decode(&ja)
	if err != nil {
		return ja, err
	}
	return ja, nil
}

// JSON answer for a hash
type JsonHash struct {
	Status string		`json:"status"`
	Error string		`json:"error"`
	Value map[string]string `json:"value"`
}

func JsonHashReader(body *io.ReadCloser) (JsonHash, error) {
	var jh JsonHash
	decoder := json.NewDecoder(*body)
	err := decoder.Decode(&jh)
	if err != nil {
		return jh, err
	}
	return jh, nil
}

type Client struct {
	conn *Connection
	instName string
}

var (
	ErrInstanceNotSelected = errors.New("Instance wasn't selected")
)

// Constructor
func NewClient(addr string, port int) *Client {
	return &Client{
		conn: NewConnection(addr, port),
	}
}

func (c *Client) simpleRequest(method string, path string, data map[string]interface{}) (string, error) {
	req, err := c.conn.request(method, c.conn.Url(path), data)
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

	return j.Value, err
}

// Create a new instance
func (c *Client) CreateInstance(instName string) error {
	_, err := c.simpleRequest("POST", "/create", map[string]interface{}{
		"name": instName,
	})

	if err != nil {
		return err
	}

	return nil
}

// Returns list of all instances
func (c *Client) ListInstances() ([]string, error) {
	req, err := c.conn.Get("/list", map[string]interface{}{})
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

// Select context of an instance
func (c *Client) SelectInstance(instName string) error {
	req, err := c.conn.Get(fmt.Sprintf("/in/%s", instName), map[string]interface{}{})
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

func (c *Client) checkContext() error {
	if c.instName == "" {
		return ErrInstanceNotSelected
	}
	return nil
}

func (c *Client) CurrentInstanceName() string {
	return c.instName
}

// Save string
func (c *Client) Set(name string, value string, ttl int) error {
	if err := c.checkContext(); err != nil {
		return err
	}

	path := fmt.Sprintf("/in/%s/set", c.instName)
	_, err := c.simpleRequest("POST", path, map[string]interface{}{
		"name":  name,
		"value": value,
		"ttl":   ttl,
	})

	if err != nil {
		return err
	}

	return nil
}

// Save an array
func (c *Client) SetArray(name string, value []string, ttl int) error {
	if err := c.checkContext(); err != nil {
		return err
	}

	req, err := c.conn.Post(fmt.Sprintf("/in/%s/arr/set", c.instName), map[string]interface{}{
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

// Get an array
func (c *Client) GetArray(name string) ([]string, error) {
	empty := []string{}
	if err := c.checkContext(); err != nil {
		return empty, err
	}

	req, err := c.conn.Get(fmt.Sprintf("/in/%s/get/%s", c.instName, name), map[string]interface{}{})
	if err != nil {
		return empty, err
	}
	defer req.Body.Close()

	j, err := JsonArrayReader(&req.Body)
	if err != nil {
		return empty, err
	}

	if j.Error != "" {
		return empty, errors.New(j.Error)
	}

	return j.Value, nil
}

// Return a hash
func (c *Client) GetHash(name string) (map[string]string, error) {
	empty := map[string]string{}
	if err := c.checkContext(); err != nil {
		return empty, err
	}

	req, err := c.conn.Get(fmt.Sprintf("/in/%s/get/%s", c.instName, name), map[string]interface{}{})
	if err != nil {
		return empty, err
	}
	defer req.Body.Close()

	j, err := JsonHashReader(&req.Body)
	if err != nil {
		return empty, err
	}

	if j.Error != "" {
		return empty, errors.New(j.Error)
	}

	return j.Value, nil
}

// Save a hash
func (c *Client) SetHash(name string, value map[string]string, ttl int) error {
	if err := c.checkContext(); err != nil {
		return err
	}

	req, err := c.conn.Post(fmt.Sprintf("/in/%s/hash/set", c.instName), map[string]interface{}{
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

// Set time to live
func (c *Client) SetTTL(name string, ttl int) error {
	if err := c.checkContext(); err != nil {
		return err
	}

	path := fmt.Sprintf("/in/%s/ttl/set", c.instName)
	_, err := c.simpleRequest("POST", path, map[string]interface{}{"instance": c.instName,
		"name": name,
		"ttl": ttl,
	})

	if err != nil {
		return err
	}

	return nil
}
// Remove time to live for variable
func (c *Client) DelTTL(name string) error {
	if err := c.checkContext(); err != nil {
		return err
	}

	req, err := c.conn.Delete(fmt.Sprintf("/in/%s/ttl/del", c.instName), map[string]interface{}{
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

func (c *Client) Get(name string) (string, error) {
	if err := c.checkContext(); err != nil {
		return "", err
	}

	req, err := c.conn.Get(fmt.Sprintf("/in/%s/get/%s", c.instName, name), map[string]interface{}{})
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

func (c *Client) Del(name string) error {
	if err := c.checkContext(); err != nil {
		return err
	}

	err := c.checkContext()
	if err != nil {
		return err
	}

	path := fmt.Sprintf("/in/%s/delete/%s", c.instName, name)
	_, err = c.simpleRequest("DELETE", path, map[string]interface{}{})
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetArrayElement(name string, index int) (string, error) {
	if err := c.checkContext(); err != nil {
		return "", err
	}

	req, err := c.conn.Get(fmt.Sprintf("/in/%s/arr/el/get", c.instName), map[string]interface{}{
		"name": name,
		"index": index,
	})
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

func (c *Client) SetArrayElement(name string, index int, value string) error {
	if err := c.checkContext(); err != nil {
		return err
	}

	path := fmt.Sprintf("/in/%s/arr/el/set", c.instName)
	_, err := c.simpleRequest("GET", path, map[string]interface{}{
		"name":  name,
		"index": index,
		"value": value,
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DelArrayElement(name string, index int) error {
	if err := c.checkContext(); err != nil {
		return err
	}

	path := fmt.Sprintf("/in/%s/arr/el/del", c.instName)
	_, err := c.simpleRequest("DELETE", path, map[string]interface{}{
		"name": name,
		"index": index,
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) AddArrayElement(name string, value string) error {
	if err := c.checkContext(); err != nil {
		return err
	}

	path := fmt.Sprintf("/in/%s/arr/el/add", c.instName)
	_, err := c.simpleRequest("POST", path, map[string]interface{}{
		"name":  name,
		"value": value,
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) SetHashElement(name string, key string, value string) error {
	if err := c.checkContext(); err != nil {
		return err
	}

	path := fmt.Sprintf("/in/%s/hash/el/set", c.instName)
	_, err := c.simpleRequest("POST", path, map[string]interface{}{
		"name":  name,
		"key":   key,
		"value": value,
	})

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetHashElement(name string, key string) (string, error) {
	if err := c.checkContext(); err != nil {
		return "", err
	}

	path := fmt.Sprintf("/in/%s/hash/el/get", c.instName)
	res, err := c.simpleRequest("GET", path, map[string]interface{}{
		"name":  name,
		"key":   key,
	})

	if err != nil {
		return "", err
	}

	return res, nil
}

func (c *Client) DelHashElement(name string, key string) error {
	if err := c.checkContext(); err != nil {
		return err
	}

	path := fmt.Sprintf("/in/%s/hash/el/del", c.instName)
	_, err := c.simpleRequest("DELETE", path, map[string]interface{}{
		"name":  name,
		"key":   key,
	})

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Keys() ([]string, error) {
	empty := []string{}

	if err := c.checkContext(); err != nil {
		return empty, err
	}

	req, err := c.conn.Get(fmt.Sprintf("/in/%s", c.instName), map[string]interface{}{})
	if err != nil {
		return empty, err
	}
	defer req.Body.Close()

	j, err := JsonReader(&req.Body)
	if err != nil {
		return empty, err
	}

	if j.Error != "" {
		return empty, errors.New(j.Error)
	}

	return j.Names, err
}

func (c *Client) Destroy(name string) error {
	if c.CurrentInstanceName() == name {
		_ = c.SelectInstance("")
	}

	_, err := c.simpleRequest("DELETE", "/destroy", map[string]interface{}{
		"name":  name,
	})

	if err != nil {
		return err
	}

	return nil
}
