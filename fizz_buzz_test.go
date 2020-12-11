package main

import (
	"testing"
)

func TestTableCheckAndAdd(t *testing.T) {
	var tests = []struct {
		index    int
		value    int
		word     string
		expected string
	}{
		{3, 3, "Fizz", "Fizz"},
		{3, 5, "Buzz", ""},
		{5, 3, "Fizz", ""},
		{5, 5, "Buzz", "Buzz"},
		{7, 3, "Fizz", ""},
		{7, 5, "Buzz", ""},
	}

	for _, test := range tests {
		if output := checkAndAdd(test.index, test.value, test.word); output != test.expected {
			t.Errorf("Test Failed: index %d; value %d; expected %s, result %s", test.index, test.value, test.expected, output)
		}
	}
}

func TestTableCreateString(t *testing.T) {
	var tests = []struct {
		input    int
		expected string
	}{
		{3, "Fizz"},
		{5, "Buzz"},
		{7, "7"},
		{15, "FizzBuzz"},
	}

	for _, test := range tests {
		if output := createString(test.input, values); output != test.expected {
			t.Errorf("Test Failed: input of %d; expected %s; received %s", test.input, test.expected, output)
		}
	}
}
