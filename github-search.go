package main

import (
	"net/http"
	"net/url"
	"bytes"
	"log"
	"fmt"
)

const base = "https://api.github.com/search"
const contentType = "application/vnd.github.v3.text-match+json"

func SearchGithub(
	search string,
	username string,
	personalKey string,
	queryParams map[string]string,
) (items []string) {

	values := url.Values{}
	values.Set("q", search)
	fmt.Println(values)

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

