package http_client

import (
	"log"
	"net/http"
	"net/http/cookiejar"
)

var (
	DOMAIN     = "localhost"
	PORT       = ":8080"
	ADDRESS    = DOMAIN + PORT
	Login      = "/auth/login"
	Signup     = "/auth/signup"
	SaveResult = "/score/newScoreRecord"
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
