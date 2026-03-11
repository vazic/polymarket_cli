package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func mockServer(handler http.HandlerFunc) *httptest.Server {
	return httptest.NewServer(handler)
}

func TestSearchMarkets(t *testing.T) {
	mockResponse := []Event{
		{
			ID:          "123",
			Title:       "Test Event",
			Description: "Test Description",
			StartDate:   "2023-01-01T00:00:00Z",
			EndDate:     "2023-12-31T00:00:00Z",
			Markets: []Market{
				{
					ConditionID:  "0x123",
					Question:     "Will it rain?",
					EndDate:      "2023-12-31T00:00:00Z",
					Outcomes:     `["Yes", "No"]`,
					Prices:       `["0.6", "0.4"]`,
					ClobTokenIds: `["tokA", "tokB"]`,
				},
			},
		},
	}

	server := mockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/events" {
			t.Errorf("Expected path /events, got %s", r.URL.Path)
		}
		if r.URL.Query().Get("query") != "rain" {
			t.Errorf("Expected query 'rain', got %s", r.URL.Query().Get("query"))
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockResponse)
	})
	defer server.Close()

	// Override URL for testing
	originalURL := GammaBaseURL
	GammaBaseURL = server.URL
	defer func() { GammaBaseURL = originalURL }()

	c := NewClient()
	events, err := c.SearchMarkets("rain", 10)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(events) != 1 {
		t.Fatalf("Expected 1 event, got %d", len(events))
	}
	if events[0].Title != "Test Event" {
		t.Errorf("Expected title 'Test Event', got %s", events[0].Title)
	}
}

func TestGetMarket(t *testing.T) {
	mockResponse := []Market{
		{
			ConditionID:  "0x123",
			Question:     "Will it rain?",
			EndDate:      "2023-12-31T00:00:00Z",
			Outcomes:     `["Yes", "No"]`,
			Prices:       `["0.6", "0.4"]`,
			ClobTokenIds: `["tokA", "tokB"]`,
		},
	}

	server := mockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/markets" {
			t.Errorf("Expected path /markets, got %s", r.URL.Path)
		}
		if r.URL.Query().Get("condition_id") != "0x123" {
			t.Errorf("Expected query 'condition_id=0x123', got %s", r.URL.Query().Get("condition_id"))
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockResponse)
	})
	defer server.Close()

	originalURL := GammaBaseURL
	GammaBaseURL = server.URL
	defer func() { GammaBaseURL = originalURL }()

	c := NewClient()
	market, err := c.GetMarket("0x123")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if market == nil {
		t.Fatalf("Expected market, got nil")
	}
	if market.Question != "Will it rain?" {
		t.Errorf("Expected question 'Will it rain?', got %s", market.Question)
	}
}
