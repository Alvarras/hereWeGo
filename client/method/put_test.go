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

func logUser(user User) {
	fmt.Printf("User Name: %s, ID : %v, Role: %s, Experience: %d, Remote: %v, Location: %s, Age: %d\n",
		user.User.Name, user.ID, user.Role, user.Experience, user.Remote, user.User.Location, user.User.Age)
}

func updateUser(baseURL, id, apiKey string, data User) (User, error) {
	fullURL := baseURL + "/" + id
	encodedData, err := json.Marshal(data)
	if err != nil {
		return User{}, fmt.Errorf("error encoding data: %w", err)
	}
	req, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(encodedData))
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

func getUserById(baseURL, id, apiKey string) (User, error) {
	fullURL := baseURL + "/" + id
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return User{}, fmt.Errorf("error creating request: %w", err)
	}
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

func TestPutUser(t *testing.T) {
	userId := "2f8282cb-e2f9-496f-b144-c0aa4ced56db"
	baseURL := "https://api.boot.dev/v1/courses_rest_api/learn-http/users"
	apiKey := generateKey()

	userData, err := getUserById(baseURL, userId, apiKey)
	if err != nil {
		fmt.Println(err)
	}
	logUser(userData)

	fmt.Printf("Updating user with id: %s\n", userData.ID)
	userData.Role = "Senior Backend Developer"
	userData.Experience = 7
	userData.Remote = true
	userData.User.Name = "Allan"

	updatedUser, err := updateUser(baseURL, userId, apiKey, userData)
	if err != nil {
		fmt.Println(err)
		return
	}
	logUser(updatedUser)
}
