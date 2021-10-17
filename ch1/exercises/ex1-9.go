// Exericse 1.9: Modify fetch to also print the HTTP status code found in resp.Status.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			s := []string{"http://", url}
			url = strings.Join(s, "")
		}
		resp, err := http.Get(url)
		fmt.Printf("fetch: %v - %s\n", resp.Status, url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		resp.Body.Close()
	}
}
