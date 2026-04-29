// asked to create a program which uses 2 go routines which prints red and green alternatively infinitely till it gets a termination signal from user
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func print() {
	greenCh := make(chan struct{})
	redCh := make(chan struct{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case <-redCh:
				fmt.Println("red")
				select {
				case greenCh <- struct{}{}:
				case <-ctx.Done():
					return
				}
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case <-greenCh:
				fmt.Println("green")
				select {
				case redCh <- struct{}{}:
				case <-ctx.Done():
					return
				}
			}
		}
	}()

	redCh <- struct{}{}

	<-sigChan
	fmt.Println("\nShutting down gracefully...")

	cancel()

	wg.Wait()

	fmt.Println("Exited cleanly")
}

// func main() {
// 	print()
// }
