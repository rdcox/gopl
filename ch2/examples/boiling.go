package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point: %gF or %gC\n", f, c)
	// Output:
	// boiling point: 212.0F or 100C
}
