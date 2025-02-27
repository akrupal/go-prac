package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// RateLimiter struct encapsulates rate limiting logic
type RateLimiter struct {
	maxRequests    int
	interval       time.Duration
	mu             sync.Mutex
	clientRequests map[string][]time.Time
}

// NewRateLimiter initializes a new rate limiter
func NewRateLimiter(maxRequests int, interval time.Duration) *RateLimiter {
	return &RateLimiter{
		maxRequests:    maxRequests,
		interval:       interval,
		clientRequests: make(map[string][]time.Time),
	}
}

// LimitMiddleware applies rate limiting to any HTTP handler
func (rl *RateLimiter) LimitMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rl.mu.Lock()
		defer rl.mu.Unlock()

		clientIP := r.RemoteAddr // Identifying user by IP (better: use API keys or user ID)
		fmt.Println(clientIP)
		now := time.Now()

		// Get or initialize request timestamps for this client
		requests, exists := rl.clientRequests[clientIP]
		fmt.Println(requests)
		fmt.Println(exists)
		if !exists {
			requests = []time.Time{}
		}

		// Remove old requests outside the interval
		var newRequests []time.Time
		for _, t := range requests {
			if now.Sub(t) <= rl.interval {
				newRequests = append(newRequests, t)
			}
		}

		// Apply rate limiting
		if len(newRequests) >= rl.maxRequests {
			http.Error(w, "Excedded the maximum allowed number of requests", http.StatusTooManyRequests)
			return
		}

		// Add new request timestamp
		newRequests = append(newRequests, now)
		fmt.Println(newRequests)
		rl.clientRequests[clientIP] = newRequests

		// Proceed to the actual handler
		next(w, r)
	}
}

// helloWorld is a sample handler function
func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func main() {
	// Initialize rate limiter with 2 requests per 5 seconds
	limiter := NewRateLimiter(2, 5*time.Second)

	http.HandleFunc("/hello", limiter.LimitMiddleware(helloWorld))

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
