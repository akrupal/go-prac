package main

import (
	"fmt"
	"sync"
)

func correct() {
	ch := make(chan *int, 4) //this was not buffered at the start
	array := []int{1, 2, 3, 4}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for _, value := range array {
			ch <- &value
		}
		close(ch) //this shouldnt be here usually a buffered channel is closed just before we read the values out of it
	}()
	go func() {
		for value := range ch {
			fmt.Println(*value) //what will be printed here
		}
		wg.Done()
	}()
	wg.Wait()
}

// func main() {

// 	//find mistakes and correct the code
// 	// a very weird code that runs because we are using buffered channels
// 	// but this is not the correct way to write code as there are 2 go routines there should be 2 in Add
// 	correct()
// }
