package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	lru := NewLRUCache(2)
	lru.Put(1, 1)
	lru.Put(2, 2) // cache is {1=1, 2=2}
	assert.Equal(t, 1, lru.Get(1))
	lru.Put(3, 3) // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	assert.Equal(t, -1, lru.Get(2))
	lru.Put(4, 4) // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	assert.Equal(t, -1, lru.Get(1))
	assert.Equal(t, 3, lru.Get(3))
	assert.Equal(t, 4, lru.Get(4))
}
