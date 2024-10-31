package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

type Issue struct {
	Title    string `json:"title"`
	Estimate int    `json:"estimate"`
}

func getIssues(url string) ([]Issue, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	var issues []Issue
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&issues); err != nil {
		return nil, fmt.Errorf("error decoding response body: %w", err)
	}
	return issues, nil
}

func TestDecoding(t *testing.T) {
	type testCase struct {
		url      string
		expected []Issue
	}

	tests := []testCase{
		{
			"https://api.boot.dev/v1/courses_rest_api/learn-http/issues?limit=1",
			[]Issue{{Title: "Fix that one bug nobody understands", Estimate: 19}},
		},
		{
			"https://api.boot.dev/v1/courses_rest_api/learn-http/issues?limit=2",
			[]Issue{
				{Title: "Fix that one bug nobody understands", Estimate: 19},
				{Title: "Implement user authentication flow", Estimate: 6},
			},
		},
	}

	// Additional test cases for submission
	if withSubmit {
		tests = append(tests, []testCase{
			{"", nil},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		issues, _ := getIssues(test.url)

		if !reflect.DeepEqual(issues, test.expected) {
			failCount++
			t.Errorf(`---------------------------------
URL:		%v
Expecting:  %+v
Actual:     %+v
Fail`, test.url, test.expected, issues)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
URL:		%v
Expecting:  %+v
Actual:     %+v
Pass
`, test.url, test.expected, issues)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
