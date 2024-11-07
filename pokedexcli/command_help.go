package main

import "fmt"

func callbackHelp() error {
	fmt.Println("Welcome to the POKEDEXCLi REPL!")
	fmt.Println("Available commands:")
	availableCommands := getCommand()
	for _, cmd := range availableCommands {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
