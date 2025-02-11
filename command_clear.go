package main

import (
	"fmt"
)

func commandClear(conf *config) error {
	fmt.Println("Clearing the cache")
	GlobalCache.Clear()
	return nil
}
