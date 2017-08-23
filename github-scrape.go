package main

import (
	"golang.org/x/net/html"
	"io"
)

func getRepositoryNames(body io.Reader) []string {
	z := html.NewTokenizer(body)
	var repoNames []string

	for tokType := z.Next(); tokType != html.ErrorToken; {
		if tokType == html.StartTagToken {
			tok := z.Token()
			isAnchor := tok.Data == "a"

			if isAnchor {
				if linkIsRepoName(tok) {
					repoNames = append(repoNames, "cake")
				}
			}
		}
	}

	return repoNames
}

func linkIsRepoName(token html.Token) bool {
	isRepo := false

	for _, attr := range token.Attr {
		if attr.Key == "class" && attr.Val == "text.bold" {
			isRepo = true
			return isRepo
		}
	}

	return isRepo
}
