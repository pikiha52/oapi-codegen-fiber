package repositories

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisRepositories struct {
	Client *redis.Client
}

type IRedisRepositories interface {
	Set(key string, data []byte, expiredTime time.Duration, ctx context.Context) error
	Hset(key string, data string, expireAt time.Time, ctx context.Context) error
	Get(key string, ctx context.Context) (string, error)
	Del(key string, ctx context.Context) error
}

func NewRedisRepositories(r *redis.Client) IRedisRepositories {
	return &redisRepositories{
		Client: r,
	}
}

// Set implements IRedisRepositories.
func (r *redisRepositories) Set(key string, data []byte, expiredTime time.Duration, ctx context.Context) error {
	err := r.Client.Set(ctx, key, data, expiredTime).Err()
	if err != nil {
		panic(err)
	}
	return nil
}

// Hset implements IRedisRepositories.
func (r *redisRepositories) Hset(key string, data string, expireAt time.Time, ctx context.Context) error {
	err := r.Client.Do(ctx, "SET", key, data).Err()

	if err != nil {
		panic(err)
	}

	err = r.Client.ExpireAt(ctx, key, expireAt).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Map set with expiration:", key)
	return nil
}

// Get implements IRedisRepositories.
func (r *redisRepositories) Get(key string, ctx context.Context) (string, error) {
	result, err := r.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", errors.New("key does not exist")
	} else if err != nil {
		return "", redis.Nil
	}

	resultTtl := r.Client.TTL(ctx, key)
	if resultTtl.Err() != nil {
		return "", redis.Nil
	}

	ttl := resultTtl.Val()

	if ttl == time.Duration(-1) {
		return "", errors.New("key does not exist or has no expiration set")
	} else if ttl == time.Duration(-2) {
		return "", errors.New("key exists, but no expiration is set")
	} else {
		fmt.Println("Time to live:", ttl)
	}

	fmt.Println("Value:", result)
	return result, nil
}

// Del implements IRedisRepositories.
func (r *redisRepositories) Del(key string, ctx context.Context) error {
	_, err := r.Client.Del(ctx, key).Result()
	if err != nil {
		return err
	}

	return nil
}
