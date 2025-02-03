package main

import (
	"fmt"
	"sync"
)

func calcSuminParallel(m, n int) {
	sumChan := make(chan int, 3)
	var finalSum int = 0
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func() {
		sum := 0
		for i := m; i < m+(n-m)/3; i++ {
			sum = sum + i
		}
		sumChan <- sum
		wg.Done()
	}()
	go func() {
		sum := 0
		for i := m + (n-m)/3; i < m+(n-m)*2/3; i++ {
			sum = sum + i
		}
		sumChan <- sum
		wg.Done()
	}()
	go func() {
		sum := 0
		for i := m + (n-m)*2/3; i < n+1; i++ {
			sum = sum + i
		}
		sumChan <- sum
		wg.Done()
	}()
	wg.Wait()
	close(sumChan)
	for i := range sumChan {
		finalSum = finalSum + i
	}
	fmt.Println("Final sum is: ", finalSum)
}

// func main() {
// 	calcSuminParallel(1, 10)
// }
