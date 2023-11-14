package main

import (
	_ "embed"
	"testing"
)

//go:embed caisson.go
var trigramexample string

func BenchmarkTrigrams(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Trigrams(trigramexample)
	}
}

func BenchmarkTrigramBytes(b *testing.B) {
	trigrams := Trigrams(trigramexample)

	for i := 0; i < b.N; i++ {
		trigrams[i%len(trigrams)].Bytes()
	}
}

func BenchmarkTrigramBytesFast(b *testing.B) {
	trigrams := Trigrams(trigramexample)

	for i := 0; i < b.N; i++ {
		trigrams[i%len(trigrams)].BytesFast()
	}
}

func BenchmarkItemise(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Itemise(Tokenize(trigramexample))
	}
}
