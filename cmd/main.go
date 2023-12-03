package main

import (
	"fmt"
	"oapi-codegen-fiber/api"
	"oapi-codegen-fiber/bootstrap"
	"oapi-codegen-fiber/cmd/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	f := fiber.New()

	f.Static("/swagger", "cmd")

	app := bootstrap.NewInitializeBootsrap()
	f.Group(fmt.Sprintf("/api/%s", app.Env.AppVersion))

	serve := handlers.NewServiceInitial(app)
	checkController := serve.CheckHandler()
	wrapper := &handlers.ServerInterfaceWrapper{
		CheckHandler: checkController,
	}

	api.RegisterHandlers(f, wrapper)
	f.Listen(app.Env.AppPort)
}
