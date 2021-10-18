package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%gF = %gC, freezing\n", freezingF, fToC(freezingF))
	fmt.Printf("%gF = %gC, boiling\n", boilingF, fToC(boilingF))
	// Output:
	// 32F = 0C, freezing
	// 212F = 100C, boiling
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
