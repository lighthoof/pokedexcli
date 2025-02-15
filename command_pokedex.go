package main

import (
	"fmt"
)

func commandPokedex(conf *config, args ...string) error {
	//Check if there are any pokemon in the pokedex
	if len(conf.PokeDex) == 0 {
		fmt.Println("Your pokedex is empty")
		return nil
	}
	//Print the list of caught pokemon
	fmt.Println("Your pokedex:")
	for k, _ := range conf.PokeDex {
		fmt.Println(" - " + k)
	}

	return nil
}
