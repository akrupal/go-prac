package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	cap int
	l   *list.List
	m   map[int]*list.Element
}

type mapValue struct {
	key   int
	value int
}

func initialiseCache(cap int) *LRUCache {
	return &LRUCache{
		cap: cap,
		l:   list.New(),
		m:   make(map[int]*list.Element),
	}
}

func (lru *LRUCache) Put(k, v int) {
	el, exist := lru.m[k]
	if exist {
		fmt.Println("org value was:", el.Value)
		el.Value = mapValue{k, v}
		lru.m[k] = el
		lru.l.MoveToFront(el)
	} else {
		e := lru.l.PushFront(mapValue{k, v})
		lru.m[k] = e
	}

	if lru.l.Len() > lru.cap {
		elemVal := lru.l.Remove(lru.l.Back())
		e, ok := elemVal.(mapValue)
		if ok {
			fmt.Println("Capacity exceeded removing last element with key", e.key, "and value", e.value)
			delete(lru.m, e.key)
		}
	}
}

func (lru *LRUCache) Get(k int) {
	el, exist := lru.m[k]
	if exist {
		fmt.Println("got:", el)
		lru.l.MoveToFront(el)
	} else {
		fmt.Println("Element not found for key", k)
	}
}

func (lru *LRUCache) DisplayList() {
}

func main() {

	lru := initialiseCache(3)
	lru.Put(1, 21)
	lru.Put(2, 42)
	lru.Put(3, 100)
	lru.Put(4, 56)
	lru.Get(1)
	lru.Get(3)

}
