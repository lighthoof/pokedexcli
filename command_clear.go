package main

import (
	"fmt"
)

func commandClear(conf *config, input string) error {
	fmt.Println("Clearing the cache")
	GlobalCache.Clear()
	return nil
}
