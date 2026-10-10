package main

import (
	"errors"
	"fmt"
	"gopokedexcli/internal/pokeapi"
)

func commandMap(cfg *config) error {
	return commandMapGeneral(cfg.nextMapUrL, cfg)
}

func commandMapBack(cfg *config) error {
	if cfg.previousMapUrl == "" {
		fmt.Println("you're on the first page")
	}
	return commandMapGeneral(cfg.previousMapUrl, cfg)
}

func commandMapGeneral(url string, cfg *config) error {
	if url == "" {
		return errors.New("pageUrl should never be empty")
	}

	locationResponse, err := pokeapi.RequestLocations(url)

	if err != nil {
		return err
	}

	cfg.nextMapUrL = locationResponse.Next
	cfg.previousMapUrl = locationResponse.Previous

	for _, location := range locationResponse.Results {
		fmt.Println(location.Name)
	}

	return nil
}
