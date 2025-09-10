package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ManoloEsS/pokedex/internal/cache"
)

type mockCache struct{}

func (c *mockCache) Add(key string, val []byte) {}
func (c *mockCache) Get(key string) ([]byte, bool) {
	return nil, false
}

func TestGetLocationAreas(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"count": 1, "next": "https://example.com/next", "previous": null, "results": [{"name": "test-location", "url": "https://example.com/location/1"}]}`)
	}))
	defer server.Close()

	client := NewClient(0)
	client.BaseURL = server.URL
	cache := cache.NewCache(0)

	resp, err := client.GetLocationAreas(nil, cache)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.Count != 1 {
		t.Errorf("expected count to be 1, got %d", resp.Count)
	}

	if *resp.Next != "https://example.com/next" {
		t.Errorf("expected next to be 'https://example.com/next', got %s", *resp.Next)
	}

	if len(resp.Results) != 1 {
		t.Errorf("expected 1 result, got %d", len(resp.Results))
	}

	if resp.Results[0].Name != "test-location" {
		t.Errorf("expected result name to be 'test-location', got %s", resp.Results[0].Name)
	}
}

// Further testing for the api package could include:
//
// 1. Testing error conditions:
//    - Create a test that uses a mock server that returns an error (e.g., a 500 status code).
//    - Assert that the `GetLocationAreas` function returns an error in this case.
//
// 2. Testing with different responses:
//    - Create tests that use mock servers that return different JSON responses, such as an empty `results` array or a `null` `next` field.
//    - Assert that the function correctly unmarshals these responses.
//
// 3. Testing with a page URL:
//    - Create a test that passes a non-nil `pageURL` to the function.
//    - Assert that the mock server receives a request for the correct URL.
//
// 4. Testing the `GetLocationArea` function (if you add it):
//    - This would be very similar to the `TestGetLocationAreas` test.
//    - You would create a mock server that returns a JSON response for a single location area.
//    - You would then assert that the function correctly unmarshals the response.
