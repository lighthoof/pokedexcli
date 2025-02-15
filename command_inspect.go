package main

import (
	"errors"
	"fmt"
)

func commandInspect(conf *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	name := args[0]
	//Checking the PokeDex for the pokemon data
	pokemon, ok := conf.PokeDex[name]
	if !ok {
		fmt.Println("you have not yet caught that pokemon")
		return nil
	}

	//Displaying the pokemon stats
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("-%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokeType := range pokemon.Types {
		fmt.Printf("- %v\n", pokeType.Type.Name)
	}

	return nil
}
