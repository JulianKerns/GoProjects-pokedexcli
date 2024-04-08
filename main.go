package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	http "github.com/JulianKerns/pokedexcli/internal"
)

func main() {

	type cliCommand struct {
		name        string
		description string
		callback    func() error
	}
	commandHelp := func() error {
		fmt.Println("Welcome to the Pokedex!")
		fmt.Println("Usage:")
		fmt.Println("")
		fmt.Println("help: Displays a help message")
		fmt.Println("exit: Exits the Pokedex")
		fmt.Println("map: Displays the nest 20 locations to explore")
		fmt.Println("mapb: Displays the 20 locations that have already been explored")
		return nil
	}
	commandExit := func() error {

		return errors.New("")
	}

	commandMap := func() error {
		http.GetMapLocations()
		return nil
	}
	commandMapb := func() error {
		return nil
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
			break

		}

		if input == commandLines["help"].name {
			commandLines["help"].callback()
		}

		if input == commandLines["map"].name {
			commandLines["map"].callback()
		}

	}

}
