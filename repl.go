package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(conf config) {
	//create an input scanner
	scanner := bufio.NewScanner(os.Stdin)
	conf.next = conf.base + conf.locations

	//start input loop
	for {
		fmt.Print("Pokedex > ")
		isScanned := scanner.Scan()
		//handle possible input errors
		if !isScanned /*&& scanner.Err() != nil*/ {
			fmt.Printf("Error reading input : %v", scanner.Err())
		}

		//clear the input
		inputString := cleanInput(scanner.Text())
		inputParameter := ""
		if len(inputString) > 1 {
			inputParameter = inputString[1]
		}
		commandInput := inputString[0]

		//handle the command
		supportedCommands := getSupportedCommands()
		command, ok := supportedCommands[commandInput]
		if ok {
			err := command.callback(&conf, inputParameter)
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
	callback    func(*config, string) error
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
			"mapf": {
				name:        "mapf",
				description: "Display next page of map locations",
				callback:    commandMapf,
			},
			"mapb": {
				name:        "mapb",
				description: "Display previous page of map locations",
				callback:    commandMapb,
			},
			"clear": {
				name:        "clear",
				description: "Clear the location cache",
				callback:    commandClear,
			},
			"explore": {
				name:        "explore",
				description: "Explore the location",
				callback:    commandExplore,
			}, "catch": {
				name:        "explore",
				description: "Try to catch a pokemon",
				callback:    commandCatch,
			},
		}
	return supportedCommands
}
