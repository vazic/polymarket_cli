package client

// Orderbook represents the response from the CLOB API book endpoint
type Orderbook struct {
	Market string      `json:"market"`
	Bids   []BookLevel `json:"bids"`
	Asks   []BookLevel `json:"asks"`
}

// BookLevel represents a single price level in the orderbook
type BookLevel struct {
	Price string `json:"price"`
	Size  string `json:"size"`
}

// GetOrderbook fetches the CLOB API order book for a specific outcome token
func (c *Client) GetOrderbook(tokenID string) (*Orderbook, error) {
	queryParams := map[string]string{
		"token_id": tokenID,
	}

	var result Orderbook
	err := c.get(CLOBBaseURL, "/book", queryParams, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
