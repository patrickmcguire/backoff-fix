package main

import (
	"log"
	"strings"
	"github.com/google/go-github/github"
	"context"
)

func fetchMatches(
	client *github.Client,
	search string,
) (result *github.CodeSearchResult) {
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
	client *github.Client,
	search string,
	exactSearch string,
) (exactResults []github.CodeResult) {

	results := fetchMatches(client, search)
	exactResults = selectExactMatches(results, exactSearch)
	return exactResults
}

