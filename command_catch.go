package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/lighthoof/pokedexcli/internal/pokeAPIHandler"
)

func commandCatch(conf *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	//Appending input pokemon name to the API link
	name := args[0]
	catchURL := conf.base + conf.pokemon + "/" + name + "/"
	pokemon, err := pokeAPIHandler.GetPokeAPIPokemon(catchURL, GlobalCache)
	if err != nil {
		return err
	}

	//Emulating time spent on catching a pokemon
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon.Name)
	time.Sleep(2 * time.Second)

	//Deciding if a pokemon is caught
	if rand.Intn(200) > pokemon.BaseExperience {
		fmt.Printf("%v escaped!\n", pokemon.Name)
		return nil
	}

	//When a pokemon is caught, add it to Pokedex
	fmt.Printf("%v was caught!\n", pokemon.Name)
	_, ok := conf.PokeDex[pokemon.Name]
	if !ok {
		fmt.Println("You may now inspect it with the inspect command.")
	}
	conf.PokeDex[pokemon.Name] = pokemon

	return nil
}
