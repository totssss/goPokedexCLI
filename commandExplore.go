package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	locationName := args[0]
	if locationName == "" {
		return errors.New("Location Name Missing")
	}

	location, err := cfg.client.GetLocation(locationName)

	if err != nil {
		return err
	}

	for _, pokemonEncounter := range location.PokemonEncounters {
		fmt.Println(pokemonEncounter.Pokemon.Name)
	}

	return nil
}
