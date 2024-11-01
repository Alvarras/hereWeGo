package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
)

const domain = "api.boot.dev"

type Issue struct {
	Title string
}

func getIssues(domain string) ([]Issue, error) {
	res, err := http.Get("https://" + domain + "/v1/courses_rest_api/learn-http/issues")
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	var issues []Issue
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&issues)
	if err != nil {
		return nil, err
	}

	return issues, nil
}

func logIssues(issues []Issue) {
	for _, issue := range issues {
		fmt.Println(issue.Title)
	}
}

func TestSubDomain(t *testing.T) {
	issues, err := getIssues(domain)
	if err != nil {
		log.Fatalf("error getting issues data: %v", err)
	}
	logIssues(issues)
}
