package main

import (
	"log"
	"net/http"

	"github.com/slonik1111/shortener/internal/handlers"
)

func main() {
	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/shorten", handlers.ShortenHandler)

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
