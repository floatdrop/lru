package lru

// Cache defines minimal interface for all cache implementations, that requires two methods: Get and Set.
type Cache[K comparable, V any] interface {
	Get(key K) *V

	Set(key K, value V) *V
}
