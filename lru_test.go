package lru

import (
	"math/rand"
	"testing"
)

// LRU should be compatible with Cache interface.
var _ Cache[int, int] = &LRU[int, int]{}

func BenchmarkLRU_Rand(b *testing.B) {
	l := New[int64, int64](8192)

	trace := make([]int64, b.N*2)
	for i := 0; i < b.N*2; i++ {
		trace[i] = rand.Int63() % 32768
	}

	b.ResetTimer()

	var hit, miss int
	for i := 0; i < 2*b.N; i++ {
		if i%2 == 0 {
			l.Set(trace[i], trace[i])
		} else {
			if l.Get(trace[i]) == nil {
				miss++
			} else {
				hit++
			}
		}
	}
	b.Logf("hit: %d miss: %d ratio: %f", hit, miss, float64(hit)/float64(miss))
}

func BenchmarkLRU_Freq(b *testing.B) {
	l := New[int64, int64](8192)

	trace := make([]int64, b.N*2)
	for i := 0; i < b.N*2; i++ {
		if i%2 == 0 {
			trace[i] = rand.Int63() % 16384
		} else {
			trace[i] = rand.Int63() % 32768
		}
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		l.Set(trace[i], trace[i])
	}
	var hit, miss int
	for i := 0; i < b.N; i++ {
		if l.Get(trace[i]) == nil {
			miss++
		} else {
			hit++
		}
	}
	b.Logf("hit: %d miss: %d ratio: %f", hit, miss, float64(hit)/float64(miss))
}

func TestLRU_zero(t *testing.T) {
	l := New[int, int](0)
	i := 5

	if e := l.Set(i, i); e == nil || *e != i {
		t.Fatalf("value should be evicted")
	}

	if e := l.Remove(i); e != nil {
		t.Fatalf("value should not be removed")
	}
}

func TestLRU_defaultkey(t *testing.T) {
	l := New[string, int](1)
	var k string
	v := 10

	if e := l.Set(k, v); e != nil {
		t.Fatalf("value should not be evicted")
	}

	if e := l.Get(k); e == nil || *e != v {
		t.Fatalf("bad returned value: %v != %v", e, v)
	}
}

func TestLRU_setget(t *testing.T) {
	l := New[int, int](128)

	if e := l.Get(5); e != nil {
		t.Fatalf("bad returned value: %v != nil", e)
	}

	if l.Set(5, 10) != nil {
		t.Fatal("should not have evictions")
	}

	if e := l.Get(5); *e != 10 {
		t.Fatalf("bad returned value: %v != %v", *e, 10)
	}

	if e := l.Set(5, 9); e == nil || *e != 10 {
		t.Fatal("old value should be evicted")
	}
}

func TestLRU_eviction(t *testing.T) {
	l := New[int, int](128)

	evictCounter := 0
	for i := 0; i < 256; i++ {
		if l.Set(i, i) != nil {
			evictCounter++
		}
	}

	if l.Len() != 128 {
		t.Fatalf("bad len: %v", l.Len())
	}

	if evictCounter != 128 {
		t.Fatalf("bad evict count: %v", evictCounter)
	}

	for i := 0; i < 128; i++ {
		if e := l.Get(i); e != nil {
			t.Fatalf("should be evicted")
		}
	}

	for i := 128; i < 256; i++ {
		if e := l.Get(i); e == nil {
			t.Fatalf("should not be evicted")
		}
	}

	for i := 128; i < 192; i++ {
		l.Remove(i)
		if e := l.Get(i); e != nil {
			t.Fatalf("should be deleted")
		}
	}
}

func TestLRU_peek(t *testing.T) {
	l := New[int, int](2)

	l.Set(1, 1)
	l.Set(2, 2)
	if v := l.Peek(1); v == nil || *v != 1 {
		t.Errorf("1 should be set to 1: %v,", v)
	}

	l.Set(3, 3)
	if l.Peek(1) != nil {
		t.Errorf("should not have updated recent-ness of 1")
	}
}
