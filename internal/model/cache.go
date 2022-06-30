package model

import (
	"fmt"
)

// Cache The main cache model. This is what users will instantiate and interact
// it using the provided API. The actual work is offloaded to the eviction
// strategy so that it is plug and play with any strategy the user writes.
type Cache struct {
	// Size The capacity of the cache. Cannot be exceeded. Must be positive.
	Size int
	// Strategy The eviction policy of the cache. Handles the actual insertion and
	// retrieval of key-val pairs.
	Strategy EvictionStrategy
	// Storage Simple hash map for efficient operations on the cache.
	Storage map[int]*Node
}

// NewCache Create a new cache with the given size and eviction strategy. Returns
// pointer to a Cache object, or an error if instantiation was unsuccessful.
func NewCache(size int, strat EvictionStrategy) (*Cache, error) {
	if size <= 0 {
		return nil, fmt.Errorf("unable to create a cache of size 0 or less")
	}

	return &Cache{
		size,
		strat,
		make(map[int]*Node),
	}, nil
}

// Insert Stores the given key value pair into the cache, evicting a key as per
// the eviction strategy if necessary. This method cannot fail.
func (c *Cache) Insert(key int, value int) {
	c.Strategy.InsertIntoCache(c, &Node{
		Key:  key,
		Val:  value,
		Next: nil,
		Prev: nil,
	})
}

// Get Retrieves the value associated with the given key in the cache, if found.
// Otherwise, returns -1 and error.
func (c *Cache) Get(key int) (int, error) {
	val, err := c.Strategy.GetFromCache(c, key)
	if err != nil {
		return -1, err
	}
	return val, nil
}

// StateOfCache Returns a map of key value pairs, which represents the current
// state of the cache. This method cannot fail.
func (c *Cache) StateOfCache() map[int]int {
	m := make(map[int]int)
	for k, v := range c.Storage {
		m[k] = v.Val
	}
	return m
}
