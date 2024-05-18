package api

import (
	"net/http"
)

const URL = "https://api.akl.gg/"

var Client = &http.Client{}

func Get(path string) (*http.Response, error) {
	return Client.Get(URL + path)
}
