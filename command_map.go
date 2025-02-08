package main

import (
	"fmt"
	"time"

	"github.com/lighthoof/pokedexcli/internal/pokeAPIHandler"
	"github.com/lighthoof/pokedexcli/internal/pokeCache"
)

func commandMapf(conf *config, mapCache *pokeCache.Cache) error {
	res, err := pokeAPIHandler.GetPokeAPILocation(conf.next)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, location := range res.Results {
		fmt.Println(location.Name)
	}
	cacheEntry := CacheEntry{
		createdAt: time.Now(),
		val:       res.Results,
	}

	conf.next = res.Next
	conf.previous = res.Previous
	return nil
}

func commandMapb(conf *config, mapCache *pokeCache.Cache) error {

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
