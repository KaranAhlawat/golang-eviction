# Cache Eviction Policy

## Run Demo
To run the build on Unix systems (Linux, MacOS, FreeBDS etc), run the following command from the shell inside the project directory root
```
./build/cache
```

## Problem Statement

Design a Cache with 2 eviction policies, Least Recently Used (LRU) and Most Recently Used (MRU)

### Requirements

1. The cache should be extendable to accommodate other strategies
2. Configurable cache size while instantiating

## Assumptions

1. It is assumed that both the keys and the values in the cache are integers.

## Design

1. `Cache` - The main struct, that will actually be used by the third party. Contains the size, the cache policy and the cache contents (storage).
2. `EvictionStrategy` - Interface to define eviction strategies/policies for the cache. It has full access to the cache, and is used to provide support for any other strategy the user may want to implement. They just need to implement the interface, and code the correct logic.
3. `LRUStrategy` - One of the two included strategies for the cache.
4. `MRUStrategy` - One of the two included strategies for the cache.

## Usage API

1. `Insert(key, val)` - Inserts the key value pair into the cache.
2. `Get(key)` - Gets the value of the key from the cache, or returns -1 and an error value if the key was not found in the cache.
3. `StateOfCache()` - Returns a map of key to value.

## Adding Strategies

To add a strategy for the cache, one must implement the `EvictionStrategy` interface. It has two methods :-
1. `InsertIntoCache`
2. `GetFromCache`

The work in the actual `Cache` API is forwarded to the eviction strategy of the cache, so that it can keep track of the cache state as required, without having to modify the cache struct itself.
