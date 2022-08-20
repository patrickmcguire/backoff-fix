package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v46/github"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"log"
	"os"
)

const searchText string = "\"backoff.Retry\" language:go"
const exactSearchText string = "backoff.Retry"

func main() {
	outPath := os.Args[1]

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

	for _, result := range results {
		repo := result.Repository
		ctx = context.Background()
		repo, _, err = client.Repositories.GetByID(ctx, *repo.ID)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("got by id", *repo.ID)
		}

		repoDir, err := FetchRepository(repo, outPath)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("got repo")
		}

		err = ExamineRepository(repoDir)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("examined")
		}
	}
}
