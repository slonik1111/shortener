package kvstorage

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var rdb *redis.Client

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:    "localhost:6379",
	})

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Ошибка подключения к Redis:", err)
	}
	log.Println("Подключение к Redis успешно")
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