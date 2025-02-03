package pokeAPIHandler

import (
	"encoding/json"
	"io"
	"net/http"
)

type Location struct {
	Name string
	Url  string
}

type Response struct {
	Count    int
	Next     string
	Previous string
	Results  []Location
}

func GetPokeAPILocation(url string) (Response, error) {
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

	//Decode json data into a struct with results
	var decodedRes Response
	err = json.Unmarshal(jsonData, &decodedRes)
	if err != nil {
		return Response{}, err
	}

	return decodedRes, nil
}
