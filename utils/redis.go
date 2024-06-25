package utils

import (
	"context"
	"go-web-app/config"
	"log"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client
var ctx = context.Background()

func ConnectRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     config.Config.GetString("REDIS_HOST") + ":" + config.Config.GetString("REDIS_PORT"),
		Password: config.Config.GetString("REDIS_PASSWORD"),
		DB:       0,
	})

	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
}
