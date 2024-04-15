package main

import (
	"encoding/json"
	"fmt"
	"time"

	pokeAPI "github.com/JulianKerns/pokedexcli/internal/pokeAPI"
	pokecache "github.com/JulianKerns/pokedexcli/internal/pokecache"
)

// variable and Pointer that are getting changed and storing the value of the current URL to traverse the locations
var startingConfig pokeAPI.Data = pokeAPI.Data{}
var startingConfigPointer *pokeAPI.Data = &startingConfig
var cache *pokecache.Cache = pokecache.NewCache(500 * time.Millisecond)

func commandMap(cfg *pokeAPI.Data) error {
	if cfg.Next == nil {
		locationsResponse, err := pokeAPI.GetLocations(cfg.Next)
		if err != nil {
			return err
		}
		cfg.Next = locationsResponse.Next
		cfg.Previous = locationsResponse.Previous

		for _, locations := range locationsResponse.Results {
			fmt.Println(locations.Name)
		}

	} else {
		cachedData, ok := cache.Get(*cfg.Next)
		if !ok {
			locationsResponse, err := pokeAPI.GetLocations(cfg.Next)
			if err != nil {
				return err
			}
			cfg.Next = locationsResponse.Next
			cfg.Previous = locationsResponse.Previous

			for _, locations := range locationsResponse.Results {
				fmt.Println(locations.Name)
			}

		} else {
			data := pokeAPI.Data{}
			errJson := json.Unmarshal(cachedData, &data)
			if errJson != nil {
				fmt.Println("Could not format into Go-struct properly")
			}
			cfg.Next = data.Next
			cfg.Previous = data.Previous

			for _, locations := range data.Results {
				fmt.Println(locations.Name)
			}

		}

	}
	return nil
}

func commandMapb(cfg *pokeAPI.Data) error {
	if cfg.Previous == nil {
		fmt.Println("you are on the first page, cant go back before going forward")
		return nil
	}

	if cfg.Previous == nil {
		locationsResponse, err := pokeAPI.GetLocations(cfg.Previous)
		if err != nil {
			return err
		}
		cfg.Next = locationsResponse.Next
		cfg.Previous = locationsResponse.Previous

		for _, locations := range locationsResponse.Results {
			fmt.Println(locations.Name)
		}

	}
	cachedData, ok := cache.Get(*cfg.Previous)
	if !ok {
		locationsResponse, err := pokeAPI.GetLocations(cfg.Previous)
		if err != nil {
			return err
		}
		cfg.Next = locationsResponse.Next
		cfg.Previous = locationsResponse.Previous

		for _, locations := range locationsResponse.Results {
			fmt.Println(locations.Name)
		}

	} else {
		data := pokeAPI.Data{}
		errJson := json.Unmarshal(cachedData, &data)
		if errJson != nil {
			fmt.Println("Could not format into Go-struct properly")
		}
		cfg.Next = data.Next
		cfg.Previous = data.Previous

		for _, locations := range data.Results {
			fmt.Println(locations.Name)
		}

	}
	return nil
}
