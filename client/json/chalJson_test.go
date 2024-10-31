package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"testing"
)

func getResources(url string) ([]map[string]any, error) {
	var resources []map[string]any

	res, err := http.Get(url)
	if err != nil {
		return resources, err
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&resources); err != nil {
		return nil, err
	}
	return resources, nil

}

func logResources(resources []map[string]any) {
	var formattedStrings []string

	for _, resource := range resources {
		for key, value := range resource {
			formattedStrings = append(formattedStrings, fmt.Sprintf("Key: %s - Value: %v", key, value))
		}
	}

	sort.Strings(formattedStrings)

	for _, str := range formattedStrings {
		fmt.Println(str)
	}
}

const baseUrl = "https://api.boot.dev"

func TestChallegeJson(t *testing.T) {
	issues, err := getResources(baseUrl + "/v1/courses_rest_api/learn-http/issues?limit=1")
	if err != nil {
		fmt.Println("Error getting issues:", err)
		return
	}
	fmt.Println("Issue:")
	logResources(issues)
	fmt.Println("---")

	projects, err := getResources(baseUrl + "/v1/courses_rest_api/learn-http/projects?limit=1")
	if err != nil {
		fmt.Println("Error getting projects:", err)
		return
	}
	fmt.Println("Project:")
	logResources(projects)
	fmt.Println("---")

	users, err := getResources(baseUrl + "/v1/courses_rest_api/learn-http/users?limit=1")
	if err != nil {
		fmt.Println("Error getting users:", err)
		return
	}
	fmt.Println("User:")
	logResources(users)
}
