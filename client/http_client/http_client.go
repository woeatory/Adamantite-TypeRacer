package http_client

import (
	"log"
	"net/http"
	"net/http/cookiejar"
)

var HttpClient = CreateClient()

func CreateClient() *http.Client {
	return &http.Client{Jar: createCookieJar()}
}

func createCookieJar() http.CookieJar {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}
	return jar
}