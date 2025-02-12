package main

import (
	"fmt"

	"github.com/lighthoof/pokedexcli/internal/pokeAPIHandler"
)

func commandExplore(conf *config, input string) error {
	exploreURL := conf.base + conf.locations + "/" + input
	res, err := pokeAPIHandler.GetPokeAPILocationDetails(exploreURL, GlobalCache)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, enounter := range res.PokemonEncounters {
		fmt.Println(enounter.Pokemon.Name)
	}

	return nil
}
