package main

import (
	"oapi-codegen-fiber/api"
	"oapi-codegen-fiber/bootstrap"
	"oapi-codegen-fiber/cmd/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	f := fiber.New()

	f.Static("/swagger", "cmd")

	app := bootstrap.NewInitializeBootsrap()
	f.Group("/api/v1.0")

	serve := handlers.NewServiceInitial(app)
	checkController := serve.CheckHandler()
	wrapper := &handlers.ServerInterfaceWrapper{
		CheckHandler: checkController,
	}

	api.RegisterHandlers(f, wrapper)
	f.Listen(":3000")
}
