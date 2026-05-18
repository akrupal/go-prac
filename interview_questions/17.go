package main

// given a slice of URL's write code that makes a get call to the URL concurrently and store the status code and response in a map with the URL as the key

import (
	// "fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type ResponseData struct {
	StatusCode int
	Response   string
	Error      string
}

// Job represents a single URL task
type Job struct {
	URL string
}

func worker(
	id int,
	jobs <-chan Job,
	results map[string]ResponseData,
	mu *sync.Mutex,
	wg *sync.WaitGroup,
) {

	defer wg.Done()

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	for job := range jobs {

		resp, err := client.Get(job.URL)
		if err != nil {

			mu.Lock()
			results[job.URL] = ResponseData{
				Error: err.Error(),
			}
			mu.Unlock()

			continue
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {

			mu.Lock()
			results[job.URL] = ResponseData{
				StatusCode: resp.StatusCode,
				Error:      err.Error(),
			}
			mu.Unlock()

			continue
		}

		mu.Lock()
		results[job.URL] = ResponseData{
			StatusCode: resp.StatusCode,
			Response:   string(body),
		}
		mu.Unlock()
	}
}

func fetchURLsWithWorkerPool(urls []string, workerCount int) map[string]ResponseData {

	jobs := make(chan Job, len(urls))

	results := make(map[string]ResponseData)

	var mu sync.Mutex
	var wg sync.WaitGroup

	// Start fixed number of workers
	for i := 1; i <= workerCount; i++ {

		wg.Add(1)

		go worker(
			i,
			jobs,
			results,
			&mu,
			&wg,
		)
	}

	// Send jobs to workers
	for _, url := range urls {
		jobs <- Job{
			URL: url,
		}
	}

	// Close jobs channel
	close(jobs)

	// Wait for all workers to complete
	wg.Wait()

	return results
}

// func main() {

// 	urls := []string{
// 		"https://jsonplaceholder.typicode.com/posts/1",
// 		"https://jsonplaceholder.typicode.com/posts/2",
// 		"https://jsonplaceholder.typicode.com/posts/3",
// 		"https://jsonplaceholder.typicode.com/posts/4",
// 		"https://jsonplaceholder.typicode.com/posts/5",
// 	}

// 	// Instead of 10000 goroutines,
// 	// we only run 5 workers
// 	workerCount := 5

// 	results := fetchURLsWithWorkerPool(
// 		urls,
// 		workerCount,
// 	)

// 	for url, data := range results {

// 		fmt.Println("URL:", url)
// 		fmt.Println("Status Code:", data.StatusCode)

// 		if data.Error != "" {
// 			fmt.Println("Error:", data.Error)
// 		} else {
// 			fmt.Println("Response Length:", len(data.Response))
// 		}

// 		fmt.Println("--------------------------------")
// 	}
// }
