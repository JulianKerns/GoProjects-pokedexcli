package internal

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
type Config struct {
	Next     *string     `json:"next"`
	Previous *string     `json:"previous"`
	Results  []Locations `json:"results"`
}

func GetInitialLocations(url string) Config {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 399 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	d := Data{}
	errJson := json.Unmarshal(body, &d)
	if errJson != nil {
		fmt.Println("Could not format into Go-struct properly")
	}

	results := d.Results

	return Config{Next: d.Next, Previous: d.Previous, Results: results}
}
