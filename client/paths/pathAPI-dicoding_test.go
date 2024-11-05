package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func getResources(path string) map[string]any {
	fullURL := "https://coffee-api.dicoding.dev" + path

	res, err := http.Get(fullURL)
	if err != nil {
		fmt.Println("Error creating Request", err)
		return nil
	}
	defer res.Body.Close()

	var resources map[string]any
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&resources)
	if err != nil {
		fmt.Println("Error decoding resources", err)
		return nil
	}
	return resources
}

func logResources(resources map[string]any) {
	for _, resource := range resources {
		jsonResource, err := json.Marshal(resource)
		if err != nil {
			fmt.Println("Error marshalling resource", err)
			continue
		}
		fmt.Printf(" -%s\n", jsonResource)
	}
}

func TestUrlPathDicoding(t *testing.T) {
	coffees := getResources("/coffees")
	fmt.Println("Coffees:")
	logResources(coffees)
	fmt.Println(" --- ")
}
