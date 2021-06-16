package gset

import (
	"sync"
)

// Gsset interface is the interface
type Gset interface {
	Append(string)
	GetSet() []string
	Lookup(string) bool
	// Union(Gset)
}

// gset satisfies Gset interface{}.
type gset struct {
	// because an empty struct takes no memory.
	set map[string]struct{}
	mu  sync.RWMutex
}

// static check if gset implements Gset
var _ Gset = (*gset)(nil)

// New initializes new Gset interface{} which is the gset implementation, basically whose concrete type is gset struct.
func New() Gset {
	return &gset{
		set: make(map[string]struct{}),
	}
}

// Append appends to the gset, this is concurrent safe.
func (s *gset) Append(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set[key] = struct{}{}
}

// This is much faster
// 1.552 ns/op
// func (s *gset) GetSet() map[string]struct{} {
// 	return s.set
// }

// GetSet returns the set as an array of string.
// Concurrent safe using Read Write Mutex, because if you might be trying to read
// the set, but at the same time some go routine is trying to write in the gset.
// This is costly O(n), where n is the size of the map
func (s *gset) GetSet() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.set) == 0 {
		return nil
	}
	st := make([]string, len(s.set))
	i := 0
	for v := range s.set {
		st[i] = v
		i++
	}
	return st
}

// Lookup returns a boolean for if the value is present in the set already.
// So true for the key exists, and false if the key doesn't exist.
func (s *gset) Lookup(key string) (ok bool) {
	_, ok = s.set[key]
	return
}

// Union takes 2 gsets as an argument, and
// returns the union of both the gsets.
func Union(a Gset, b Gset) Gset {
	c := New()
	for _, v := range a.GetSet() {
		c.Append(v)
	}
	for _, v := range b.GetSet() {
		c.Append(v)
	}
	return c
}
