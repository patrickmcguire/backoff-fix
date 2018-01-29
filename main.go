package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const searchText string = "\"backoff.Retry\" language:go"
const exactSearchText string = "backoff.Retry"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	username := os.Getenv("USERNAME")
	apiKey := os.Getenv("GITHUB_TOKEN")
	results := SearchGithub(searchText, exactSearchText, username, apiKey)
	fmt.Println(results)
}

