package network

import (
	"net/http"
)

func Call(url string) (*http.Response, error) {
	return http.Get(url)
}
