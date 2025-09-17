package db

import (
    "database/sql"
    _ "github.com/lib/pq" // PostgreSQL драйвер
    "log"
)

var db *sql.DB

func Connect() {
	var err error
	db, err = sql.Open("postgres", "host=localhost port=5432 user=shortener_user password=secret dbname=shortener sslmode=disable")

	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Ошибка проверки подключения к базе данных:", err)
	}
}

func AddURL(shortCode, originalURL string) error {
	_, err := db.Exec("INSERT INTO urls (short_code, original_url) VALUES ($1, $2)", shortCode, originalURL)
	return err
}

func GetOriginalURL(shortCode string) (string, error) {
	var originalURL string
	err := db.QueryRow("SELECT original_url FROM urls WHERE short_code = $1", shortCode).Scan(&originalURL)
	return originalURL, err
}