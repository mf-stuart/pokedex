package main

import (
	"encoding/json"
	"fmt"
	"os"
	"pokedex/internal/pokeapi"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func(args []string, pokedex *map[string]pokeapi.Pokemon, configPointer *pokeapi.LocationAreaApiConfig) error
}

var supportedCommands map[string]cliCommand

func init() {
	supportedCommands = map[string]cliCommand{
		"help": {
			Name:        "help",
			Description: "Show this help message",
			Callback:    commandHelp,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Show pokedex",
			Callback:    commandPokedex,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"map": {
			Name:        "map",
			Description: "Displays the next 20 locations",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the previous 20 locations",
			Callback:    commandMapb,
		},
		"explore": {
			Name:        "explore",
			Description: "takes a location arg to explore that location",
			Callback:    commandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Tries to catch a pokemonLite",
			Callback:    commandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Displays pokemon information",
			Callback:    commandInspect,
		},
	}
}

func commandHelp(args []string, pokedex *map[string]pokeapi.Pokemon, loc *pokeapi.LocationAreaApiConfig) error {

	var commandList string
	for _, cmd := range supportedCommands {
		commandList += fmt.Sprintf("\n%s: %s", cmd.Name, cmd.Description)
	}

	fmt.Print(
		"Welcome to the Pokedex!" +
			"\nUsage:" +
			"\n\n" +
			commandList +
			"\n",
	)
	return nil
}

func commandPokedex(args []string, pokedex *map[string]pokeapi.Pokemon, loc *pokeapi.LocationAreaApiConfig) error {
	if len(*pokedex) < 1 {
		return fmt.Errorf("No Pokemon in Pokedex")
	}

	fmt.Println("Your Pokedex:")
	for _, entry := range *pokedex {
		fmt.Printf("- %s\n", entry.Name)
	}

	return nil
}

func commandExit(args []string, pokedex *map[string]pokeapi.Pokemon, loc *pokeapi.LocationAreaApiConfig) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(args []string, pokedex *map[string]pokeapi.Pokemon, loc *pokeapi.LocationAreaApiConfig) error {
	err := pokeapi.Map(loc.Next, loc)
	if err != nil {
		return err
	}
	return nil
}

func commandMapb(args []string, pokedex *map[string]pokeapi.Pokemon, loc *pokeapi.LocationAreaApiConfig) error {
	err := pokeapi.Map(loc.Previous, loc)
	if err != nil {
		return err
	}
	return nil
}

func commandExplore(args []string, pokedex *map[string]pokeapi.Pokemon, loc *pokeapi.LocationAreaApiConfig) error {
	if len(args) != 1 {
		return fmt.Errorf("Usage: explore <location>")
	}
	fmt.Printf("Exploring %s...\n", args[0])
	err := pokeapi.Explore(args[0])
	if err != nil {
		return err
	}
	return nil
}

func commandCatch(args []string, pokedex *map[string]pokeapi.Pokemon, loc *pokeapi.LocationAreaApiConfig) error {
	if len(args) != 1 {
		return fmt.Errorf("Usage: catch <Pokemon>")
	}
	err := pokeapi.Catch(args[0], pokedex)
	if err != nil {
		return err
	}
	return nil
}

func commandInspect(args []string, pokedex *map[string]pokeapi.Pokemon, loc *pokeapi.LocationAreaApiConfig) error {
	if len(args) != 1 {
		return fmt.Errorf("Usage: inspect <Pokemon>")
	}

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", args[0])

	if poke, ok := (*pokedex)[url]; ok {

		bytes, _ := json.MarshalIndent(poke, "", "  ")
		fmt.Println(string(bytes))

	} else {
		return fmt.Errorf("pokemon not in pokedex: %s", args[0])
	}
	return nil
}
