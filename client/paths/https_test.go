package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

type User struct {
	Role       string `json: "role"`
	ID         string `json: "id"`
	Experience int    `json: "experience"`
	Remote     bool   `json: "remote"`
	User       struct {
		Name     string `json: "name"`
		Location string `json: "location"`
		Age      int    `json: "age"`
	} `json: "user"`
}

const URL = "https://api.boot.dev/v1/courses_rest_api/learn-http/users"

func errIfNotHTTPS(URL string) error {
	url, err := url.Parse(URL)
	if err != nil {
		return err
	}
	if url.Scheme != "https" {
		return errors.New("URL must be HTTPS")
	}
	return nil
}

func getUserById(baseURL, id string) (User, error) {
	fullURL := baseURL + "/" + id
	if err := errIfNotHTTPS(fullURL); err != nil {
		return User{}, err
	}
	res, err := http.Get(fullURL)
	if err != nil {
		return User{}, err
	}
	var user User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
func logUser(user User) {
	fmt.Printf("User Name: %s, Role : %s, Experience: %d, Remote: %v \n",
		user.User.Name, user.Role, user.Experience, user.Remote)
}

func TestHTTPS(t *testing.T) {
	uuid := "2f8282cb-e2f9-496f-b144-c0aa4ced56db"
	user, err := getUserById(URL, uuid)
	if err != nil {
		fmt.Println(err)
		return
	}
	logUser(user)
}
