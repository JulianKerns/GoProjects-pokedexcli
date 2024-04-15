package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pokeAPI "github.com/JulianKerns/pokedexcli/internal/pokeAPI"
)

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

type cliCommandExplore struct {
	name        string
	description string
	callback    func(...string) error
}

var commandLines = map[string]cliCommand{
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

var commandLinesMap = map[string]cliCommandMap{
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

var commandLinesExplore = map[string]cliCommandExplore{
	"explore": {
		name:        "explore",
		description: "Gives out a List of all the possoble Pokemon encounters in the given area",
		callback:    commandExplore,
	},
}

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		cleaned := CleanInput(input)
		if len(cleaned) == 0 {
			continue
		}

		command := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		if command == commandLines["exit"].name {
			commandLines["exit"].callback()
		}

		if command == commandLines["help"].name {
			commandLines["help"].callback()
		}

		if command == commandLinesMap["map"].name {
			commandLinesMap["map"].callback(startingConfigPointer)
		}

		if command == commandLinesMap["mapb"].name {
			commandLinesMap["mapb"].callback(startingConfigPointer)

		}

		if command == commandLinesExplore["explore"].name {
			commandLinesExplore["explore"].callback(args...)
		}

	}
}

func CleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)

	return words
}
