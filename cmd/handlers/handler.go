package handlers

import (
	"oapi-codegen-fiber/bootstrap"
	"oapi-codegen-fiber/controller"

	fiber "github.com/gofiber/fiber/v2"
)

type MyHandler struct {
	Application bootstrap.Application
}

func NewServiceInitial(app bootstrap.Application) MyHandler {
	return MyHandler{
		Application: app,
	}
}

type ServerInterfaceWrapper struct {
	CheckHandler controller.ICheckController
}

// Check implements api.ServerInterface.
func (h *ServerInterfaceWrapper) Check(c *fiber.Ctx) error {
	return h.CheckHandler.Execute(c)
}
