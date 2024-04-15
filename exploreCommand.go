package main

import (
	"encoding/json"
	"fmt"

	pokeAPI "github.com/JulianKerns/pokedexcli/internal/pokeAPI"
)

func commandExplore(location string) error {
	cachedData, ok := cache.Get(location)
	if !ok {
		pokemonList, err := pokeAPI.ExploreLocation(location)
		if err != nil {
			return err
		}

		for _, pokemon := range pokemonList {
			fmt.Println("Pokemon found:")
			fmt.Printf("- %s\n", pokemon)
		}
	}

	locationData := pokeAPI.LocationData{}

	errLoc := json.Unmarshal(cachedData, &locationData)
	if errLoc != nil {
		return errLoc
	}

	pokemonEncounters := locationData.PokemonEncounters

	for _, pokemon := range pokemonEncounters {
		fmt.Println("Pokemon found:")
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)

	}

	return nil
}
