package pokeAPIHandler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/lighthoof/pokedexcli/internal/pokeCache"
)

func GetPokeAPILocation(url string, GlobalCache *pokeCache.Cache) (LocationListResponse, error) {
	//Send GET request to get location data from provided URL
	res, err := http.Get(url)
	if err != nil {
		return LocationListResponse{}, err
	}
	defer res.Body.Close()

	//Read body into a json data string
	jsonData, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationListResponse{}, err
	}

	//Check if the record already exists in the cache and decode it
	var decodedRes LocationListResponse
	cache, ok := GlobalCache.Get(url)
	if ok {
		err = json.Unmarshal(cache, &decodedRes)
	} else {
		GlobalCache.Add(url, jsonData)
		err = json.Unmarshal(jsonData, &decodedRes)
	}

	if err != nil {
		return LocationListResponse{}, err
	}

	return decodedRes, nil
}

func GetPokeAPILocationDetails(url string, GlobalCache *pokeCache.Cache) (LocationDetailResponse, error) {
	//Send GET request to get location data from provided URL
	res, err := http.Get(url)
	if err != nil {
		return LocationDetailResponse{}, err
	}
	defer res.Body.Close()

	//Read body into a json data string
	jsonData, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationDetailResponse{}, err
	}

	//Check if the record already exists in the cache and decode it
	var decodedRes LocationDetailResponse
	cache, ok := GlobalCache.Get(url)
	if ok {
		err = json.Unmarshal(cache, &decodedRes)
	} else {
		GlobalCache.Add(url, jsonData)
		err = json.Unmarshal(jsonData, &decodedRes)
	}

	if err != nil {
		return LocationDetailResponse{}, err
	}

	return decodedRes, nil
}
