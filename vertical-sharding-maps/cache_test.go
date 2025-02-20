package main

import (
	"fmt"
	"testing"
)

func TestCache(t *testing.T) {
	cache := NewCache()

	// this should fail the race condition test
	// try running using go test ./... -v --race
	for i := range 10 {
		go func(val int) {
			cache.Set(fmt.Sprint(val), val)
		}(i)
	}
}
