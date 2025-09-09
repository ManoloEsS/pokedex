package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ManoloEsS/pokedex/internal/api"
	internal "github.com/ManoloEsS/pokedex/internal/api"
)

type Config struct {
	pokeapiClient    *api.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &Config{
		pokeapiClient:    internal.NewClient(5),
		nextLocationsURL: nil,
		prevLocationsURL: nil,
	}

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			cleaned := cleanInput(scanner.Text())
			if len(cleaned) < 1 {
				fmt.Println("no command entered...")
				continue
			}
			commandName := cleaned[0]
			if command, ok := getCommands()[commandName]; ok {
				err := command.callback(cfg)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown command")
				continue
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input: ", err)
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	return strings.Fields(strings.TrimSpace(lowered))
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 locations",
			callback:    commandMapb,
		},
	}
}
