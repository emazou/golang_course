package main

import (
	"log"
	"net/http"
	"rpsweb/handlers"
)

func main() {
	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))
	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.PlayGame)
	router.HandleFunc("/about", handlers.About)

	port := ":8080"

	log.Printf("Server is running on port %s ", port)

	log.Fatal(http.ListenAndServe(port, router))
}