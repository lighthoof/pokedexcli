package main

import (
	"fmt"
	"os"

	"github.com/lighthoof/pokedexcli/internal/pokeCache"
)

func commandExit(conf *config, mapCache *pokeCache.Cache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
