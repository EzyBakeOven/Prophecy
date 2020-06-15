package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Realm string
}

func processconfig() {
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("Error failed to process config file:", err)
	}
	fmt.Println("Welcome to", configuration.Realm)
}

func main() {
	processconfig()
}
