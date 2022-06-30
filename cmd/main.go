package main

import (
	"cache-assign/internal/model"
	"cache-assign/internal/strategy"
	"fmt"
)

// main The starting execution point. Demonstrates how to use the cache through its public API.
func main() {
	// Create a cache strategy
	strat := strategy.NewLRUStrategy()

	// Create a cache with a configurable size
	cache, err := model.NewCache(5, &strat)
	if err != nil {
		fmt.Printf("error occured: %s\n", err.Error())
		return
	}

	// Perform some operations
	cache.Insert('A', 10)
	cache.Insert('B', 20)
	cache.Insert('C', 30)
	cache.Insert('D', 40)
	cache.Insert('E', 50)
	cache.Insert('F', 60)
	cache.Insert('C', 100)
	cache.Insert('G', 69)

	// Check the state of the cache.
	fmt.Println("Current state of cache is: ")
	state := cache.StateOfCache()
	for k, v := range state {
		fmt.Printf("Key: %c, Value: %d\n", k, v)
	}

	val, err := cache.Get('D')
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("\nKey D found in cache! Value is %d\n\n", val)
	}

	cache.Insert('B', 70)

	// Check the state of the cache.
	fmt.Println("Current state of cache is: ")
	state = cache.StateOfCache()
	for k, v := range state {
		fmt.Printf("Key: %c, Value: %d\n", k, v)
	}
}
