package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
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

	startUrl := cnfLinks.ConfLinks[0].URL + cnfLinks.ConfLinks[1].URL
	startRepl(startUrl)
}
