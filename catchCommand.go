package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"

	pokeAPI "github.com/JulianKerns/pokedexcli/internal/pokeAPI"
)

func commandCatch(pokemon ...string) error {
	if len(pokemon) != 1 {
		return errors.New("no location area provided")
	}
	pokemonName := pokemon[0]
	cachedData, ok := cache.Get(pokemonName)

	if !ok {
		pokemonInfo, err := pokeAPI.GettingPokemonInfo(pokemonName)
		if err != nil {
			return err
		}
		baseExperience := pokemonInfo.BaseExperience
		fmt.Printf("Trying to Catch %s:\n", pokemonInfo.Name)
		if CatchingPokemon(baseExperience) {
			fmt.Printf("%s was caught\n", pokemonInfo.Name)
			pokedex[pokemonName] = pokemonInfo
			fmt.Println(pokedex[pokemonName].Height)
		} else {
			fmt.Printf("%s escaped the ball!\n", pokemonInfo.Name)
		}
		return nil
	}
	pokemonData := pokeAPI.PokemonInfo{}
	errLoc := json.Unmarshal(cachedData, &pokemonData)
	if errLoc != nil {
		fmt.Println("Could not format into Go-struct properly")
	}
	baseExperience := pokemonData.BaseExperience
	fmt.Printf("Trying to Catch %s:\n", pokemonData.Name)
	if CatchingPokemon(baseExperience) {
		fmt.Printf("%s was caught\n", pokemonData.Name)
		pokedex[pokemonName] = pokemonData
		fmt.Println(pokedex[pokemonName].Height)
	} else {
		fmt.Printf("%s escaped the ball!\n", pokemonData.Name)
	}
	return nil
}

func CatchingPokemon(catchRate int) bool {

	if catchRate >= 300 {
		randomNumber := rand.Intn(100)
		if randomNumber <= 10 {
			return true
		} else {
			return false
		}
	}
	if catchRate >= 200 {
		randomNumber := rand.Intn(100)
		if randomNumber <= 25 {
			return true
		} else {
			return false
		}

	}
	if catchRate >= 100 {
		randomNumber := rand.Intn(100)
		if randomNumber <= 50 {
			return true
		} else {
			return false
		}
	}
	if catchRate < 100 {
		randomNumber := rand.Intn(100)
		if randomNumber <= 75 {
			return true
		} else {
			return false
		}
	}
	return false

}
