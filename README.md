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

func main() {
	cache := lru.New[string, int](256)

	cache.Set("Hello", 5)

	if e := cache.Get("Hello"); e != nil {
		fmt.Println(*e)
		// Output: 5
	}
}
```

## TTL

You can wrap values into `Expiring[T any]` struct to release memory on timer (or manually in `Valid` method).

<details>
    <summary>Example implementation</summary>

```go
import (
    "fmt"
    "time"

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

**Note:** Althou this short implementation frees memory after ttl duration, it will not erase entry for key in cache. It can be a problem, if you do not check nillnes after getting element from cache and call `Set` afterwards.
</details>

## Benchmarks

```
floatdrop/lru:
    BenchmarkLRU_Rand-8   	 8802915	       131.7 ns/op	      24 B/op	       1 allocs/op
    BenchmarkLRU_Freq-8   	 9392769	       127.8 ns/op	      24 B/op	       1 allocs/op

hashicorp/golang-lru:
    BenchmarkLRU_Rand-8   	 5992782	       195.8 ns/op	      76 B/op	       3 allocs/op
    BenchmarkLRU_Freq-8   	 6355358	       186.1 ns/op	      71 B/op	       3 allocs/op

jellydator/ttlcache:
    BenchmarkLRU_Rand-8   	 4447654	       253.5 ns/op	     144 B/op	       2 allocs/op
    BenchmarkLRU_Freq-8   	 4837938	       240.9 ns/op	     137 B/op	       2 allocs/op
```
