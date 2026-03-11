package client

import "strconv"

// Market represents a simplified market object returned from Gamma API
type Market struct {
	ConditionID  string `json:"conditionId"`
	Question     string `json:"question"`
	EndDate      string `json:"endDate"`
	Outcomes     string `json:"outcomes"`      // JSON encoded string array
	Prices       string `json:"outcomePrices"` // JSON encoded string array
	ClobTokenIds string `json:"clobTokenIds"`  // JSON encoded string array
}

// Event represents an event object from the /events endpoint containing markets
type Event struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	StartDate   string   `json:"startDate"`
	EndDate     string   `json:"endDate"`
	Markets     []Market `json:"markets"`
}

// SearchMarkets queries the Gamma API for events matching the text
func (c *Client) SearchMarkets(query string, limit int) ([]Event, error) {
	queryParams := map[string]string{
		"query": query,
		"limit": strconv.Itoa(limit),
	}

	var results []Event
	err := c.get(GammaBaseURL, "/events", queryParams, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// GetMarket fetches details for a specific condition ID from the Gamma API
func (c *Client) GetMarket(conditionID string) (*Market, error) {
	queryParams := map[string]string{
		"condition_id": conditionID,
	}

	var results []Market
	err := c.get(GammaBaseURL, "/markets", queryParams, &results)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, nil // Not found
	}
	return &results[0], nil
}
