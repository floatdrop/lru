package lru

import "testing"

// Nop should be compatible with Cache interface.
var _ Cache[int, int] = &Nop[int, int]{}

func TestNop(t *testing.T) {
	l := &Nop[int, int]{}

	if e := l.Set(5, 5); e == nil || *e != 5 {
		t.Fatal("value should be evicted")
	}

	if l.Get(5) != nil {
		t.Fatal("no values should be in nop cache")
	}
}
