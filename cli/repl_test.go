package cli

import (
	"testing"
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
	expectedCommands := []string{"help", "map", "mapb", "exit"}

	if len(commands) != len(expectedCommands) {
		t.Errorf("Expected %d commands, but got %d", len(expectedCommands), len(commands))
	}

	for _, cmdName := range expectedCommands {
		if _, ok := commands[cmdName]; !ok {
			t.Errorf("Expected command '%s' not found", cmdName)
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

