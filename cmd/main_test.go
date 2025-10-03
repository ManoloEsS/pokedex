package main

import (
	"testing"
	"time"

	"github.com/ManoloEsS/pokedex/cli"
	"github.com/ManoloEsS/pokedex/internal/api"
)

// TestMainPackageImports verifies that all necessary packages are imported correctly
func TestMainPackageImports(t *testing.T) {
	// This test ensures that the main package compiles and all imports are valid
	// If this test runs, it means the package structure is correct
	t.Log("Main package imports are valid")
}

// TestConfigCreation verifies that a Config can be created with the expected structure
func TestConfigCreation(t *testing.T) {
	// Create a config similar to what main() does
	cfg := &cli.Config{
		PokeapiClient: api.NewClient(5*time.Second, 5*time.Minute),
		Pokedex:       make(map[string]api.PokemonData),
	}

	// Verify the config is not nil
	if cfg == nil {
		t.Fatal("Config should not be nil")
	}

	// Verify Pokedex map is initialized
	if cfg.Pokedex == nil {
		t.Error("Pokedex map should be initialized")
	}

	// Verify Pokedex is empty initially
	if len(cfg.Pokedex) != 0 {
		t.Errorf("Expected Pokedex to be empty, got %d entries", len(cfg.Pokedex))
	}
}

// TestClientCreation verifies that the API client can be created
func TestClientCreation(t *testing.T) {
	timeout := 5 * time.Second
	cacheInterval := 5 * time.Minute

	// Create a client - this verifies it compiles and returns without error
	client := api.NewClient(timeout, cacheInterval)

	// Use type assertion to verify we got a Client back
	var _ api.Client = client
}
