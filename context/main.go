package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Work done") //this will only be executed if timeout time is greater than the time set above
	case <-ctx.Done():
		fmt.Println("Cancelled:", ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go doWork(ctx)

	//wait for 3 sec to allow processing to complete
	time.Sleep(3 * time.Second)
}
