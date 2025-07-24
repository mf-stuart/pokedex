package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedex/internal/pokecache"
)

func pokeapiGetJSON[T any](url string, target *T) error {
	getReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("pokeapiGetJSON: %w", err)
	}

	res, err := apiClient.Do(getReq)
	if err != nil {
		return fmt.Errorf("error while fetching %s: %w", url, err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, target); err != nil {
		return err
	}

	pokecache.ClientCache.Add(url, data)
	return nil
}
