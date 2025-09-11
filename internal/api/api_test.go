package api

import (
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
