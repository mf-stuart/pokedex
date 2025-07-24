package pokeapi

import (
	"net/http"
)

type Pokemon struct {
	Name           string        `json:"name"`
	Height         int           `json:"height"`
	Weight         int           `json:"weight"`
	BaseExperience int           `json:"base_experience"`
	Stats          []PokemonStat `json:"stats"`
	Types          []Types       `json:"types"`
}
type PokemonStat struct {
	Stat   Stat `json:"stat"`
	Effort int  `json:"effort"`
	Base   int  `json:"base"`
}
type Stat struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	GameIndex    int    `json:"game_index"`
	IsBattleOnly bool   `json:"is_battle_only"`
}
type Types struct {
	Slot int  `json:"slot"`
	Type Type `json:"type"`
}
type Type struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
type pokemonLite struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
type locationArea struct {
	PokemonEncounters []pokemonEncounters `json:"pokemon_encounters"`
}
type pokemonEncounters struct {
	Pokemon pokemonLite `json:"pokemon"`
}
type locationLite struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
type locationAreaApiResponse struct {
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []locationLite `json:"results"`
}
type LocationAreaApiConfig struct {
	Next     string
	Previous string
}

var LocationAreaPage = LocationAreaApiConfig{
	Next:     "https://pokeapi.co/api/v2/location-area/",
	Previous: "",
}

var apiClient = http.Client{}
