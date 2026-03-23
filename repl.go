package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/PicusB/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(conf *config) error
}

type config struct {
	pokeapiClient pokeapi.Client
	nextURL       *string
	previousURL   *string
}

func cleanInput(text string) []string {
	returnString := []string{}
	splitString := strings.Split(text, " ")
	for _, element := range splitString {
		trimmedString := strings.TrimSpace(element)
		if trimmedString == "" {
			continue
		}
		lower := strings.ToLower(trimmedString)
		returnString = append(returnString, lower)
	}
	return returnString
}

func startRepl(conf *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := scanner.Text()
		cleanWords := cleanInput(words)
		if len(cleanWords) == 0 {
			fmt.Println("No Command Received")
			continue
		}
		commandName := cleanWords[0]
		commands := getCommands()
		command, exists := commands[commandName]
		if !exists {
			fmt.Printf("No command by the name of %s\n", commandName)
			continue
		}
		command.callback(conf)
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Provide information to user",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the name of 20 locations in the Pokemon world",
			callback:    commandMap,
		},
	}
}
