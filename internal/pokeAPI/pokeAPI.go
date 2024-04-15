package pokeAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	pokecache "github.com/JulianKerns/pokedexcli/internal/pokecache"
)

func GetLocations(pageURL *string) (Data, error) {
	baseURL := "https://pokeapi.co/api/v2/location-area"
	if pageURL == nil {
		pageURL = &baseURL
	}

	res, err := http.Get(*pageURL)
	if err != nil {
		return Data{}, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 399 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return Data{}, err
	}
	cache := pokecache.NewCache(500 * time.Millisecond)
	go cache.Add(*pageURL, body)

	d := Data{}
	errJson := json.Unmarshal(body, &d)
	if errJson != nil {
		fmt.Println("Could not format into Go-struct properly")
	}

	return d, nil
}

func ExploreLocation(location string) ([]string, error) {
	baseURL := "https://pokeapi.co/api/v2/location-area"

	locationURL := fmt.Sprintf(baseURL+"/%s", location)
	res, err := http.Get(locationURL)
	if err != nil {
		return []string{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []string{}, err
	}
	cache := pokecache.NewCache(500 * time.Millisecond)
	go cache.Add(locationURL, body)

	locationData := LocationData{}

	errLoc := json.Unmarshal(body, &locationData)
	if errLoc != nil {
		fmt.Println("Could not format into Go-struct properly")
	}

	pokemonEncounters := locationData.PokemonEncounters
	pokemonList := []string{}
	for _, pokemon := range pokemonEncounters {
		pokemonList = append(pokemonList, pokemon.Pokemon.Name)

	}
	return pokemonList, nil

}

func GettingPokemonInfo(location string) (PokemonInfo, error) {
	baseURL := "https://pokeapi.co/api/v2/pokemon"

	pokemonURL := fmt.Sprintf(baseURL+"/%s", location)
	res, err := http.Get(pokemonURL)
	if err != nil {
		return PokemonInfo{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonInfo{}, err
	}

	cache := pokecache.NewCache(500 * time.Millisecond)
	go cache.Add(pokemonURL, body)

	pokemonData := PokemonInfo{}
	errLoc := json.Unmarshal(body, &pokemonData)
	if errLoc != nil {
		fmt.Println("Could not format into Go-struct properly")
	}

	return pokemonData, nil

}
