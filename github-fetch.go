package main

import (
	"github.com/google/go-github/github"
	"gopkg.in/src-d/go-git.v4"
	"os"
	"encoding/json"
)

func FetchRepository(repo * github.Repository, location string) (err error) {
	cloneURL := repo.GetGitURL()
	url := *repo.URL
	path := location + url
	json.NewEncoder(os.Stdout).Encode(repo)
	_, err = git.PlainClone(path, false, &git.CloneOptions{
		URL: cloneURL,
		Progress: os.Stdout,
	})

	return err
}

