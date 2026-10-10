package main

import (
	"fmt"
)

func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Print("Welcome to the Pokedex!")
	fmt.Print("Usage:")
	fmt.Print()
	for _, cmd := range cfg.commandRegistry {
		fmt.Printf("%v: %v\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
