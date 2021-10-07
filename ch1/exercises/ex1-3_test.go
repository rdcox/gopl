package main

import (
	"strings"
	"testing"
)

var args = []string{"arg1", "arg1.0", "arg3", "arg2", "arg3", "arg4", "arg5bitlongerthantherest", "more", "and", "more", "and", "arg2", "arg3", "arg4", "arg5bitlongerthantherest"}

// echo1
func concat1(args []string) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	// fmt.Println(s)
}

// echo2
func concat2(args []string) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	// fmt.Println(s)
}

// echo3
func join1(args []string) {
	strings.Join(args, " ")
}

func BenchmarkConcat1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concat1(args)
	}
}

func BenchmarkConcat2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concat2(args)
	}
}

func BenchmarkJoin1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		join1(args)
	}
}
