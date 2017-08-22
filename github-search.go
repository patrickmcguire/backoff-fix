package main

import (
	"net/http"
	"bytes"
	"log"
	"fmt"
)

const base = "https://api.github.com/search"
const contentType = "application/vnd.github.v3.text-match+json"

func SearchGithub(search, username, personalKey string) (items []string) {
	var url = base + "?q=" + search
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(username, personalKey)
	req.Header.Add("Accept", contentType)

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

