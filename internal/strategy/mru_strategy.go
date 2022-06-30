package strategy

import (
	"cache-assign/internal/model"
	"fmt"
)

// MRUEviction Represents the MRU eviction strategy for a cache. It can be
// plugged in while creating a new cache.
type MRUEviction struct {
	// left Housekeeping field, to keep insertion and retrieval efficient. Points to
	// the least recently used node.
	left *model.Node
	// right Housekeeping field, to keep insertion and retrieval efficient. Points to
	// the most recently used node.
	right *model.Node
}

// NewMRUStrategy Create a new MRU eviction policy instance, for a cache
// instance.
func NewMRUStrategy() MRUEviction {
	mru := MRUEviction{
		left:  &model.Node{0, 0, nil, nil},
		right: &model.Node{0, 0, nil, nil},
	}
	mru.left.Next = mru.right
	mru.right.Prev = mru.left

	return mru
}

// remove Helper function to delete the given node from the linked list.
func (mru *MRUEviction) remove(node *model.Node) {
	prev, next := node.Prev, node.Next
	prev.Next, next.Prev = next, prev
}

// insert Helper function to add the given node to the head of the list, making
// it the most recently used node.
func (mru *MRUEviction) insert(node *model.Node) {
	prev, next := mru.right.Prev, mru.right
	prev.Next = node
	next.Prev = node
	node.Prev, node.Next = prev, next
}

// InsertIntoCache A function to insert the given node into the given cache. It
// sets the node to be the most recently used node. If key already exists, update
// its value.
func (mru *MRUEviction) InsertIntoCache(cache *model.Cache, node *model.Node) {
	// Check if key exists in cache.
	mNode, ok := cache.Storage[node.Key]

	if ok {
		// If it does, we update it and make it the most recently used element.
		mru.remove(mNode)
	} else {
		// If not, we check if adding it would exceed cache limit. If it will, remove the
		// most recently used element Else, just insert and make the inserted node most
		// recently used.
		if len(cache.Storage) == cache.Size {
			recent := mru.right.Prev
			mru.remove(recent)
			delete(cache.Storage, recent.Key)
		}
	}
	mru.insert(node)
	cache.Storage[node.Key] = node
}

// GetFromCache Retrieves the value associated with the key in the given cache.
// If the key is not found, returns -1 and an error.
func (mru *MRUEviction) GetFromCache(cache *model.Cache, key int) (int, error) {
	if node, ok := cache.Storage[key]; ok {
		mru.remove(node)
		mru.insert(node)
		return node.Val, nil
	}
	return -1, fmt.Errorf("key %d not in cache", key)
}
