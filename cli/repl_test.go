package cli

import (
	"testing"
	"time"

	"github.com/ManoloEsS/pokedex/internal/api"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello   world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hello  world",
			expected: []string{"hello", "world"},
		},
		{
			input:    " hello  world",
			expected: []string{"hello", "world"},
		},
		{
			input:    " Hello WorlD ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("%#v and %#v are not the same length", actual, c.expected)
			t.Fail()
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("%#v and %#v are not the same word", word, expectedWord)
				t.Fail()
			}
		}
	}
}

func TestGetCommands(t *testing.T) {
	commands := getCommands()
	expectedCommands := []string{"help", "map", "mapb", "explore", "catch", "exit", "inspect", "pokedex"}

	if len(commands) != len(expectedCommands) {
		t.Errorf("Expected %d commands, but got %d", len(expectedCommands), len(commands))
	}

	for _, cmdName := range expectedCommands {
		if _, ok := commands[cmdName]; !ok {
			t.Errorf("Expected command '%s' not found", cmdName)
		}
	}
}

func TestConfigInitialization(t *testing.T) {
	client := api.NewClient(5*time.Second, 5*time.Minute)
	cfg := &Config{
		PokeapiClient:    client,
		nextLocationsURL: nil,
		prevLocationsURL: nil,
		Pokedex:          make(map[string]api.PokemonData),
	}

	// Test that the config is properly initialized
	// We can't compare clients directly, so we'll test functionality
	if cfg.nextLocationsURL != nil {
		t.Error("Expected nextLocationsURL to be nil initially")
	}

	if cfg.prevLocationsURL != nil {
		t.Error("Expected prevLocationsURL to be nil initially")
	}

	if cfg.Pokedex == nil {
		t.Error("Expected Pokedex to be initialized")
	}

	if len(cfg.Pokedex) != 0 {
		t.Error("Expected Pokedex to be empty initially")
	}

	// Test that we can add to the Pokedex
	testPokemon := api.PokemonData{
		Name: "test-pokemon",
		ID:   1,
	}
	cfg.Pokedex["test"] = testPokemon

	if len(cfg.Pokedex) != 1 {
		t.Error("Expected Pokedex to have one entry after adding")
	}

	if cfg.Pokedex["test"].Name != "test-pokemon" {
		t.Error("Expected to retrieve the added pokemon correctly")
	}
}

func TestCleanInputEdgeCases(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "   ",
			expected: []string{},
		},
		{
			input:    "\t\n\r",
			expected: []string{},
		},
		{
			input:    "SINGLE",
			expected: []string{"single"},
		},
		{
			input:    "Multiple    Spaces   Between",
			expected: []string{"multiple", "spaces", "between"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Input %q: expected %#v, got %#v", c.input, c.expected, actual)
			continue
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("Input %q: expected word %q at index %d, got %q", c.input, c.expected[i], i, actual[i])
			}
		}
	}
}

func TestCommandStructureValidation(t *testing.T) {
	commands := getCommands()

	for cmdName, cmd := range commands {
		if cmd.name != cmdName {
			t.Errorf("Command %q: name field %q doesn't match map key", cmdName, cmd.name)
		}

		if cmd.description == "" {
			t.Errorf("Command %q: description is empty", cmdName)
		}

		if cmd.callback == nil {
			t.Errorf("Command %q: callback is nil", cmdName)
		}
	}
}

// Further testing for the cli package could include:
//
// 1. Testing the command callbacks:
//    - This would involve creating a mock for the `PokeapiClient` to avoid making real API calls.
//    - For each command, you would:
//      - Create a `Config` with the mocked client.
//      - Call the command's callback function.
//      - Assert that the client's methods were called with the expected parameters.
//      - Assert that the `Config` struct is updated correctly (e.g., `NextLocationsURL`, `PrevLocationsURL`).
//
// 2. Testing the REPL loop (`StartRepl`):
//    - This is more complex to test as it involves `os.Stdin`.
//    - One approach is to use an `io.Pipe` to simulate user input and capture the output.
//    - You could write a series of commands to the pipe and then assert that the output is what you expect.
//
// Example for testing `commandMapf` (pseudo-code):
//
// func TestCommandMapf(t *testing.T) {
// 	// Create a mock client
// 	mockClient := &MockPokeapiClient{}
//
// 	// Expected URL for the first call
// 	initialURL := "https://pokeapi.co/api/v2/location-area"
//
// 	// Setup the config with the mock client
// 	cfg := &Config{
// 		PokeapiClient:    mockClient,
// 		NextLocationsURL: &initialURL,
// 	}
//
// 	// Define the expected response from the mock client
// 	mockResponse := api.RespShallowLocations{
// 		Next:     "https://pokeapi.co/api/v2/location-area?offset=20&limit=20",
// 		Previous: nil,
// 		Results:  []api.Result{{Name: "location_1"}, {Name: "location_2"}},
// 	}
//
// 	// Set the expectation on the mock client
// 	mockClient.On("GetLocationAreas", &initialURL).Return(mockResponse, nil)
//
// 	// Call the function
// 	err := commandMapf(cfg)
//
// 	// Assert that there was no error
// 	if err != nil {
// 		t.Errorf("Unexpected error: %v", err)
// 	}
//
// 	// Assert that the client's method was called
// 	mockClient.AssertExpectations(t)
//
// 	// Assert that the config was updated correctly
// 	if *cfg.NextLocationsURL != "https://pokeapi.co/api/v2/location-area?offset=20&limit=20" {
// 		t.Errorf("NextLocationsURL was not updated correctly")
// 	}
// }

// Further testing for the `catch` command could include:
//
// 1. Testing API Error Handling:
//    - Mock the `GetPokemon` method to return an error.
//    - Assert that `commandCatch` correctly propagates this error.
//
// 2. Testing Already Caught Pokemon:
//    - Add a pokemon to the `playerPokedex` map.
//    - Call `commandCatch` with the same pokemon name.
//    - Assert that the function returns the expected "already caught" message/error.
//
// 3. Testing Edge Cases for `BaseExperience`:
//    - Test with a pokemon having a `BaseExperience` of 0.
//    - Test with a very high `BaseExperience` to ensure the probability calculation is correct.
