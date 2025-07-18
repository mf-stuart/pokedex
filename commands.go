package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var supportedCommands map[string]cliCommand

func init() {
	supportedCommands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Show this help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandHelp() error {

	var commandList string
	for _, cmd := range supportedCommands {
		commandList += fmt.Sprintf("\n%s: %s", cmd.name, cmd.description)
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

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
