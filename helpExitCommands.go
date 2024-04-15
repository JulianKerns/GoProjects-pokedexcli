package main

import (
	"fmt"
	"os"
)

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exits the Pokedex")
	fmt.Println("explore <area-name>: Gives out a List of all the possoble Pokemon encounters in an area")
	fmt.Println("map: Displays the nest 20 locations to explore")
	fmt.Println("mapb: Displays the previouse visited 20 locations")
	return nil
}
func commandExit() error {
	os.Exit(0)
	return nil
}
