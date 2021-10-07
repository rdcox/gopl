package main

import (
	"fmt"
	"os"
)

// Has [] around the ends of the string because default slice formatting
func main() {
	fmt.Println(os.Args[1:])
}
