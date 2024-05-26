package lrucache

type node[K comparable, V any] struct {
	key        K
	value      V
	next, prev *node[K, V]
}

func (n *node[K, V]) remove() {
	n.next.prev = n.prev
	n.prev.next = n.next
}

type LRUCache[K comparable, V any] struct {
	capacity   int
	cache      map[K]*node[K, V]
	head, tail *node[K, V]
}

// Create a new LRUCache with the given capacity.
func New[K comparable, V any](capacity int) *LRUCache[K, V] {
	c := &LRUCache[K, V]{
		capacity: capacity,
		cache:    make(map[K]*node[K, V]),
		head:     &node[K, V]{},
		tail:     &node[K, V]{},
	}
	c.head.next = c.tail
	c.tail.prev = c.head
	return c
}

// Get a value from the cache. The second return value is false if the key is not found.
func (c *LRUCache[K, V]) Get(key K) (V, bool) {
	if n, ok := c.cache[key]; ok {
		n.remove()
		c.tail.prev.next = n
		n.prev = c.tail.prev
		c.tail.prev = n
		n.next = c.tail
		return n.value, true
	}
	var v V
	return v, false
}

// Put a key-value pair into the cache. If the key already exists, the value is updated.
func (c *LRUCache[K, V]) Put(key K, value V) {
	if n, ok := c.cache[key]; ok {
		n.remove()
	}
	n := &node[K, V]{key: key, value: value}
	c.cache[key] = n
	c.tail.prev.next = n
	n.prev = c.tail.prev
	c.tail.prev = n
	n.next = c.tail
	if len(c.cache) > c.capacity {
		delete(c.cache, c.head.next.key)
		c.head.next.remove()
	}
}
