package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// a struct to contain supported commands
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func startRepl() {
	//creating a registry of supported commands
	supportedCommands :=
		map[string]cliCommand{
			"exit": {
				name:        "exit",
				description: "Exit the Pokedex",
				callback:    commandExit,
			},
		}

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
		for _, command := range supportedCommands {
			if commandInput == command.name {
				command.callback()
			} else {
				fmt.Println("Unknown command")
			}
		}

	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	time.Sleep(1 * time.Second)
	os.Exit(0)
	return nil
}

func cleanInput(text string) []string {
	resultString := strings.Fields(strings.ToLower(text))
	return resultString
}
