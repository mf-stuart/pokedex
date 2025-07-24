package pokeapi

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"pokedex/internal/pokecache"
)

const MAX_XP = 400

func Catch(pokemon string, pokedex *map[string]Pokemon) error {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon)
	err := getCatch(url, pokedex)
	if err != nil {
		return err
	}
	return nil
}

func getCatch(url string, pokedex *map[string]Pokemon) error {

	var jsonResponse Pokemon

	entry, exists := pokecache.ClientCache.Get(url)
	if exists {
		fmt.Println("-----Cache Hit!------")
		if err := json.Unmarshal(entry, &jsonResponse); err != nil {
			return fmt.Errorf("could not parse JSON: %w", err)
		}

	} else {
		err := pokeapiGetJSON(url, &jsonResponse)
		if err != nil {
			return fmt.Errorf("pokemon not found")
		}
	}

	baseExperience := jsonResponse.BaseExperience
	name := jsonResponse.Name

	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	catchChance := 1.0 - float64(baseExperience)/float64(MAX_XP)
	randChance := rand.Float64()
	if randChance <= catchChance {
		fmt.Printf("%s was caught!\n", name)
		(*pokedex)[url] = jsonResponse
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}
