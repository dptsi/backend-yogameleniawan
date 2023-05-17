package go_web

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
