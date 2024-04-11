package pokeAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

	d := Data{}
	errJson := json.Unmarshal(body, &d)
	if errJson != nil {
		fmt.Println("Could not format into Go-struct properly")
	}

	return d, nil
}