package main

func main() {
	config := &config{
		commandRegistry: getCommands(),
	}
	startRepl(config)
}
