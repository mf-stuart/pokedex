package pokeapi

import (
	"encoding/json"
	"fmt"
	"pokedex/internal/pokecache"
)

func Explore(explorable string) error {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", explorable)
	return getExplore(url)
}

func getExplore(url string) error {

	var jsonResponse locationArea

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
	fmt.Println("Found pokemonLite:")
	for _, entry := range jsonResponse.PokemonEncounters {
		fmt.Println(entry.Pokemon.Name)
	}
	return nil

}
