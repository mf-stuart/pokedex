package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputScanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		inputScanner.Scan()
		text := inputScanner.Text()
		wordsSlice := cleanInput(text)
		command := wordsSlice[0]

		value, exists := supportedCommands[command]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		err := value.callback()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {

	trimmedText := strings.TrimSpace(text)

	words := strings.Fields(trimmedText)

	if len(words) == 0 {
		return []string{""}
	}

	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	return words
}
