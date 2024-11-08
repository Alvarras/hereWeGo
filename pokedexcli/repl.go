package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(" >")

		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]

		availableCommands := getCommand()
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Unknown command: ", command)
			continue
		}
		err := command.callback(cfg)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display help message",
			callback:    callbackHelp,
		},
		"map": {
			name:        "map",
			description: "List the next location areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapBack",
			description: "List the previous location areas",
			callback:    callbackMapBack,
		},
		"exit": {
			name:        "exit",
			description: "Exit the REPL",
			callback:    callbackExit,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
