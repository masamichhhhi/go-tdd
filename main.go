package main

import (
	"log"
	"net/http"

	"github.com/masamichhhhi/go-tdd/server"
)

func main() {
	server := &server.PlayerServer{server.NewInMemoryPlayerStore()}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
