package model

// EvictionStrategy Represents the behaviour of an eviction strategy. Each
// eviction strategy the user wishes to support must follow this interface. The
// insertion and retrieval are offloaded to the strategy itself so that it can do
// any housekeeping operations it needs to without needing to modify the cache
// structure itself.
type EvictionStrategy interface {
	InsertIntoCache(*Cache, *Node)
	GetFromCache(*Cache, int) (int, error)
}
