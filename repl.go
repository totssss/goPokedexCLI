package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func startRepl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())

		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		if command, ok := getCommands()[commandName]; ok {
			command.callback(cfg, args...)
		} else {
			fmt.Println("Unknown command")
		}
	}
}
