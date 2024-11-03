package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
)

type User struct {
	Role       string `json:"role"`
	ID         string `json:"id"`
	Experience int    `json:"experience"`
	Remote     bool   `json:"remote"`
	User       struct {
		Name     string `json:"name"`
		Location string `json:"location"`
		Age      int    `json:"age"`
	} `json:"user"`
}

func logUsers(users []User) {
	for _, user := range users {
		fmt.Printf("User Name: %s, Role: %s, Experience: %d, Remote: %v\n", user.User.Name, user.Role, user.Experience, user.Remote)
	}
}

func getUsers(url string) ([]User, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	var users []User
	err = json.NewDecoder(res.Body).Decode(&users)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}
	return users, nil
}

func TestGetUsers(t *testing.T) {
	url := "https://api.boot.dev/v1/courses_rest_api/learn-http/users"
	users, err := getUsers(url)
	if err != nil {
		log.Fatal("Error getting users:", err)
	}
	logUsers(users)
}
