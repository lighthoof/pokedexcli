package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/lighthoof/pokedexcli/internal/pokeAPIHandler"
	"github.com/lighthoof/pokedexcli/internal/pokeCache"
)

type Link struct {
	XMLName xml.Name `xml:"link"`
	Type    string
	URL     string
}

type Links struct {
	XMLName   xml.Name `xml:"links"`
	ConfLinks []Link   `xml:"link"`
}

type config struct {
	base      string
	locations string
	pokemon   string
	next      string
	previous  string
	PokeDex   map[string]pokeAPIHandler.Pokemon
}

var GlobalCache *pokeCache.Cache

func main() {
	//Open the config file
	xmlConfigFile, err := os.Open("url_config.xml")
	if err != nil {
		fmt.Printf("Can't open config : %v\n", err)
	}
	defer xmlConfigFile.Close()

	//Reading the config file
	byteConfig, err := io.ReadAll(xmlConfigFile)
	if err != nil {
		fmt.Printf("Can't read config : %v\n", err)
	}

	//Decode the config into a struct
	var cnfLinks Links
	err = xml.Unmarshal(byteConfig, &cnfLinks)
	if err != nil {
		fmt.Printf("Can't decode config : %v\n", err)
	}

	conf := config{
		base:      cnfLinks.ConfLinks[0].URL,
		locations: cnfLinks.ConfLinks[1].URL,
		pokemon:   cnfLinks.ConfLinks[2].URL,
		next:      "",
		previous:  "",
		PokeDex:   map[string]pokeAPIHandler.Pokemon{},
	}

	//create the cache
	GlobalCache = pokeCache.NewCache(5 * time.Second)
	//initialize PokeDex
	startRepl(conf)
}
