package netbox

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	baseURL    string
	token      string
	httpClient *http.Client
}

// NewClient crée un nouveau client NetBox
func NewClient(baseURL, token string) *Client {
	return &Client{
		baseURL:    strings.TrimSuffix(baseURL, "/"),
		token:      token,
		httpClient: &http.Client{},
	}
}

// APIError représente une erreur de l'API NetBox
type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("NetBox API error %d: %s", e.StatusCode, e.Message)
}

func (c *Client) newRequest(method, path string, body io.Reader) (*http.Request, error) {
	url := fmt.Sprintf("%s/api/%s", c.baseURL, strings.TrimPrefix(path, "/"))
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Token "+c.token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	return req, nil
}

func (c *Client) doRequest(req *http.Request, result interface{}) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		return &APIError{
			StatusCode: resp.StatusCode,
			Message:    string(body),
		}
	}

	if result != nil {
		return json.Unmarshal(body, result)
	}

	return nil
}

// Get récupère un objet par son ID
func (c *Client) Get(endpoint string, id int, result interface{}) error {
	path := fmt.Sprintf("%s/%d/", endpoint, id)
	req, err := c.newRequest("GET", path, nil)
	if err != nil {
		return err
	}
	return c.doRequest(req, result)
}

// List récupère une liste d'objets avec filtrage
func (c *Client) List(endpoint string, params map[string]string, result interface{}) error {
	path := endpoint + "/"
	if len(params) > 0 {
		query := url.Values{}
		for k, v := range params {
			query.Add(k, v)
		}
		path += "?" + query.Encode()
	}

	req, err := c.newRequest("GET", path, nil)
	if err != nil {
		return err
	}
	return c.doRequest(req, result)
}

// Create crée un nouvel objet
func (c *Client) Create(endpoint string, data interface{}, result interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := c.newRequest("POST", endpoint+"/", strings.NewReader(string(jsonData)))
	if err != nil {
		return err
	}
	return c.doRequest(req, result)
}
