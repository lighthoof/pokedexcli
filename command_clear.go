package main

import (
	"fmt"
)

func commandClear(conf *config, args ...string) error {
	fmt.Println("Clearing the cache")
	GlobalCache.Clear()
	return nil
}
