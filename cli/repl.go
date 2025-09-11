package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ManoloEsS/pokedex/internal/api"
)

type Config struct {
	PokeapiClient    api.Client
	nextLocationsURL *string
	prevLocationsURL *string
	Pokedex          map[string]api.PokemonData
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, string) error
}

func StartRepl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			cleaned := cleanInput(scanner.Text())
			if len(cleaned) < 1 {
				fmt.Println("no command entered...")
				continue
			}
			commandName := cleaned[0]
			parameter := ""
			if len(cleaned) > 1 {
				parameter = cleaned[1]
			}
			if command, ok := commands[commandName]; ok {
				err := command.callback(cfg, parameter)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown command")
				continue
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input: ", err)
		}
	}
}

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	lowered := strings.ToLower(trimmed)
	return strings.Fields(lowered)
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
		"explore": {
			name:        "explore",
			description: "Displays the names of the pokemon in the specified area using 'explore <area name>'",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch the specified pokemon using 'catch <pokemon name>'",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
