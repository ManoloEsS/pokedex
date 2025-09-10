# Gemini Project Context: Go Best Practices

This document outlines the context and best practices for developing the Pokedex CLI application.

## Project Overview

This project is a command-line interface (CLI) application for interacting with the PokeAPI. It allows users to explore Pokemon data directly from their terminal.

## Go Version

The project uses Go version `1.25.1`.

## Coding Style and Standards

- **Formatting**: All Go code should be formatted with `gofmt` or `goimports` before committing.
- **Linting**: Use `golangci-lint` to identify and fix common issues.
- **Naming Conventions**: Follow standard Go naming conventions.
  - `camelCase` for internal variables and functions.
  - `PascalCase` for exported identifiers.
- **Comments**: Add comments to explain complex logic, especially in public functions.

## Error Handling

- **Consistent Error Handling**: Use the `if err != nil` pattern consistently.
- **Error Wrapping**: When returning errors from nested calls, wrap them with additional context using `fmt.Errorf` or a similar approach.

## Testing

- **Testing Framework**: Use the standard `testing` package.
- **Unit Tests**: Write unit tests for all new functionality.
- **Table-Driven Tests**: Use table-driven tests for functions with multiple inputs and outputs.
- **Test Coverage**: Aim for high test coverage for all packages.

## Dependencies

- **Dependency Management**: Manage dependencies using Go modules (`go.mod` and `go.sum`).

## API Interaction

- **API Client**: All interactions with the PokeAPI should be handled by the client in the `internal/api` package.
- **Client Struct**: The `Client` struct in `internal/api/client.go` should be used to make API requests.

## Caching

- **Cache Implementation**: The `internal/cache` package provides a cache for API responses.
- **Cache Usage**: Use the cache to store and retrieve frequently accessed data to reduce API calls.
