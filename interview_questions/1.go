package main

import (
	"fmt"
	"sync"
)

func main() {
	//write a code that takes in 2 go routines
	// one prints the odd number and the other prints even numbers
	// print numbers from 1 to 10 sequentially by making the go routines communicate
	oddCh := make(chan bool)
	evenCh := make(chan bool)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 1; i < 11; i += 2 {
			<-oddCh
			fmt.Println(i)
			evenCh <- true
		}
		wg.Done()
	}()
	go func() {
		for i := 2; i < 11; i += 2 {
			<-evenCh
			fmt.Println(i)
			if i < 10 {
				oddCh <- true
			}
		}
		wg.Done()
	}()
	oddCh <- true
	wg.Wait()

	close(oddCh)
	close(evenCh)
}
