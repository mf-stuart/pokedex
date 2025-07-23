package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedex/internal/pokecache"
)

type locationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
type locationAreaApiResponse struct {
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []locationArea `json:"results"`
}
type LocationAreaApiConfig struct {
	Next     string
	Previous string
}

var LocationAreaEndpoint LocationAreaApiConfig = LocationAreaApiConfig{
	Next:     "https://pokeapi.co/api/v2/location-area/",
	Previous: "",
}

var apiClient http.Client = http.Client{}

func FetchLocationArea(url string, loc *LocationAreaApiConfig) error {
	return getLocationArea(url, loc)
}

func getLocationArea(url string, loc *LocationAreaApiConfig) error {

	var jsonResponse locationAreaApiResponse

	entry, exists := pokecache.LocationCache.Get(url)
	if exists {
		fmt.Println("-----Cache Hit!------")
		if err := json.Unmarshal(entry, &jsonResponse); err != nil {
			return fmt.Errorf("could not parse JSON: %w", err)
		}

		loc.Next = jsonResponse.Next
		loc.Previous = jsonResponse.Previous

		for _, location := range jsonResponse.Results {
			fmt.Println(location.Name)
		}
		return nil
	}

	getMap, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("could not make GET request: %w", err)
	}

	mapRes, err := apiClient.Do(getMap)
	if err != nil {
		return fmt.Errorf("could not make GET request: %w", err)
	}
	defer mapRes.Body.Close()

	data, err := io.ReadAll(mapRes.Body)
	if err != nil {
		return fmt.Errorf("could not read body: %w", err)
	}

	if err := json.Unmarshal(data, &jsonResponse); err != nil {
		return fmt.Errorf("could not parse JSON: %w", err)
	}

	for _, location := range jsonResponse.Results {
		fmt.Println(location.Name)
	}

	loc.Next = jsonResponse.Next
	loc.Previous = jsonResponse.Previous

	pokecache.LocationCache.Add(url, data)

	return nil
}
