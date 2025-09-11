# Pokedex CLI

This is a command-line interface (CLI) application for interacting with the [PokeAPI](https://pokeapi.co/). It allows users to explore Pokemon data directly from their terminal.

## Packages

This project is organized into the following packages:

-   `cmd`: This is the main entry point of the application. It initializes the configuration and starts the CLI.
-   `cli`: This package handles the command-line interface. It reads user input, processes commands, and displays the output.
-   `internal/api`: This package is responsible for communicating with the PokeAPI. It provides a client for making HTTP requests and unmarshaling the JSON responses.
-   `internal/cache`: This package provides a caching mechanism for API responses. It helps to reduce the number of requests made to the PokeAPI and improve performance.

## Installation

To install and run this project, you will need to have [Go](https://go.dev/) installed on your system.

1.  Clone the repository:

    ```bash
    git clone https://github.com/ManoloEsS/pokedex.git
    ```

2.  Navigate to the project directory:

    ```bash
    cd pokedex
    ```

3.  Run the application:

    ```bash
    go run ./cmd/main.go
    ```

This will start the Pokedex CLI, and you will be greeted with a `Pokedex >` prompt.

## Usage

The Pokedex CLI provides the following commands:

-   `help`: Displays a help message with a list of available commands.
-   `map`: Displays the names of the next 20 location areas.
-   `mapb`: Displays the names of the previous 20 location areas.
-   `exit`: Exits the Pokedex CLI.
