package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "Hello Mutiara",
			expected: []string{
				"hello",
				"mutiara",
			},
		},
		{
			input: "Hello Varras",
			expected: []string{
				"hello",
				"varras",
			},
		},
		{
			input: "Electrical Engineering Unsoed",
			expected: []string{
				"electrical",
				"engineering",
				"unsoed",
			},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("The lenghts are not equal: %v vs %v",
				len(actual),
				len(c.expected),
			)
			continue
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := c.expected[i]
			if actualWord != expectedWord {
				t.Errorf("The words are not equal: %v vs %v", actualWord, expectedWord)
			}
		}

	}
}
