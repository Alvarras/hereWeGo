package dasar

import (
	"strconv"
	"testing"
)

func fizzbuzz(n int) string {
	if n%15 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	} else {
		return strconv.Itoa(n)
	}
}

func TestFizzbuzz(t *testing.T) {
	type testCase struct {
		n        int
		expected string
	}
	tests := []testCase{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{4, "4"},
		{5, "Buzz"},
		{6, "Fizz"},
		{7, "7"},
		{8, "8"},
		{9, "Fizz"},
		{10, "Buzz"},
		{11, "11"},
		{12, "Fizz"},
		{13, "13"},
		{14, "14"},
		{15, "FizzBuzz"},
		{16, "16"},
		{17, "17"},
		{18, "Fizz"},
		{19, "19"},
		{20, "Buzz"},
	}

	for _, test := range tests {
		output := fizzbuzz(test.n)
		if output == test.expected {
			t.Logf("SUCCESS: For %d, got expected %s", test.n, output)
		} else {
			t.Errorf("FAILED: For %d, expected %s, but got %s", test.n, test.expected, output)
		}
	}
}
