package main

import (
	"github.com/google/go-github/github"
	"gopkg.in/src-d/go-git.v4"
	"os"
	"fmt"
	"strconv"
)

func FetchRepository(repo * github.Repository, location string) (repoDir string, err error) {
	cloneURL := repo.GetGitURL()
	repoDir = location + strconv.FormatInt(*repo.ID, 10)
	fmt.Println(repoDir)
	_, err = git.PlainClone(repoDir, false, &git.CloneOptions{
		URL: cloneURL,
		Progress: os.Stdout,
	})

	return repoDir, err
}

