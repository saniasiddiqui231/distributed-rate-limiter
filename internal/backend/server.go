package backend

import (
	"fmt"
	"net/http"
)

func NewServer() *http.Server {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Request reached the backend!")
	})

	return &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}
}
