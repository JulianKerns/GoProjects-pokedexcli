package main

import (
	"fmt"

	pokeAPI "github.com/JulianKerns/pokedexcli/internal/pokeAPI"
)

var pokedex = make(map[string]pokeAPI.PokemonInfo)

func commandInspect(pokemon ...string) error {
	pokemonName := pokemon[0]
	pokemonInfo, ok := pokedex[pokemonName]
	if !ok {
		fmt.Println("Cant inspect a Pokemon that has not been caught!")
		return nil
	}
	name := pokemonInfo.Name
	heigth := pokemonInfo.Height
	weigth := pokemonInfo.Weight
	stats := pokemonInfo.Stats
	types := pokemonInfo.Types

	fmt.Printf("name: %s\n", name)
	fmt.Printf("heigth: %v\n", heigth)
	fmt.Printf("weigth: %v\n", weigth)
	fmt.Println("Stats:")
	for _, stat := range stats {
		statName := stat.Stat.Name
		baseStat := stat.BaseStat
		fmt.Printf("-%s: %v\n", statName, baseStat)
	}
	fmt.Println("Types: ")
	for _, typing := range types {
		pokemonTypes := typing.Type.Name
		fmt.Printf("- %s\n", pokemonTypes)

	}

	return nil
}

func commandPokedex() error {
	if len(pokedex) == 0 {
		fmt.Println("You have not caugth any Pokemon so far!")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for key := range pokedex {
		fmt.Printf("- %s\n", key)

	}

	return nil

}
