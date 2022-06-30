package strategy

import (
	"cache-assign/internal/model"
	"fmt"
)

// LRUEviction Represents the LRU eviction strategy for a cache. It can be
// plugged in while creating a new cache.
type LRUEviction struct {
	// left Housekeeping field, to keep insertion and retrieval efficient. Points to
	// the least recently used node.
	left *model.Node
	// right Housekeeping field, to keep insertion and retrieval efficient. Points to
	// the head of the linked list.
	right *model.Node
}

// NewLRUStrategy Create a new LRU eviction policy instance, for a cache
// instance.
func NewLRUStrategy() LRUEviction {
	lru := LRUEviction{
		left:  &model.Node{0, 0, nil, nil},
		right: &model.Node{0, 0, nil, nil},
	}
	// We are setting the left and right pointers to point to each other at the
	// start, since no elements exist.
	lru.left.Next = lru.right
	lru.right.Prev = lru.left

	return lru
}

// remove Helper function to delete the given node from the linked list.
func (lru *LRUEviction) remove(node *model.Node) {
	prev, next := node.Prev, node.Next
	prev.Next, next.Prev = next, prev
}

// insert Helper function to add the given node to the head of the list, making
// it the most recently accessed node.
func (lru *LRUEviction) insert(node *model.Node) {
	prev, next := lru.right.Prev, lru.right
	prev.Next = node
	next.Prev = node
	node.Next, node.Prev = next, prev
}

// InsertIntoCache A function to insert the given node into the given cache. It
// sets the node to be the most recently used node. If key already exists, update
// its value.
func (lru *LRUEviction) InsertIntoCache(cache *model.Cache, node *model.Node) {
	// Check if key exists.
	if mNode, ok := cache.Storage[node.Key]; ok {
		// If true, remove it from the list
		lru.remove(mNode)
	}
	// Upsert the key
	cache.Storage[node.Key] = node
	// Insert as the most recently used in the linked list
	lru.insert(node)

	// If cache size is exceeded, remove the least recently used node.
	if len(cache.Storage) > cache.Size {
		last := lru.left.Next
		lru.remove(last)
		delete(cache.Storage, last.Key)
	}
}

// GetFromCache Retrieves the value associated with the key in the given cache.
// If the key is not found, returns -1 and an error.
func (lru *LRUEviction) GetFromCache(cache *model.Cache, key int) (int, error) {
	if node, ok := cache.Storage[key]; ok {
		// Make the node the most recently used node.
		lru.remove(node)
		lru.insert(node)
		return node.Val, nil
	}
	return -1, fmt.Errorf("key %d not in cache", key)
}
