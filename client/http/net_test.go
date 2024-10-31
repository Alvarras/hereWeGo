package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"
)

const issueURL = "https://api.boot.dev/v1/courses_rest_api/learn-http/issues"

func getIssueData(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error cretaing request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error cretaing request: %w", err)
	}
	return data, nil
}

func prettify(data string) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, []byte(data), "", " ")
	if err != nil {
		return "", fmt.Errorf("error indenting JSOn: %w", err)
	}

	return prettyJSON.String(), nil
}

func TestNet(t *testing.T) {
	issues, err := getIssueData(issueURL)
	if err != nil {
		log.Fatalf("error getting issue data: %v", err)
	}

	prettyData, err := prettify(string(issues))
	if err != nil {
		log.Fatalf("error getting issue data: %v", err)
	}
	fmt.Println(prettyData)
}
