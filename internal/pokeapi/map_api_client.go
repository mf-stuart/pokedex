package pokeapi

import (
	"encoding/json"
	"fmt"
	"pokedex/internal/pokecache"
)

func Map(url string, loc *LocationAreaApiConfig) error {
	return getMap(url, loc)
}

func getMap(url string, loc *LocationAreaApiConfig) error {

	var jsonResponse locationAreaApiResponse

	entry, exists := pokecache.ClientCache.Get(url)
	if exists {
		fmt.Println("-----Cache Hit!------")
		if err := json.Unmarshal(entry, &jsonResponse); err != nil {
			return fmt.Errorf("could not parse JSON: %w", err)
		}

	} else {
		err := pokeapiGetJSON(url, &jsonResponse)
		if err != nil {
			return err
		}
	}

	for _, location := range jsonResponse.Results {
		fmt.Println(location.Name)
	}

	loc.Next = jsonResponse.Next
	loc.Previous = jsonResponse.Previous

	return nil
}
