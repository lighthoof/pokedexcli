package main

import (
	"fmt"

	"github.com/lighthoof/pokedexcli/internal/pokeAPIHandler"
)

func commandMapf(conf *config) error {
	res, err := pokeAPIHandler.GetPokeAPILocation(conf.next)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, location := range res.Results {
		fmt.Println(location.Name)
	}

	conf.next = res.Next
	conf.previous = res.Previous

	return nil
}

func commandMapb(conf *config) error {

	if conf.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	res, err := pokeAPIHandler.GetPokeAPILocation(conf.previous)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, location := range res.Results {
		fmt.Println(location.Name)
	}

	conf.next = res.Next
	conf.previous = res.Previous
	return nil
}
