package main

import (
	"encoding/json"
	"errors"
	"fmt"

	pokeAPI "github.com/JulianKerns/pokedexcli/internal/pokeAPI"
)

func commandExplore(location ...string) error {
	if len(location) != 1 {
		return errors.New("no location area provided")
	}
	locationArea := location[0]
	cachedData, ok := cache.Get(locationArea)
	if !ok {
		pokemonList, err := pokeAPI.ExploreLocation(locationArea)
		if err != nil {
			return err
		}
		fmt.Println("Pokemon found:")
		for _, pokemon := range pokemonList {

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
