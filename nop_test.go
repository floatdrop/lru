package lru

// Nop should be compatible with Cache interface.
var _ Cache[int, int] = &Nop[int, int]{}
