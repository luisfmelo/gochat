package main

import (
	"log"
	"luisfmelo/gochat/pkg"
	"net/http"
)

func main() {
	log.Println("Starting Websocket chat")
	hub := pkg.Hub{
		Clients: map[string]pkg.Client{},
	}

	// HTTP API
	http.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) { pkg.GetUsers(w, r, hub) })
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) { pkg.HealthCheck(w, r) })

	// WEB SOCKETS
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) { pkg.Websocket(w, r, hub) })

	log.Fatal(http.ListenAndServe(":8080", nil))
}
