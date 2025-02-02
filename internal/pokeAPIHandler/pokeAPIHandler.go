package pokeAPIHandler

import (
	"encoding/json"
	"io"
	"net/http"
)

type Location struct {
	name string
	url  string
}

type response struct {
	count    int
	next     string
	previous string
	results  []Location
}

func GetPokeAPILocation(url string) ([]Location, error) {
	//Send GET request to get location data from provided URL
	res, err := http.Get(url)
	if err != nil {
		return []Location{}, err
	}
	defer res.Body.Close()

	//Read body into a json data string
	jsonData, err := io.ReadAll(res.Body)
	if err != nil {
		return []Location{}, err
	}

	//Decode json data into a struct with results
	var decodedRes response
	err = json.Unmarshal(jsonData, &decodedRes)
	if err != nil {
		return []Location{}, err
	}

	return decodedRes.results, nil
}
