package client

import (
	"bytes"
	"encoding/json"
	"fmt"
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
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("X-API-KEY", apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	var users []User
	err = json.NewDecoder(res.Body).Decode(&users)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}
	return users, nil

}

func logUsers(users []User) {
	for _, user := range users {
		fmt.Printf("User Name: %s, Role: %s, Experience: %d, Remote: %v\n", user.User.Name, user.Role, user.Experience, user.Remote)
	}
}

func createUser(url, apiKey string, data User) (User, error) {
	encodedData, err := json.Marshal(data)
	if err != nil {
		return User{}, fmt.Errorf("error encoding data: %w", err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(encodedData))
	if err != nil {
		return User{}, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", apiKey)

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return User{}, fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	var user User
	decoderData := json.NewDecoder(res.Body)
	err = decoderData.Decode(&user)
	if err != nil {
		return User{}, fmt.Errorf("error decoding response: %w", err)
	}
	return user, nil

}

func TestPostUsers(t *testing.T) {
	userToCreate := User{
		Role:       "Junior Developer",
		Experience: 2,
		Remote:     true,
		User: struct {
			Name     string `json:"name"`
			Location string `json:"location"`
			Age      int    `json:"age"`
		}{
			Name:     "Dan",
			Location: "NOR",
			Age:      29,
		},
	}

	url := "https://api.boot.dev/v1/courses_rest_api/learn-http/users"
	apiKey := generateKey()

	fmt.Println("Retrieving user data...")
	userDataFirst, err := getUsers(url, apiKey)
	if err != nil {
		fmt.Println("Error retrieving users:", err)
		return
	}
	logUsers(userDataFirst)
	fmt.Println("---")

	fmt.Println("Creating new character...")
	creationResponse, err := createUser(url, apiKey, userToCreate)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	jsonData, _ := json.Marshal(creationResponse)
	fmt.Printf("Creation response body: %s\n", string(jsonData))
	fmt.Println("---")

	fmt.Println("Retrieving user data...")
	userDataSecond, err := getUsers(url, apiKey)
	if err != nil {
		fmt.Println("Error retrieving users:", err)
		return
	}
	logUsers(userDataSecond)
	fmt.Println("---")
}
