package bootstrap

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func RedisClient(redisHost string, redisPort string, redisUsername string, redisPassword string) *redis.Client {
	redisUrl := redisHost + ":" + redisPort
	host := redisUrl
	password := redisPassword
	client := redis.NewClient(
		&redis.Options{
			Addr:     host,
			Username: redisUsername,
			Password: password,
			DB:       0,
		},
	)

	if err := client.Ping(context.TODO()).Err(); err != nil {
		panic("Failed to start redis client")
	}
	Redis = client

	fmt.Println("Connected to redis client with port: " + redisPort)
	return Redis
}
