package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server ...")
	cfg, err := NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	client := NewClient(cfg)
	handler := NewHandler(cfg, client)
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/health", handler.Health)
	http.HandleFunc("/auth", handler.Auth)
	http.HandleFunc("/auth/callback", handler.Callback)

	log.Println("listening on port 9000")
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatal(err)
	}
}
