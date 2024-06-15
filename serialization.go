package main

import (
	"encoding/json"
	"fmt"
)

type ApiResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Password string `json:"-"`
}

func main() {
	sampleResponse := &ApiResponse{
		ID:       1,
		FullName: "Billy Joe",
		Username: "joe",
		Password: "supersecretvalue",
	}
	stringBytes, err := json.Marshal(sampleResponse)
	if err != nil {
		fmt.Println("encountered error!", err)
	}
	fmt.Println(string(stringBytes))

	inputText := []byte("{\"id\": 2, \"full_name\": \"Jane Doe\", \"username\": \"jane\"}")
	var inputInstance ApiResponse
	if err := json.Unmarshal(inputText, &inputInstance); err != nil {
		fmt.Println("Failed to deserialize input text!", err)
	}
	fmt.Println("input instance:", inputInstance)
}
