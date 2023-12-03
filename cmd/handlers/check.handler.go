package handlers

import (
	"oapi-codegen-fiber/controller"
	"oapi-codegen-fiber/repositories"
	"oapi-codegen-fiber/usecase"
)

func (h *MyHandler) CheckHandler() controller.ICheckController {
	redisRepo := repositories.NewRedisRepositories(h.Application.Redis)
	checkUsecase := usecase.NewCheckUsecase(redisRepo)
	checkController := controller.NewCheckController(checkUsecase)
	return checkController
}
