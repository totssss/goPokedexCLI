package main

import "gitlab.insa-rouen.fr/tsavary/goPokedexCLI/internal/pokeapi"

type config struct {
	commandRegistry map[string]cliCommand
	nextMapUrL      string
	previousMapUrl  string
}

func main() {
	config := &config{
		commandRegistry: getCommands(),
		nextMapUrL:      pokeapi.BaseLocationURL,
		previousMapUrl:  "",
	}
	startRepl(config)
}
