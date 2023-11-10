package main

import (
	"reflect"
	"testing"
)

func TestTrigrams(t *testing.T) {
	for _, testCase := range []struct {
		text     string
		expected []string
	}{
		{"", []string{}},
		{"ab", []string{}},
		{"abc", []string{"abc"}},
		{"ABC", []string{"ABC"}},
		{"abcd", []string{"abc", "bcd"}},
		{"a b c", []string{"a b", " b ", "b c"}},
	} {
		testCase := testCase
		t.Run(testCase.text, func(t *testing.T) {
			result := Trigrams(testCase.text)
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("Trigrams(%v)=%v, expected %v", testCase.text, result, testCase.expected)
			}
		})
	}
}

func TestTokenize(t *testing.T) {
	for _, testCase := range []struct {
		text     string
		expected []string
	}{
		{"", []string{}},
		{"ab", []string{}},
		{"abc", []string{"abc"}},
		{"ABC", []string{"abc"}},
		{"AbCd", []string{"abc", "bcd"}},

		// Multi-word tests
		{"abc def", []string{"abc", "def"}},
		{"abcd efg", []string{"abc", "bcd", "efg"}},
		{"abc abc", []string{"abc", "abc"}},
		{"abc de", []string{"abc"}},
		{"abc de fghi", []string{"abc", "fgh", "ghi"}},
	} {
		testCase := testCase
		t.Run(testCase.text, func(t *testing.T) {
			result := Tokenize(testCase.text)
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("Tokenize(%v)=%v, expected %v", testCase.text, result, testCase.expected)
			}
		})
	}
}
