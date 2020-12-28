package main

// NewInMemoryPlayerStore returns an instance of a PlayerStore
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

// InMemoryPlayerStore returns a store map
type InMemoryPlayerStore struct {
	store map[string]int
}

// RecordWin adds one to player win count
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

// GetPlayerScore returns the score from player name
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}
