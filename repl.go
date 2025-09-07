package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()

		if !ok {
			fmt.Errorf("Error scanning user input: %w", scanner.Err)
		}

		inputWords := cleanInput(scanner.Text())
		if len(inputWords) == 0 {
			continue
		}

		command, exists := getCommands()[inputWords[0]]
		if !exists {
			fmt.Println("Unknown command.")
			continue
		}

		err := command.callback()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
/*	splitStrs := strings.Split(text, " ")
	var cleanInputs []string

	for _, s := range splitStrs {
		normValue := strings.ToLower(strings.TrimSpace(s))
		if normValue == "" {
			continue
		}
		cleanInputs = append(cleanInputs, normValue)
	}
*/
	output := strings.ToLower(text)
	cleanInputs := strings.Fields(output)
	return cleanInputs
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
	"exit": {
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	},
	"help": {
		name: "help",
		description: "Displays a help message",
		callback: commandHelp,
	},
}

}
