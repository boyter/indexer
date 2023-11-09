package main

import (
	_ "embed"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

//go:embed main.go
var main_go string

func BenchmarkTrigrams(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Trigrams(main_go)
	}
}

func TestTrigrams(t *testing.T) {
	tcs := []struct {
		in   string
		want []string
	}{
		{"", nil},
		{"a", nil},
		{"ab", nil},
		{"abc", []string{"abc"}},
		{"abcd", []string{"abc", "bcd"}},
	}
	for _, tc := range tcs {
		got := Trigrams(tc.in)
		if !cmp.Equal(got, tc.want, cmpopts.EquateEmpty()) {
			t.Errorf("Trigrams(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}
