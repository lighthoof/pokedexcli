package main

import (
	"fmt"

	"github.com/lighthoof/pokedexcli/internal/pokeAPIHandler"
)

func commandExplore(conf *config, args ...string) error {
	//Appending input location name to the API link
	name := args[0]
	exploreURL := conf.base + conf.locations + "/" + name
	//Getting the location details
	res, err := pokeAPIHandler.GetPokeAPILocationDetails(exploreURL, GlobalCache)
	if err != nil {
		return err
	}

	for _, enounter := range res.PokemonEncounters {
		fmt.Println(enounter.Pokemon.Name)
	}

	return nil
}
