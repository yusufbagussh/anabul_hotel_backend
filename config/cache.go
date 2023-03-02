package config

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"os"
)

var ctx = context.Background()

func SetupRedisConnection() *redis.Client {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load to env file")
	}

	redisHost := os.Getenv("REDIS_HOST")

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:6379", redisHost),
		Password: "",
		DB:       0,
	})

	return client
}
