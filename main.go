package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	resultString := strings.Fields(strings.ToLower(text))
	return resultString
}

func main() {
	fmt.Println("Hello, World!")
}
