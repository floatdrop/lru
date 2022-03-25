package lru_test

import (
	"fmt"

	"github.com/floatdrop/lru"
)

func ExampleLRU() {
	cache := lru.New[string, int](256)

	cache.Set("Hello", 5)

	fmt.Println(*cache.Get("Hello"))
	// Output: 5
}
