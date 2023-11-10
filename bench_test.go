package main

import "testing"

func BenchmarkTrigrams(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Trigrams("this is a test")
	}
}

func BenchmarkTokenize(b *testing.B) {
	const text = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`
	var result []Trigram
	for n := 0; n < b.N; n++ {
		result = Tokenize(text)
	}
	_ = result
}
