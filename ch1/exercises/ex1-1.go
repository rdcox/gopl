// Exercise 1.1: Modify the echo program to also print os.Args[0], the name of the command that invoked it
package main

import (
	"fmt"
	"os"
	"strings"
)

// Solution
func main() {
	// could also just fmt.Println(os.Args) - this cuts out the []
	fmt.Println(strings.Join(os.Args, " "))
}
