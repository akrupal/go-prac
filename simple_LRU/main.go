package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	c int
	l *list.List
	m map[int]*list.Element
}

func initialiseCashe(cap int) LRUCache {
	return LRUCache{
		c: cap,
		l: list.New(),
		m: make(map[int]*list.Element),
	}
}

func (lr *LRUCache) Get(key int) {
	el, ok := lr.m[key]
	if ok {
		val := el.Value
		fmt.Println(val)
		lr.l.PushFront(el)
	} else {
		fmt.Println("Value not found for key ", key)
	}
}

func (lr *LRUCache) Put(key, value int) {
	el, ok := lr.m[key]
	if ok {
		e := lr.l.PushFront(el)
		lr.m[key] = e
	} else {
		e := lr.l.PushFront(value)
		lr.m[key] = e
	}

	if lr.l.Len() > lr.c {
		e := lr.l.Back()
		a, ok := e.Value.(int)
		if ok {
			fmt.Println("More than capacity removing last element with value ", a)
		}
		lr.l.Remove(e)
		delete(lr.m, a/10)
	}
}

func main() {
	cache := initialiseCashe(3)
	cache.Put(1, 10)
	cache.Put(2, 20)
	cache.Put(3, 30)
	cache.Put(4, 40)

	cache.Get(1)
	cache.Get(4)
}
