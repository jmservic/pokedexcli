package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/jmservic/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name string
	description string
	callback func(*config) error
	config *config
}

type config struct {
	pokeapiClient *pokeapi.Client
	Next *string
	Previous *string
	Args []string
}


func startRepl(c *pokeapi.Client) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands(c)
	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()

		if !ok {
			fmt.Errorf("Error scanning user input: %w", scanner.Err())
		}

		inputWords := cleanInput(scanner.Text())
		if len(inputWords) == 0 {
			continue
		}

		command, exists := commands[inputWords[0]]
		if !exists {
			fmt.Println("Unknown command.")
			continue
		}
		
		if len(inputWords) > 1 {
			command.config.Args = inputWords[1:]
		} else {
			command.config.Args = nil
		}

		err := command.callback(command.config)

		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	cleanInputs := strings.Fields(output)
	return cleanInputs
}

func getCommands(c *pokeapi.Client) map[string]cliCommand {
	defaultConfig := config{
		pokeapiClient: c,
	}
	mapConfig := config{
		pokeapiClient: c,
	}

	return map[string]cliCommand {
	"exit": {
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
		config: &defaultConfig,
	},
	"help": {
		name: "help",
		description: "Displays a help message",
		callback: commandHelp,
		config: &defaultConfig,
	},
	"map": {
		name: "map",
		description: "Displays next 20 location in the Pokemon world.",
		callback: commandMapf,
		config: &mapConfig,
	},
	"mapb": {
		name: "mapb",
		description: "Displays the previous 20 locations in the Pokemon world.",
		callback: commandMapb,
		config: &mapConfig,
	},
	"explore": {
		name: "explore",
		description: "Explores the the given location name or id (explore <NAME | ID>), printing pokemon at the location.",
		callback: commandExplore,
		config: &defaultConfig,
	},
}

}
