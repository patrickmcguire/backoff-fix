package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"http"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//apiKey := os.Getenv("GITHUB_TOKEN")
	//fmt.Println(apiKey)
}

