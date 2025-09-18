package handlers

import (
	"encoding/json"
	"html/template"	
	"math/rand"
	"net/http"
	"github.com/slonik1111/shortener/internal/db"
	"github.com/slonik1111/shortener/internal/kvstorage"
	"log"
)

var page = template.Must(template.ParseFiles("templates/index.html"))

type ShortenRequest struct {
	URL string `json:"url"`
}

// Обработчик корневого пути и перенаправления
func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		page.Execute(w, nil)
		return
	}
	short := r.URL.Path[1:]

	if fullURL, err := kvstorage.GetOriginalURL(short); err == nil {
		log.Println("Взято из Redis:", fullURL)
		http.Redirect(w, r, fullURL, http.StatusFound)
		return
	}

	if fullURL, err := db.GetOriginalURL(short); err == nil {
		log.Println("Взято из db:", fullURL)
		http.Redirect(w, r, fullURL, http.StatusFound)
		return 
	}
	http.NotFound(w, r)
}

// Генерация случайной строки для короткого URL
func generateShortURL() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 6)	
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// Генерация короткого URL
func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	
	longURL := r.FormValue("url")
	
	if longURL == "" {
		var req ShortenRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.URL == "" {
			http.Error(w, "Некорректный URL", http.StatusBadRequest)
			return
		}
		longURL = req.URL
	}

	generatedShort := generateShortURL()
	shortURL := "http://localhost:8080/" + generatedShort

	kvstorage.AddURL(generatedShort, longURL)
	db.AddURL(generatedShort, longURL)
	contentType := r.Header.Get("Content-Type")

	if contentType == "application/json" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"short_url": shortURL})
	} else {
		page.Execute(w, shortURL)
	}
}