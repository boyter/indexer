package main

import "testing"

var res []string

/*
goos: darwin
goarch: arm64
pkg: indexer
BenchmarkTokenize-10    	  458174	      2656 ns/op	    6848 B/op	       4 allocs/op
PASS
ok  	indexer	2.477s
*/
func BenchmarkTokenize(b *testing.B) {
	const text = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`
	var result []string
	for n := 0; n < b.N; n++ {
		result = Tokenize(text)
	}
	res = result
}
