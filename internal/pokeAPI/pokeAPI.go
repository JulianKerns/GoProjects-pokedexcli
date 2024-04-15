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

type Locations struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Data struct {
	Count    int         `json:"count"`
	Next     *string     `json:"next"`
	Previous *string     `json:"previous"`
	Results  []Locations `json:"results"`
}

type LocationData struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int `json:"chance"`
				ConditionValues []struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"condition_values"`
				MaxLevel int `json:"max_level"`
				Method   struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

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
