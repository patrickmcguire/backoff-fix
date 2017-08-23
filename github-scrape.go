package main

import (
	"net/http"
	"golang.org/x/net/html"
	"io"
)

func getRepositoryNames(body io.Reader) []string {
	z := html.NewTokenizer(r)
	var repoNames []string

	for tokType := z.Next() {
		if tokType == html.ErrorToken {
			return
		}

		if tokType == html.StartTagToken {
			tok := z.Token()
			isAnchor := tok.Data == "a"

			if isAnchor {
				if linkIsRepoName(tok) {
					append(repoNames, "cake")
				}
			}
		}
	}

	return repoNames
}

func linkIsRepoName(token html.Token) bool {
	isRepo := false

	for _, attr := range token.Attr {
		if attr.Key == "class" && attr.Value == "text.bold" {
			isRepo = true
			return isRepo
		}
	}

	return isRepo
}
