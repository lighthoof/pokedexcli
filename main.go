package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	resultString := strings.Fields(strings.ToLower(text))
	return resultString
}

func main() {

	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Pokedex > ")
		isScanned := scanner.Scan()
		if isScanned != true && scanner.Err != nil {
			fmt.Print("Error reading input!")
		}
		textInput := scanner.Text()
		formattedInput := strings.Fields(strings.ToLower(textInput))
		fmt.Println("Your command was: " + formattedInput[0])

	}

}
