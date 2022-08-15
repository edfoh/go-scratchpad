package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
}

/*
https://leetcode.com/problems/lru-cache/

Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

Implement the LRUCache class:

LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
int get(int key) Return the value of the key if the key exists, otherwise return -1.
void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation, evict the least recently used key.
The functions get and put must each run in O(1) average time complexity.
*/
type LRUCache struct {
	list     *list.List
	nodes    map[int]*list.Element
	capacity int
}

type keyValue struct {
	Key   int
	Value int
}

func newKeyValue(key, value int) *keyValue {
	return &keyValue{key, value}
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		list:     list.New(),
		nodes:    map[int]*list.Element{},
		capacity: capacity,
	}
}

func (cache *LRUCache) Get(key int) int {
	res := -1
	if ele, ok := cache.nodes[key]; ok {
		kv := ele.Value.(*keyValue)
		res = kv.Value
		cache.list.MoveToFront(ele)
	}
	return res
}

func (cache *LRUCache) Put(key int, value int) {
	// if key exist, update value and move to front of linked list
	if ele, ok := cache.nodes[key]; ok {
		updateValue(ele, value)
		cache.list.MoveToFront(ele)
	} else {
		// if not found, add new node and update map
		cache.nodes[key] = cache.list.PushFront(newKeyValue(key, value))

		// check if capacity has been reached, if so remove last element
		if cache.list.Len() > cache.capacity {
			cache.removeLast()
		}
	}
}

func updateValue(ele *list.Element, value int) {
	kv := ele.Value.(*keyValue)
	kv.Value = value
}

func (cache *LRUCache) removeLast() {
	lastEle := cache.list.Back()
	lastKv := lastEle.Value.(*keyValue)
	delete(cache.nodes, lastKv.Key)
	cache.list.Remove(lastEle)
}
