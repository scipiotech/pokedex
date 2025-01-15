package main

import (
	"bufio"
	"fmt"
	"github.com/scipiotech/pokedex/internal/pokeapi"
	"github.com/scipiotech/pokedex/internal/pokecache"
	"os"
	"strings"
)

type config struct {
	pokeapiClient        pokeapi.Client
	nextLocationsURL     *string
	previousLocationsURL *string
	cache                *pokecache.Cache
}

func startRepl(cfg *config) {

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command.")
		}
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the names of 20 location areas",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous page of 20 location areas",
			callback:    commandMapb,
		},
	}
}
