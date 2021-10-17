// Exercise 1.10: Find a web site that produces a large amount of data. Investigate caching by
// running fetchall twice in succession to see whether the reported time changes much. Do
// you get the same content each time? Modify fetchall to print its output to a file so it
// can be examined.
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
	out, err := os.Create("response.txt")
	if err != nil {
		ch <- fmt.Sprint("creating file")
	}
	nbytes, err := io.Copy(out, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading: %s %v", url, err)
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%2.fs %7d %s", secs, nbytes, url)
}

// Ran against https://reddit.com
// Trial 1: 3s 740282 bytes
// Trial 2: 2s 739540 bytes
// Trial 3: 5s 740103 bytes
// Trial 4: 2s 739719 bytes
// Trial 5: 6s 740430 bytes
// Differences in size is likely due to the dynamic nature of the webpage.
// It's size fluctuates as its content changes.
// Differences in execution time is due, partially to changes in size, but
// also due to other undeterministic networking factors, i.e. routing, traffic, etc.
