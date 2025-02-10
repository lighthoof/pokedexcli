package pokeAPIHandler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/lighthoof/pokedexcli/internal/pokeCache"
)

func GetPokeAPILocation(url string, GlobalCache *pokeCache.Cache) (Response, error) {
	//Send GET request to get location data from provided URL
	res, err := http.Get(url)
	if err != nil {
		return Response{}, err
	}
	defer res.Body.Close()

	//Read body into a json data string
	jsonData, err := io.ReadAll(res.Body)
	if err != nil {
		return Response{}, err
	}

	//Check if the record already exists in the cache and decode it
	var decodedRes Response
	cache, ok := GlobalCache.Get(url)
	if ok {
		fmt.Println("Getting a record from the cache")
		err = json.Unmarshal(cache, &decodedRes)
	} else {
		fmt.Println("Adding a record to the cache")
		GlobalCache.Add(url, jsonData)
		err = json.Unmarshal(jsonData, &decodedRes)
	}

	if err != nil {
		return Response{}, err
	}

	return decodedRes, nil
}
