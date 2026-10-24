package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	pokemonName := args[0]
	if pokemonName == "" {
		return errors.New("Pokemon Name Missing")
	}

	pokemon, hasPokemon := cfg.pokedex[pokemonName]

	if !hasPokemon {
		fmt.Printf("You don't have caught %s\n", pokemonName)
		return fmt.Errorf("You don't have caught %s yet", pokemonName)
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d:\n", pokemon.Height)
	fmt.Printf("Weight: %d:\n", pokemon.Weight)
	fmt.Println("Stats:")

	for _, statStruct := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", statStruct.Stat.Name, statStruct.BaseStat)
	}

	fmt.Println("Types:")
	for _, typeStruct := range pokemon.Types {
		fmt.Printf("  -%s\n", typeStruct.Type.Name)
	}
	return nil
}
