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

	commandHelp := func() error {
		fmt.Println("Welcome to the Pokedex!")
		fmt.Println("Usage:")
		fmt.Println("")
		fmt.Println("help: Displays a help message")
		fmt.Println("exit: Exits the Pokedex")
		fmt.Println("map: Displays the nest 20 locations to explore")
		fmt.Println("mapb: Displays the previouse visited 20 locations")
		return nil
	}
	commandExit := func() error {
		os.Exit(0)
		return nil
	}
	// variable and Pointer that are getting changed and storing the value of the current URL to traverse the locations
	var startingConfig pokeAPI.Data = pokeAPI.Data{}
	var startingConfigPointer *pokeAPI.Data = &startingConfig

	commandMap := func(cfg *pokeAPI.Data) error {
		locationsResponse, err := pokeAPI.GetLocations(cfg.Next)
		if err != nil {
			return err
		}
		cfg.Next = locationsResponse.Next
		cfg.Previous = locationsResponse.Previous

		for _, locations := range locationsResponse.Results {
			fmt.Println(locations.Name)
		}
		return nil
	}

	commandMapb := func(cfg *pokeAPI.Data) error {
		if cfg.Previous == nil {
			fmt.Println("you are on the first page, cant go back before going forward")
			return nil

		}
		locationsResponse, err := pokeAPI.GetLocations(cfg.Previous)
		if err != nil {
			return err
		}
		cfg.Next = locationsResponse.Next
		cfg.Previous = locationsResponse.Previous

		for _, locationsb := range locationsResponse.Results {
			fmt.Println(locationsb.Name)
		}

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
