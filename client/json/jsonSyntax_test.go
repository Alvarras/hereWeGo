package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

const issueList = `
	[
	{
		"id": 0,
		"name": "Fix the thing",
		"estimate": 0.5,
		"completed": false
	},
	{
		"id": 1,
		"name": "Unstick the widget",
		"estimate": 30,
		"completed": false
	}
	]

`

const userObject = `
{
	"name": "Wayne Lagner",
	"role": "Developer",
	"remote": true
}
`

func isValidJSON(input string) bool {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(input), "", "  ")
	return err == nil
}

func TestIsValidJSON(t *testing.T) {
	type testCase struct {
		input string
	}
	tests := []testCase{
		{issueList},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{userObject},
		}...)
	}

	for _, test := range tests {
		if output := isValidJSON(test.input); !output {
			t.Errorf(`Test Failed. Input:
%v
  =>
expected isValidJSON: %v
actual isValidJSON: %v
`,
				test.input, true, output)
		} else {
			fmt.Printf(`Test Passed. Input:
%v
  =>
expected isValidJSON: %v
actual isValidJSON: %v
`,
				test.input, true, output)
		}
		fmt.Println("==============================")
	}
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
