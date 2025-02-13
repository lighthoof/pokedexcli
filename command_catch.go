package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lighthoof/pokedexcli/internal/pokeAPIHandler"
)

func commandCatch(conf *config, input string) error {
	catchURL := conf.base + conf.pokemon + "/" + input + "/"
	pokemon, err := pokeAPIHandler.GetPokeAPIPokemon(catchURL, GlobalCache)
	if err != nil {
		fmt.Println(err)
		return err
	}

	//Emulating time spent on catching a pokemon
	fmt.Printf("Throwing a Pokeball at %v ...\n", pokemon.Name)
	time.Sleep(2 * time.Second)

	//Deciding if a pokemon is caught
	if rand.Intn(200) > pokemon.BaseExperience {
		fmt.Printf("%v was caught!\n", pokemon.Name)
	} else {
		fmt.Printf("%v escaped!\n", pokemon.Name)
	}

	return nil
}
