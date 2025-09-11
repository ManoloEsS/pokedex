package api

import (
	"net/url"
	"testing"
	"time"
)

func TestGetLocationAreas(t *testing.T) {
	client := NewClient(5*time.Second, 5*time.Minute)

	resp, err := client.GetLocationAreas(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.Count == 0 {
		t.Errorf("expected a non-zero count")
	}

	if resp.Next == nil {
		t.Errorf("expected a non-nil next URL")
	}

	if len(resp.Results) == 0 {
		t.Errorf("expected at least one result")
	}
}

func TestNewClient(t *testing.T) {
	timeout := 10 * time.Second
	cacheInterval := 5 * time.Minute

	client := NewClient(timeout, cacheInterval)

	// Check that client is properly initialized
	if client.httpClient.Timeout != timeout {
		t.Errorf("expected timeout %v, got %v", timeout, client.httpClient.Timeout)
	}

	// We can't directly access cache interval, but we can verify the cache exists
	// by checking if we can use Add/Get methods
	testKey := "test-key"
	testValue := []byte("test-value")
	client.cache.Add(testKey, testValue)

	value, ok := client.cache.Get(testKey)
	if !ok {
		t.Error("expected to retrieve value from cache")
	}
	if string(value) != string(testValue) {
		t.Errorf("expected %s, got %s", testValue, value)
	}
}

func TestURLConstruction(t *testing.T) {
	tests := []struct {
		name     string
		basePath string
		parts    []string
		expected string
	}{
		{
			name:     "location areas",
			basePath: BaseURL,
			parts:    []string{"location-area"},
			expected: "https://pokeapi.co/api/v2/location-area",
		},
		{
			name:     "specific area",
			basePath: BaseURL,
			parts:    []string{"location-area", "canalave-city-area"},
			expected: "https://pokeapi.co/api/v2/location-area/canalave-city-area",
		},
		{
			name:     "pokemon endpoint",
			basePath: BaseURL,
			parts:    []string{"pokemon", "pikachu"},
			expected: "https://pokeapi.co/api/v2/pokemon/pikachu",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := url.JoinPath(tt.basePath, tt.parts...)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestClientMethodSignatures(t *testing.T) {
	client := NewClient(5*time.Second, 5*time.Minute)

	// Test GetLocationAreas with nil URL
	_, err := client.GetLocationAreas(nil)
	if err == nil {
		t.Log("GetLocationAreas with nil URL executed (may fail due to network)")
	}

	// Test GetLocationAreas with custom URL
	customURL := "https://example.com/custom"
	_, err = client.GetLocationAreas(&customURL)
	if err != nil {
		t.Logf("GetLocationAreas with custom URL failed as expected: %v", err)
	}

	// Test GetAreaData
	_, err = client.GetAreaData("test-area")
	if err != nil {
		t.Logf("GetAreaData failed as expected: %v", err)
	}

	// Test GetPokemonData
	_, err = client.GetPokemonData("test-pokemon")
	if err != nil {
		t.Logf("GetPokemonData failed as expected: %v", err)
	}
}

func TestBaseURLConstant(t *testing.T) {
	expectedBaseURL := "https://pokeapi.co/api/v2"
	if BaseURL != expectedBaseURL {
		t.Errorf("expected BaseURL to be %s, got %s", expectedBaseURL, BaseURL)
	}

	// Verify URL is valid
	_, err := url.Parse(BaseURL)
	if err != nil {
		t.Errorf("BaseURL is not a valid URL: %v", err)
	}
}

// Further testing for the api package could include:
//
// 1. Testing caching:
//    - Call `GetLocationAreas` twice for the same URL.
//    - To verify that the second call is cached, you would need to inspect the cache.
//    - This would require exposing the cache from the client or using a mock.
//
// 2. Unit testing with a mock server:
//    - The current test is an integration test because it makes a real network request.
//    - To write a unit test, you would need to use an `httptest.Server`.
//    - This would require making the `baseURL` in the client configurable.
//    - For example, by making `baseURL` a field of the `Client` struct:
//
//      type Client struct {
//          baseURL    string
//          httpClient http.Client
//          cache      cache.Cache
//      }
//
//      // In the test:
//      server := httptest.NewServer(...)
//      client := NewClient(...)
//      client.baseURL = server.URL // Override the base URL to point to the test server
//
// 3. Testing error conditions:
//    - With a mock server, you could simulate API errors (e.g., 500 status code)
//    - and assert that `GetLocationAreas` returns an appropriate error.
