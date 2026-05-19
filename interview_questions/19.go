package main

//lets say I have request that is being made to a server and there to not make the user wait we use async things 1st thing being storing relevant info into DB 2nd being using a broker like Kafka there is one more method using context where even if the parent context ends when the user is sent a 200 OK response the child context can still continue

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func process(ctx context.Context, userID string) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Finished processing for:", userID)

	case <-ctx.Done():
		fmt.Println("Processing cancelled:", ctx.Err())
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	userID := "123"

	// Independent context
	bgCtx, cancel := context.WithTimeout(
		context.Background(),
		30*time.Second,
	)

	// Important
	go func() {
		defer cancel()

		process(bgCtx, userID)
	}()

	// Respond immediately
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Request accepted"))
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}