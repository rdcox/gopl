package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%2.fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
	}
	nbytes, err := io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading: %s %v", url, err)
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%2.fs %7d %s", secs, nbytes, url)
}

// Queried: https://httpstat.us/524
// Returned 0 bytes
// Processed for 0 seconds
// Couldn't think of a good way to replicate a website *actually* not responding
// My guess would be that it hangs our program until we get a response or timeout
