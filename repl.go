package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	//create an input scanner
	scanner := bufio.NewScanner(os.Stdin)

	//start input loop
	for {
		fmt.Print("Pokedex > ")
		isScanned := scanner.Scan()
		//handle possible input errors
		if !isScanned /*&& scanner.Err() != nil*/ {
			fmt.Printf("Error reading input : %v", scanner.Err())
		}

		//clear the input
		commandInput := cleanInput(scanner.Text())[0]

		//handle the command
		supportedCommands := getSupportedCommands()
		command, ok := supportedCommands[commandInput]
		if ok {
			err := command.callback()
			if err != nil {
				fmt.Printf("Error executing a command: %v\n", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

// a struct to contain supported commands
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func cleanInput(text string) []string {
	resultString := strings.Fields(strings.ToLower(text))
	return resultString
}

func getSupportedCommands() map[string]cliCommand {
	//creating a registry of supported commands
	supportedCommands :=
		map[string]cliCommand{
			"exit": {
				name:        "exit",
				description: "Exit the Pokedex",
				callback:    commandExit,
			},
			"help": {
				name:        "help",
				description: "Displays a help message",
				callback:    commandHelp,
			},
		}
	return supportedCommands
}
