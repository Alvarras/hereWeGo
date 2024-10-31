package client

import (
	"encoding/json"
	"fmt"
	"testing"
)

func marshalAll[T any](items []T) ([][]byte, error) {
	result := make([][]byte, len(items))

	for i, item := range items {
		marshaled, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}
		result[i] = marshaled
	}
	return result, nil
}

type JelloUser struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type JelloBoard struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	TaskCount   int    `json:"taskCount"`
}

func TestEncoding(t *testing.T) {
	type testCase struct {
		inputs   []any
		expected [][]byte
	}

	tests := []testCase{
		{
			inputs: []any{
				JelloUser{"001", "Sir Bedevere the Wise", "Scientist"},
				JelloUser{"002", "Sir Lancelot the Brave", "Knight"},
				JelloBoard{"Quest for the Holy Grail", "Tasks related to finding the Grail", 5},
			},
			expected: [][]byte{
				[]byte(`{"id":"001","name":"Sir Bedevere the Wise","role":"Scientist"}`),
				[]byte(`{"id":"002","name":"Sir Lancelot the Brave","role":"Knight"}`),
				[]byte(`{"name":"Quest for the Holy Grail","description":"Tasks related to finding the Grail","taskCount":5}`),
			},
		},
		{
			inputs: []any{
				JelloUser{"003", "Sir Galahad the Pure", "Knight"},
				JelloBoard{"Defeat the Killer Rabbit", "Prepare for battle with the Rabbit of Caerbannog", 7},
				JelloBoard{"Use the Holy Hand Grenade", "Instructions on deploying the Holy Hand Grenade of Antioch", 3},
			},
			expected: [][]byte{
				[]byte(`{"id":"003","name":"Sir Galahad the Pure","role":"Knight"}`),
				[]byte(`{"name":"Defeat the Killer Rabbit","description":"Prepare for battle with the Rabbit of Caerbannog","taskCount":7}`),
				[]byte(`{"name":"Use the Holy Hand Grenade","description":"Instructions on deploying the Holy Hand Grenade of Antioch","taskCount":3}`),
			},
		},
	}
	if withSubmit {
		tests = append(tests, testCase{
			inputs: []any{
				JelloUser{"004", "Sir Robin the Not-Quite-So-Brave-As-Sir-Lancelot", "Minstrel"},
				JelloBoard{"Avoid Dangerous Situations", "Strategies for running away bravely", 2},
			},
			expected: [][]byte{
				[]byte(`{"id":"004","name":"Sir Robin the Not-Quite-So-Brave-As-Sir-Lancelot","role":"Minstrel"}`),
				[]byte(`{"name":"Avoid Dangerous Situations","description":"Strategies for running away bravely","taskCount":2}`),
			},
		})
	}

	// Running the tests
	for _, test := range tests {
		output, err := marshalAll(test.inputs)
		if err != nil {
			t.Errorf(`Test Failed: %v
  unexpected error: %v
`, test.inputs, err)
			continue
		}
		if len(output) != len(test.expected) {
			t.Errorf(`Test Failed: %v
  expected length: %v
  actual length:   %v
`, test.inputs, len(test.expected), len(output))
			continue
		}
		for i, jsonOutput := range output {
			if string(jsonOutput) != string(test.expected[i]) {
				t.Errorf(`Test Failed at index %d:
  input:    %v
  expected: %s
  actual:   %s
`, i, test.inputs[i], test.expected[i], jsonOutput)
			} else {
				fmt.Printf(`Test Passed at index %d:
  input:    %v
  expected: %s
  actual:   %s
`, i, test.inputs[i], test.expected[i], jsonOutput)
			}
		}
		fmt.Println("==============================")
	}
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
