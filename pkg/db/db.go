package db

import (
	"github.com/redis/go-redis/v9"
)

type Handler struct {
	DB *redis.Client
}

func Init(url string) Handler {
	db := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "",
		DB:       0,
	})
	return Handler{db}
}
