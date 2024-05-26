package lrucache_test

import (
	"testing"

	"github.com/adriansahlman/lrucache"
)

func TestLRUCache(t *testing.T) {
	c := lrucache.New[string, int](3)
	putWithTestGet := func(t *testing.T, k string, v int) {
		t.Helper()
		c.Put(k, v)
		vv, ok := c.Get(k)
		if !ok {
			t.Fatalf("expected key '%s' to be in cache", k)
		}
		if vv != v {
			t.Fatalf("expected value %d, got %d", v, vv)
		}
	}
	putWithTestGet(t, "a", 1)
	putWithTestGet(t, "b", 2)
	putWithTestGet(t, "c", 3)
	// "a" should be evicted
	putWithTestGet(t, "d", 4)
	v, ok := c.Get("a")
	if ok {
		t.Fatalf("expected key 'a' to be evicted, but got value %d", v)
	}
	// evict "c" instead of "b" by first fetching "b"
	_, ok = c.Get("b")
	if !ok {
		t.Fatalf("expected key 'b' to be in cache")
	}
	putWithTestGet(t, "e", 5)
	_, ok = c.Get("b")
	if !ok {
		t.Fatalf("expected key 'b' to be in cache")
	}
	v, ok = c.Get("c")
	if ok {
		t.Fatalf("expected key 'c' to be evicted, but got value %d", v)
	}
}

func Example() {
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
