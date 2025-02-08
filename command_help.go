package main

import (
	"fmt"

	"github.com/lighthoof/pokedexcli/internal/pokeCache"
)

func commandHelp(conf *config, mapCache *pokeCache.Cache) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println()

	//looping over available commands for help output
	for _, command := range getSupportedCommands() {
		name := command.name
		desc := command.description
		fmt.Println(name + ": " + desc)
	}

	return nil
}
