package client

import (
	"fmt"
	"net/url"
	"testing"
)

func getDomainNameFromURL(rawURl string) (string, error) {
	parsedURL, err := url.Parse(rawURl)
	if err != nil {
		return "", fmt.Errorf("failed to parse URL: %w", err)
	}
	return parsedURL.Hostname(), nil
}

func TestParsing(t *testing.T) {
	type testCase struct {
		address   string
		expectErr bool
		expected  string
	}

	tests := []testCase{
		{"https://boot.dev/learn/learn-python", false, "boot.dev"},
		{"https://youtube.com", false, "youtube.com"},
	}

	// Additional test cases for submission
	if withSubmit {
		tests = append(tests, []testCase{
			{"://example.com", true, ""},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		output, err := getDomainNameFromURL(test.address)

		if test.expectErr && (err != nil) != test.expectErr {
			failCount++
			t.Errorf(`---------------------------------
URL:		%v
ExpectErr:	%v
GotErr:		%v
Fail`, test.address, test.expectErr, err != nil)
		} else if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
URL:		%v
Expecting:	%v
Actual:		%v
Fail`, test.address, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
URL:		%v
Expecting:  %v
Actual:     %v
Pass
`, test.address, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
