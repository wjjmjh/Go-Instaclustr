package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	username      string
	api_key       string
	api_host_name string
}

func main() {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error when parsing configuration file:", err)
	}
}
