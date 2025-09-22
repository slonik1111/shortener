package db

import (
    "database/sql"
    _ "github.com/lib/pq" // PostgreSQL драйвер
    "log"
	"os"
	"fmt"
)

var DB *sql.DB

func Connect() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	if host == "" {
		host = "localhost"
		port = "5432"
		user = "shortener_user"
		password = "secret"
		name = "shortener"
	}

	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		user, password, name, host, port
	)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Ошибка проверки подключения к базе данных:", err)
	}

	log.Printf("Успешное подключение к Postgres (host=%s port=%s db=%s)\n", host, port, name)
}


func AddURL(shortCode, originalURL string) error {
	_, err := DB.Exec("INSERT INTO urls (short_code, original_url) VALUES ($1, $2)", shortCode, originalURL)
	return err
}

func GetOriginalURL(shortCode string) (string, error) {
	var originalURL string
	err := DB.QueryRow("SELECT original_url FROM urls WHERE short_code = $1", shortCode).Scan(&originalURL)
	return originalURL, err
}