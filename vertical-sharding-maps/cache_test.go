package main

import (
	"fmt"
	"testing"
)

func TestCache(t *testing.T) {
	cache := NewShardMap(10)

	// this should fail the race condition test (1st commit maps)

	// it will pass the race condition test when we change map to struct
	// for an individual struct the problem is all the operations will pause till the map is locked
	
	// so as a solution for this we will shard the struct that is divide the data into multiple compartments
	// so that even when one is blocked the others can still be used
	// try running using go test ./... -v --race
	for i := range 10 {
		go func(val int) {
			cache.Set(fmt.Sprint(val), val)
		}(i)
	}
}
