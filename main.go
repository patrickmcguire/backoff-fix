package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"github.com/google/go-github/github"
	"context"
	"golang.org/x/oauth2"
)

const searchText string = "\"backoff.Retry\" language:go"
const exactSearchText string = "backoff.Retry"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("GITHUB_TOKEN")
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: apiKey},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	results, err := SearchGithub(client, searchText, exactSearchText)
	if err != nil {
		log.Fatal(err)
	}

	for _, result := range(results) {
		repo := result.Repository
		ctx = context.Background()
		repo, _, err = client.Repositories.GetByID(ctx, *repo.ID)
		if err != nil {
			fmt.Println(err)
		}

		err := FetchRepository(repo, "/tmp/gituhbrepos/")
		if err != nil {
			fmt.Println(err)
		}
	}
}

