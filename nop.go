package lru

// Nop implements Cache interface, but does nothing inside (all sets are ignored, all get operations return nil).
type Nop[K comparable, V any] struct{}

// Get method of Nop cache always returns nil.
func (n *Nop[K, V]) Get(key K) *V {
	return nil
}

// Set method of Nop cache always returns value pointer (as if it was immediately evicted).
func (n *Nop[K, V]) Set(key K, value V) *Evicted[K, V] {
	return &Evicted[K, V]{key, value}
}

// Len method of Nop cache always returns 0.
func (n *Nop[K, V]) Len() int {
	return 0
}

// Remove method of Nop cache always returns nil.
func (n *Nop[K, V]) Remove(key K) *V {
	return nil
}

// Peek method of Nop cache always returns nil.
func (n *Nop[K, V]) Peek(key K) *V {
	return nil
}
