package main

import (
	"log"

	"github.com/saniasiddiqui231/distributed-rate-limiter/internal/backend"
)

func main() {

	server := backend.NewServer()

	log.Println("Backend running on http://localhost:8081")

	log.Fatal(server.ListenAndServe())
}
