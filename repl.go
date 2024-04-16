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

type cliCommandExploreCatch struct {
	name        string
	description string
	callback    func(...string) error
}

var commandLines = map[string]cliCommand{
	"help": {
		name:        "help",
		description: "Displays a help message!",
		callback:    commandHelp,
	},
	"exit": {
		name:        "exit",
		description: "Exits the Pokedex!",
		callback:    commandExit,
	},
	"pokedex": {
		name:        "pokedex",
		description: "Displays all the Pokemon you caught so far!",
		callback:    commandPokedex,
	},
}

var commandLinesMap = map[string]cliCommandMap{
	"map": {
		name:        "map",
		description: "Displays the next 20 locations to explore!",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Displays the 20 locations that have already been explored!",
		callback:    commandMapb,
	},
}

var commandLinesExploreCatch = map[string]cliCommandExploreCatch{
	"explore": {
		name:        "explore",
		description: "Gives out a List of all the possoble Pokemon encounters in the given area!",
		callback:    commandExplore,
	},
	"catch": {
		name:        "catch",
		description: "Tries to catch the given Pokemon!",
		callback:    commandCatch,
	},
	"inspect": {
		name:        "inspect",
		description: "Gives out vital information about a caught Pokemon!",
		callback:    commandInspect,
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

		if command == commandLines["pokedex"].name {
			commandLines["pokedex"].callback()
		}

		if command == commandLinesMap["map"].name {
			commandLinesMap["map"].callback(startingConfigPointer)
		}

		if command == commandLinesMap["mapb"].name {
			commandLinesMap["mapb"].callback(startingConfigPointer)

		}

		if command == commandLinesExploreCatch["explore"].name {
			commandLinesExploreCatch["explore"].callback(args...)
		}

		if command == commandLinesExploreCatch["catch"].name {
			commandLinesExploreCatch["catch"].callback(args...)
		}

		if command == commandLinesExploreCatch["inspect"].name {
			commandLinesExploreCatch["inspect"].callback(args...)
		}

	}
}

func CleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)

	return words
}
