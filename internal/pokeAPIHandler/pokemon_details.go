package pokeAPIHandler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/lighthoof/pokedexcli/internal/pokeCache"
)

func GetPokeAPIPokemon(url string, GlobalCache *pokeCache.Cache) (PokemonResponse, error) {
	//Send GET request to get pokemon data from provided URL
	res, err := http.Get(url)
	if err != nil {
		return PokemonResponse{}, err
	}
	defer res.Body.Close()

	//Read body into a json data string
	jsonData, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonResponse{}, err
	}

	//Check if the record already exists in the cache and decode it
	var decodedRes PokemonResponse
	cache, ok := GlobalCache.Get(url)
	if ok {
		err = json.Unmarshal(cache, &decodedRes)
	} else {
		GlobalCache.Add(url, jsonData)
		err = json.Unmarshal(jsonData, &decodedRes)
	}

	if err != nil {
		return PokemonResponse{}, err
	}

	return decodedRes, nil
}
