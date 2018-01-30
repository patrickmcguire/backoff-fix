package main

import (
	"log"
	"github.com/google/go-github/github"
	"context"
	"golang.org/x/oauth2"
	"strings"
)

func fetchMatches(
	search string,
	username string,
	personalKey string,
) (result *github.CodeSearchResult) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: personalKey},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	opts := &github.SearchOptions{TextMatch: true}

	results, _, err := client.Search.Code(context.Background(), search, opts)


	if err != nil {
		log.Fatal(err)
	}

	return results
}

func selectExactMatches(result *github.CodeSearchResult, exactSearch string) []github.CodeResult {
	// this doesn't currently handle pagination, but the current result set is one page
	var exactResults []github.CodeResult
	codeResults := result.CodeResults
	for _, result := range(codeResults) {
		matches := result.TextMatches
		shouldAdd := false
		for _, match := range(matches) {
			if (strings.Contains(*match.Fragment, exactSearch)) {
				shouldAdd = true
			}
		}

		if (shouldAdd) {
			exactResults = append(exactResults, result)
		}
	}

	return exactResults
}

func SearchGithub(
	search string,
	exactSearch string,
	username string,
	personalKey string,
) (exactResults []github.CodeResult) {
	results := fetchMatches(search, username, personalKey)
	exactResults = selectExactMatches(results, exactSearch)
	return exactResults
}

