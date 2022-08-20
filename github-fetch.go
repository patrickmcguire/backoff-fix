package main

import (
	"fmt"
	"github.com/google/go-github/v46/github"
	"gopkg.in/src-d/go-git.v4"
	"os"
	"path/filepath"
	"strconv"
)

func FetchRepository(repo *github.Repository, location string) (repoDir string, err error) {
	cloneURL := repo.GetSSHURL()
	fmt.Println(cloneURL)
	repoDir = filepath.Join(location, strconv.FormatInt(*repo.ID, 10))
	if _, err := os.Stat(repoDir); os.IsNotExist(err) {
		err := os.Mkdir(repoDir, 0700)
		if err != nil {
			fmt.Println(err)
		}
		// TODO: handle error
	}

	fmt.Println(repoDir)
	_, err = git.PlainClone(repoDir, false, &git.CloneOptions{
		URL:      cloneURL,
		Progress: os.Stdout,
	})

	return repoDir, err
}
