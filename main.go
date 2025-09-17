package main

import (
	"log"
	"net/http"
 	_ "github.com/lib/pq"
	"github.com/slonik1111/shortener/internal/handlers"
	"github.com/slonik1111/shortener/internal/db"
)

func main() {
	db.Connect()
	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/shorten", handlers.ShortenHandler)

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
