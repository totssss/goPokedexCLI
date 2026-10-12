package main

import (
	"time"

	"gitlab.insa-rouen.fr/tsavary/goPokedexCLI/internal/pokeapi"
)

type config struct {
	commandRegistry map[string]cliCommand
	nextMapUrL      string
	previousMapUrl  string
	client          pokeapi.Client
}

func main() {
	config := &config{
		commandRegistry: getCommands(),
		nextMapUrL:      pokeapi.BaseLocationURL,
		previousMapUrl:  "",
		client:          pokeapi.NewClient(5 * time.Second),
	}
	startRepl(config)
}
