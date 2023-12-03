package controller

import (
	"oapi-codegen-fiber/api"
	"oapi-codegen-fiber/usecase"

	"github.com/gofiber/fiber/v2"
)

type checkController struct {
	CheckUsecase usecase.ICheckUsecase
}

type ICheckController interface {
	Execute(c *fiber.Ctx) error
}

func NewCheckController(checkUsecase usecase.ICheckUsecase) ICheckController {
	return &checkController{
		CheckUsecase: checkUsecase,
	}
}

// Execute implements ICheckController.
func (co *checkController) Execute(c *fiber.Ctx) error {
	var apiRes api.GlobalResponses

	err := co.CheckUsecase.Execute(c.Context())
	if err != nil {
		apiRes.ResponseCode = "500"
		apiRes.ResponseMessage = "Internal server error!"
		return c.Status(500).JSON(apiRes)
	}

	apiRes.ResponseCode = "200"
	apiRes.ResponseMessage = "Go is running!"
	return c.JSON(apiRes)
}
