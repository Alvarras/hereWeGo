package client

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
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

func generateKey() string {
	const characters = "ABCDEF0123456789"
	result := ""
	rand.New(rand.NewSource(0))
	for i := 0; i < 16; i++ {
		result += string(characters[rand.Intn(len(characters))])
	}
	return result
}

func getUsers(url, apiKey string) ([]User, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-API-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var users []User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func logUsers(users []User) {
	for _, user := range users {
		fmt.Printf("User Name: %s, Role: %s, Experience: %d, Remote: %v\n",
			user.User.Name, user.Role, user.Experience, user.Remote)
	}
}

func deleteUser(baseURL, apiKey, id string) error {
	fullURL := baseURL + "/" + id
	req, err := http.NewRequest("DELETE", fullURL, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("X-API-Key", apiKey)

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("error deleting user: %s", res.Status)
	}
	return nil
}

func TestDeleteUser(t *testing.T) {
	userId := "5079832d-a0a3-4e4f-9cd6-fdfc633c0fa6"
	url := "https://api.boot.dev/v1/courses_rest_api/learn-http/users"
	apiKey := generateKey()

	users, err := getUsers(url, apiKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Logging user records:")
	logUsers(users)
	fmt.Println("---")

	err = deleteUser(url, userId, apiKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted user with id: %s\n", userId)
	fmt.Println("---")

	newUsers, err := getUsers(url, apiKey)
	if err != nil {
		log.Fatal(err)
	}
	logUsers(newUsers)
	fmt.Println("---")
}
