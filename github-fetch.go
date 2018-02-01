package main

import (
	"github.com/google/go-github/github"
	"gopkg.in/src-d/go-git.v4"
	"os"
	"fmt"
	"strconv"
)

func FetchRepository(repo * github.Repository, location string) (err error) {
	cloneURL := repo.GetGitURL()
	path := location + strconv.FormatInt(*repo.ID, 10)
	fmt.Println(path)
	_, err = git.PlainClone(path, false, &git.CloneOptions{
		URL: cloneURL,
		Progress: os.Stdout,
	})

	return err
}

