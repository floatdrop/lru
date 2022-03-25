# lru
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)

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

## Benchmarks

```
floatdrop/lru:
    BenchmarkLRU_Rand-8   	 9373129	       124.4 ns/op	       8 B/op	       1 allocs/op
    BenchmarkLRU_Freq-8   	10124833	       120.3 ns/op	       8 B/op	       1 allocs/op

hashicorp/golang-lru:
    BenchmarkLRU_Rand-8   	 5992782	       195.8 ns/op	      76 B/op	       3 allocs/op
    BenchmarkLRU_Freq-8   	 6355358	       186.1 ns/op	      71 B/op	       3 allocs/op
```
