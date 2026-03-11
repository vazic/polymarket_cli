package client

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetOrderbook(t *testing.T) {
	mockResponse := Orderbook{
		Market: "tokA",
		Bids: []BookLevel{
			{Price: "0.59", Size: "1000"},
		},
		Asks: []BookLevel{
			{Price: "0.61", Size: "500"},
		},
	}

	server := mockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/book" {
			t.Errorf("Expected path /book, got %s", r.URL.Path)
		}
		if r.URL.Query().Get("token_id") != "tokA" {
			t.Errorf("Expected query 'token_id=tokA', got %s", r.URL.Query().Get("token_id"))
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockResponse)
	})
	defer server.Close()

	originalURL := CLOBBaseURL
	CLOBBaseURL = server.URL
	defer func() { CLOBBaseURL = originalURL }()

	c := NewClient()
	orderbook, err := c.GetOrderbook("tokA")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if orderbook == nil {
		t.Fatalf("Expected orderbook, got nil")
	}
	if len(orderbook.Bids) != 1 || orderbook.Bids[0].Price != "0.59" {
		t.Errorf("Expected 1 bid at 0.59, got %+v", orderbook.Bids)
	}
}
