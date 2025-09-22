package kvstorage

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
	"os"
)

var rdb *redis.Client

func InitRedis() {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "6379"
	}

	addr := host + ":" + port
	rdb = redis.NewClient(&redis.Options{
		Addr: addr,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.Fatal("Ошибка подключения к Redis:", err)
	}
	log.Println("Подключение к Redis успешно:", addr)
}


func AddURL(shortCode, originalURL string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := rdb.Set(ctx, shortCode, originalURL, time.Hour).Err()
	if err != nil {
		log.Println("Ошибка сохранения в Redis:", err)
	}
	return err
}

func GetOriginalURL(shortCode string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	originalURL, err := rdb.Get(ctx, shortCode).Result()
	if err != nil {
		log.Println("Ошибка чтения из Redis:", err)
		return "", err
	}
	return originalURL, err
}