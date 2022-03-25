package lru

// Nop implements Cache interface, but does nothing inside (all sets are ignored, all get operations return nil).
type Nop[K comparable, V any] struct{}

// Get method of Nop cache always returns nil.
func (n *Nop[K, V]) Get(key K) *V {
	return nil
}

// Set method of Nop cache always returns value pointer (as if it was immediately evicted).
func (n *Nop[K, V]) Set(key K, value V) *V {
	return &value
}
