package bootstrap

import (
	"oapi-codegen-fiber/internal/env"

	"github.com/gofiber/fiber/v2/log"
	"github.com/redis/go-redis/v9"
)

type Application struct {
	Env    *env.ENV
	Logger *log.Logger
	Redis  *redis.Client
}

func NewInitializeBootsrap() Application {
	app := Application{}
	app.Env = env.NewEnv()
	app.Redis = RedisClient(app.Env.RedisHost, app.Env.RedisPort, app.Env.RedisUsername, app.Env.RedisPassword)
	return app
}
