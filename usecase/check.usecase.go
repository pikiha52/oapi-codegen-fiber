package usecase

import (
	"context"
	"fmt"
	"oapi-codegen-fiber/repositories"
	"time"
)

type checkUsecase struct {
	RedisRepositories repositories.IRedisRepositories
}

type ICheckUsecase interface {
	Execute(ctx context.Context) error
}

func NewCheckUsecase(redisRepo repositories.IRedisRepositories) ICheckUsecase {
	return &checkUsecase{
		RedisRepositories: redisRepo,
	}
}

// Execute implements ICheckUsecase.
func (c *checkUsecase) Execute(ctx context.Context) error {
	var data string
	key := "checkKey"
	get, err := c.RedisRepositories.Get(key, ctx)
	if err != nil {
		if err.Error() == "key does not exist" {
			fmt.Println("key does not exist: " + key)
			expireAt := time.Now().Add(time.Minute)
			err := c.RedisRepositories.Hset(key, "first cache", expireAt, ctx)
			if err != nil {
				return err
			}
			fmt.Println("redis cache success set key: " + key)
			get, _ := c.RedisRepositories.Get(key, ctx)
			data = get
		} else {
			return err
		}
	} else {
		fmt.Println("redis get success key: " + key)
		data = get
	}

	fmt.Println(data)
	err = c.RedisRepositories.Del(key, ctx)
	if err != nil {
		return err
	}

	fmt.Println("redis del success key: " + key)
	return nil
}
