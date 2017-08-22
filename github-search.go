package main

import (
	"net/http"
	"bytes"
	"log"
)

var base = "https://github.com/search"

func SearchGithub(search, username, personalKey string) (items []string) {
	req, err := http.NewRequest("GET", "https://gitub.github.com", nil)
	req.SetBasicAuth(username, personalKey)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	s := buf.String() // Does a complete copy of the bytes in the buffer.
	return []string{s}
}

