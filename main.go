package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	argsWithoutProg := os.Args[1:]

	apiKey := os.Getenv("GITHUB_TOKEN")
	opts := make(map[string]string)

	fmt.Println(SearchGithub(argsWithoutProg[0], "patrickmcguire", apiKey, opts))
}

