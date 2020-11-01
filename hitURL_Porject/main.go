package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	var results = make(map[string]string)

	urls := []string{
		"https://www.airbnb.com",
		"https://www.amazone.com",
		"https://www.google.com",
		"https://www.reddit.com",
		"https://soundcloud.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
	}

	for _, url := range urls {
		var status = "OK"
		err := hitURL(url)
		if err != errRequestFailed {
			status = "Failed"
		}

		results[url] = status
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking: ", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return errRequestFailed
	}

	if resp.StatusCode >= 400 {
		fmt.Println(resp.StatusCode)
		return errRequestFailed
	}

	return nil
}
