package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"math"
	"strings"
	"time"
)

func fetchMatches(
	client *github.Client,
	search string,
) (results []github.CodeResult, err error) {

	opts := &github.SearchOptions{TextMatch: true, ListOptions: github.ListOptions{Page: 0, PerPage: 100}}
	for {
		result, response, firstErr := client.Search.Code(context.Background(), search, opts)
		if firstErr != nil {
			retryCount := 1
			for retryCount < 15 {
				backoff(retryCount)
				result, response, err = client.Search.Code(context.Background(), search, opts)
				if err == nil {
					break
				}
			}

			return results, firstErr
		}

		for _, codeResult := range result.CodeResults {
			results = append(results, codeResult)
		}

		if response.NextPage == 0 {
			break
		}

		fmt.Println("Page", opts.ListOptions.Page, "completed")
		opts.ListOptions.Page = response.NextPage
		time.Sleep(5000 * time.Millisecond)
	}

	return results, nil
}

func selectExactMatches(results []github.CodeResult, exactSearch string) []github.CodeResult {
	// this doesn't currently handle pagination, but the current result set is one page
	var exactResults []github.CodeResult
	for _, result := range results {
		matches := result.TextMatches
		shouldAdd := false
		for _, match := range matches {
			if strings.Contains(*match.Fragment, exactSearch) {
				shouldAdd = true
			}
		}

		if shouldAdd {
			exactResults = append(exactResults, result)
		}
	}

	return exactResults
}

func backoff(attemptNumber int) {
	sleepTime := math.Pow(1.5, (float64(attemptNumber))) * 5000
	rounded := math.Round(sleepTime)
	time.Sleep(time.Duration(rounded))
}

func SearchGithub(
	client *github.Client,
	search string,
	exactSearch string,
) (exactResults []github.CodeResult, err error) {

	results, err := fetchMatches(client, search)
	if err != nil {
		return nil, err
	}
	exactResults = selectExactMatches(results, exactSearch)
	return exactResults, nil
}
