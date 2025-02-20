package main

import (
	"fmt"
	"testing"
)

func TestCache(t *testing.T) {
	cache := NewCache()

	// this should fail the race condition test
	// it will pass the race condition test when we change map to struct
	// try running using go test ./... -v --race
	for i := range 10 {
		go func(val int) {
			cache.Set(fmt.Sprint(val), val)
		}(i)
	}
}
