package main

import (
	_ "embed"
	"testing"
)

//go:embed caisson.go
var trigramexample string

func BenchmarkTrigrams(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Trigrams(trigramexample)
	}
}

func BenchmarkTrigramsMerovius(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TrigramsMerovius(trigramexample)
	}
}

func BenchmarkTrigramsDancantos(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TrigramsDancantos(trigramexample)
	}
}

func BenchmarkTrigramsJamesrom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TrigramsJamesrom(trigramexample)
	}
}

func BenchmarkTrigramsFfmiruz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TrigramsFfmiruz(trigramexample)
	}
}
