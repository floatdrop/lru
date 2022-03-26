# lru
[![Go Reference](https://pkg.go.dev/badge/github.com/floatdrop/lru.svg)](https://pkg.go.dev/github.com/floatdrop/lru)
[![CI](https://github.com/floatdrop/lru/actions/workflows/ci.yml/badge.svg)](https://github.com/floatdrop/lru/actions/workflows/ci.yml)
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)
[![Go Report Card](https://goreportcard.com/badge/github.com/floatdrop/lru)](https://goreportcard.com/report/github.com/floatdrop/lru)

Thread safe GoLang LRU cache.

## Example

```go
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
```

## TTL

```go
import (
	"fmt"

	"github.com/floatdrop/lru"
)

type Expiring[T any] struct {
	value *T
}

func (E *Expiring[T]) Valid() *T {
	if E == nil {
		return nil
	}

	return E.value
}

func WithTTL[T any](value T, ttl time.Duration) Expiring[T] {
	e := Expiring[T]{
		value: &value,
	}

	time.AfterFunc(ttl, func() {
		e.value = nil // Release memory
	})

	return e
}

func main() {
	l := lru.New[string, Expiring[string]](256)

	l.Set("Hello", WithTTL("Bye", time.Hour))

	if e := l.Get("Hello").Valid(); e != nil {
		fmt.Println(*e)
	}
}
```

## Benchmarks

```
floatdrop/lru:
    BenchmarkLRU_Rand-8   	 9373129	       124.4 ns/op	       8 B/op	       1 allocs/op
    BenchmarkLRU_Freq-8   	10124833	       120.3 ns/op	       8 B/op	       1 allocs/op

hashicorp/golang-lru:
    BenchmarkLRU_Rand-8   	 5992782	       195.8 ns/op	      76 B/op	       3 allocs/op
    BenchmarkLRU_Freq-8   	 6355358	       186.1 ns/op	      71 B/op	       3 allocs/op
```
