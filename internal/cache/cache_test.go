package cache

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}
func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}

func TestNewCache(t *testing.T) {
	interval := 100 * time.Millisecond
	cache := NewCache(interval)

	// Verify cache is properly initialized
	if cache.cacheData == nil {
		t.Error("expected cacheData to be initialized")
	}

	if cache.mux == nil {
		t.Error("expected mutex to be initialized")
	}

	// Test that we can immediately use the cache
	testKey := "test-key"
	testValue := []byte("test-value")
	cache.Add(testKey, testValue)

	value, ok := cache.Get(testKey)
	if !ok {
		t.Error("expected to find key immediately after adding")
	}
	if string(value) != string(testValue) {
		t.Errorf("expected %s, got %s", testValue, value)
	}
}

func TestReapMethod(t *testing.T) {
	cache := NewCache(1 * time.Hour) // Long interval to prevent automatic reaping
	
	// Add some entries
	oldTime := time.Now().UTC().Add(-2 * time.Hour)
	recentTime := time.Now().UTC().Add(-30 * time.Minute)
	
	// Manually add entries with specific timestamps
	cache.mux.Lock()
	cache.cacheData["old-key"] = cacheEntry{
		createdAt: oldTime,
		val:       []byte("old-value"),
	}
	cache.cacheData["recent-key"] = cacheEntry{
		createdAt: recentTime,
		val:       []byte("recent-value"),
	}
	cache.mux.Unlock()

	// Verify both entries exist
	_, ok := cache.Get("old-key")
	if !ok {
		t.Error("expected old-key to exist before reaping")
	}
	_, ok = cache.Get("recent-key")
	if !ok {
		t.Error("expected recent-key to exist before reaping")
	}

	// Call reap with 1 hour threshold
	now := time.Now().UTC()
	cache.reap(now, 1*time.Hour)

	// Old entry should be gone, recent entry should remain
	_, ok = cache.Get("old-key")
	if ok {
		t.Error("expected old-key to be reaped")
	}
	_, ok = cache.Get("recent-key")
	if !ok {
		t.Error("expected recent-key to remain after reaping")
	}
}

func TestConcurrentAccess(t *testing.T) {
	cache := NewCache(1 * time.Minute)
	numGoroutines := 10
	numOperations := 100

	var wg sync.WaitGroup
	
	// Start multiple goroutines that add data
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				key := fmt.Sprintf("key-%d-%d", id, j)
				value := fmt.Sprintf("value-%d-%d", id, j)
				cache.Add(key, []byte(value))
			}
		}(i)
	}

	// Start multiple goroutines that read data
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				key := fmt.Sprintf("key-%d-%d", id, j)
				_, _ = cache.Get(key) // We don't care about the result, just testing for races
			}
		}(i)
	}

	wg.Wait()

	// If we get here without a race condition, the test passes
	t.Log("Concurrent access test completed successfully")
}

// Further testing for the cache package could include:
//
// 1. Testing the reap loop:
//    - Create a new cache with a very short interval (e.g., 10 milliseconds).
//    - Add an entry.
//    - Wait for a duration longer than the interval.
//    - Try to get the entry and assert that it has been reaped (i.e., `ok` is false).
//
// 2. Testing concurrency:
//    - Create a test that uses multiple goroutines to call `Add` and `Get` on the same cache concurrently.
//    - This can help ensure that the mutex is working correctly and there are no race conditions.
//    - You can use the `go test -race` command to detect race conditions during testing.
