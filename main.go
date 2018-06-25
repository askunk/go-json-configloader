package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Username []string
	APIKey   []string
	BaseURL  []string
}

func main() {

	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := Config{}
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(config.Username, config.APIKey, config.BaseURL)
}
