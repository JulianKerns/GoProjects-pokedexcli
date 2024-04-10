package main

import (
	"bufio"
	"fmt"
	"os"

	internal "github.com/JulianKerns/pokedexcli/internal"
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
		callback    func(*string, *string) error
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
	startingConfig := internal.GetInitialLocations("https://pokeapi.co/api/v2/location-area")
	var startingPointer *internal.Config = &startingConfig

	commandMap := func(mapNext, mapPrevious *string) error {
		for _, locations := range startingPointer.Results {
			fmt.Println(locations.Name)
		}
		*startingPointer = internal.GetInitialLocations(*mapNext)
		fmt.Println(*startingPointer.Next)
		fmt.Println(*startingPointer.Previous)
		return nil
	}

	commandMapb := func(mapNext, mapPrevious *string) error {
		if mapPrevious == nil {
			fmt.Println("cant go back before going forward")
			return nil

		}
		prevlocation := internal.GetInitialLocations(*mapPrevious)
		for _, locationsb := range prevlocation.Results {
			fmt.Println(locationsb.Name)
		}

		*startingPointer = internal.GetInitialLocations(*mapPrevious)
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
			commandLinesMap["map"].callback(startingPointer.Next, startingPointer.Previous)
		}

		if input == commandLinesMap["mapb"].name {
			commandLinesMap["mapb"].callback(startingPointer.Next, startingPointer.Previous)

		}

	}

}
