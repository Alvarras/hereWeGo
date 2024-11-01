package client

import (
	"fmt"
	"testing"
)

func getMailtoLinkForEmail(email string) string {
	return "mailto:" + email
}

func TestParseEmail(t *testing.T) {
	type testCase struct {
		email    string
		expected string
	}

	tests := []testCase{
		{"wayne.lagner@dev.boot", "mailto:wayne.lagner@dev.boot"},
		{"heckmann@what.de", "mailto:heckmann@what.de"},
		{"a.liar@pants.fire", "mailto:a.liar@pants.fire"},
	}

	// Additional test cases for submission
	if withSubmit {
		tests = append(tests, []testCase{
			{"", "mailto:"},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		output := getMailtoLinkForEmail(test.email)
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Email:		%v
Expecting:  %v
Actual:     %v
Fail`, test.email, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Email:		%v
Expecting:  %v
Actual:     %v
Pass
`, test.email, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
