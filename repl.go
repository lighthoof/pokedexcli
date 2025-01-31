package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		isScanned := scanner.Scan()
		if isScanned != true && scanner.Err() != nil {
			fmt.Printf("Error reading input : %v", scanner.Err())
		}

		//textInput := scanner.Text()
		commandInput := cleanInput(scanner.Text()) //strings.Fields(strings.ToLower(textInput))
		fmt.Println("Your command was: " + commandInput[0])

	}
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	time.Sleep(2 * time.Second)
	os.Exit(0)
	return nil
}

func cleanInput(text string) []string {
	resultString := strings.Fields(strings.ToLower(text))
	return resultString
}
