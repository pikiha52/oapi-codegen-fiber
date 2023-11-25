package controller

import (
	"oapi-codegen-fiber/api"

	"github.com/gofiber/fiber/v2"
)

type checkController struct {
	// initialize for usecase or anythink in controller
}

type ICheckController interface {
	Execute(c *fiber.Ctx) error
}

func NewCheckController() ICheckController {
	return &checkController{}
}

// Execute implements ICheckController.
func (*checkController) Execute(c *fiber.Ctx) error {
	var apiRes api.GlobalResponses
	apiRes.ResponseCode = "200"
	apiRes.ResponseMessage = "Go is running!"
	return c.JSON(apiRes)
}
