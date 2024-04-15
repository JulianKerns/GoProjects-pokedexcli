package main

import (
	"bufio"
	"fmt"
	"os"

	pokeAPI "github.com/JulianKerns/pokedexcli/internal/pokeAPI"
)

func main() {

	type cliCommand struct {
		name        string
		description string
		callback    func() error
	}

	type cliCommandMap struct {
		name        string
		description string
		callback    func(*pokeAPI.Data) error
	}

	commandLines := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
		},
	}

	commandLinesMap := map[string]cliCommandMap{
		"map": {
			name:        "map",
			description: "Displays the next 20 locations to explore",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the 20 locations that have already been explored",
			callback:    commandMapb,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		if input == commandLines["exit"].name {
			commandLines["exit"].callback()
		}

		if input == commandLines["help"].name {
			commandLines["help"].callback()
		}

		if input == commandLinesMap["map"].name {
			commandLinesMap["map"].callback(startingConfigPointer)
		}

		if input == commandLinesMap["mapb"].name {
			commandLinesMap["mapb"].callback(startingConfigPointer)

		}

	}

}
