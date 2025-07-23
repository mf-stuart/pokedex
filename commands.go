package main

import (
	"fmt"
	"os"
	"pokedex/internal/pokeapi"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func(args []string, configPointer *pokeapi.LocationAreaApiConfig) error
}

var supportedCommands map[string]cliCommand

func init() {
	supportedCommands = map[string]cliCommand{
		"help": {
			Name:        "help",
			Description: "Show this help message",
			Callback:    commandHelp,
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
	}
}

func commandHelp(args []string, loc *pokeapi.LocationAreaApiConfig) error {

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

func commandExit(args []string, loc *pokeapi.LocationAreaApiConfig) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(args []string, loc *pokeapi.LocationAreaApiConfig) error {
	err := pokeapi.FetchLocationArea(loc.Next, loc)
	if err != nil {
		return err
	}
	return nil
}

func commandMapb(args []string, loc *pokeapi.LocationAreaApiConfig) error {
	err := pokeapi.FetchLocationArea(loc.Previous, loc)
	if err != nil {
		return err
	}
	return nil
}

func commandExplore(args []string, loc *pokeapi.LocationAreaApiConfig) error {
	if len(args) != 1 {
		return fmt.Errorf("Usage: pokedex explore <key>")
	}

	return nil
}
