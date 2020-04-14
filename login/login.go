package main

import (
	"log"
	"net/http"
	"net/url"
)

func main() {
	var client *http.Client
	loginURL := "https://www.flightradar24.com/premium/sig_in"
	data := url.Values{
		"email":    {"abdul@antabur.com"},
		"password": {"Robonson10"},
	}

	response, err := client.PostForm(loginURL, data)

	if err != nil {
		log.Fatalln(err.Error())
	}

	defer response.Body.Close()
}
