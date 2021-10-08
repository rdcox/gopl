// Exercise: 1.4: Modify dup2 to print the names of all files in which each duplicated line occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	dup2()
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~")
	dup2_modified()
}

func dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func dup2_modified() {
	counts := make(map[string]int)
	foundFiles := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2_modified: %v\n", err)
			}
			countLinesAndFiles(f, counts, foundFiles)
			f.Close()
		}
	}
	fmt.Printf("Line\t# of Duplicates\n")
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("\n%s\t%d\n", line, n)
			fmt.Println("Found in files:")
			for _, f := range foundFiles[line] {
				fmt.Printf("\t%s\n", f)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

func countLinesAndFiles(f *os.File, counts map[string]int, foundFiles map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		foundFiles[input.Text()] = append(foundFiles[input.Text()], f.Name())
	}
}
