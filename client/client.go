package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var (
	GammaBaseURL = "https://gamma-api.polymarket.com"
	CLOBBaseURL  = "https://clob.polymarket.com"
)

// Client is the base HTTP client for interacting with Polymarket APIs
type Client struct {
	httpClient *http.Client
}

// NewClient creates a new Polymarket API client
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}

// get performs a generic HTTP GET request to the specified endpoint with query params
func (c *Client) get(baseURL, endpoint string, queryParams map[string]string, target interface{}) error {
	u, err := url.Parse(baseURL + endpoint)
	if err != nil {
		return fmt.Errorf("invalid URL: %w", err)
	}

	if len(queryParams) > 0 {
		q := u.Query()
		for k, v := range queryParams {
			q.Add(k, v)
		}
		u.RawQuery = q.Encode()
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(bodyBytes))
	}

	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}
