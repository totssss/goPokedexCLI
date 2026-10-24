package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	pokemonName := args[0]
	if pokemonName == "" {
		return errors.New("Pokemon Name Missing")
	}

	pokemon, err := cfg.client.GetPokemon(pokemonName)

	if err != nil {
		return err
	}

	catchChance := pokemon.BaseExperience

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	throwSuccess := rand.Intn(catchChance) <= 40

	if throwSuccess {
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.pokedex[pokemonName] = pokemon
		return nil
	}

	fmt.Printf("%s escaped!\n", pokemonName)
	return nil
}
