# LRU Cache
Simple implementation of a LRU cache using generics in Go.

## Usage
```go
package main

import "github.com/adriansahlman/lrucache"

func main() {
    // Make a cache with a capacity of up to 3 values
	c := lrucache.New[string, int](3)

    // Fill the cache
	c.Put("a", 1)
	c.Put("b", 2)
	c.Put("c", 3)

    // We should now be able to get a value from the cache
    // that we previously stored.
	v, ok := c.Get("a")
	if !ok {
		panic("expected key 'a' to be in cache")
	}
	if v != 1 {
		panic("expected value of key 'a' to be 1")
	}

	// by fetching the value for "a" we move it to the end
	// of the queue, so it should no longer be the next
	// value to be evicted. Instead we can now expect "b"
	// to be the next value to be evicted when adding a new value.
	c.Put("d", 4)
	_, ok = c.Get("b")
	if ok {
		panic("expected key 'b' to be evicted")
	}
	_, ok = c.Get("a")
	if !ok {
		panic("expected key 'a' to be in cache")
	}
}
```
