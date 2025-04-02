package helpers

import (
	"context"
	"strings"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.ClusterClient

func SetupRedis() {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: strings.Split(GetEnv("REDIS_HOST", "localhost:6379"), ","),
	})

	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		Logger.Error("Failed to connect to Redis: ", err)
		return
	}
	Logger.Info("PING REDIS: " + ping)

	RedisClient = client
}
