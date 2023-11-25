package handlers

import "oapi-codegen-fiber/controller"

func (h *MyHandler) CheckHandler() controller.ICheckController {
	checkController := controller.NewCheckController()
	return checkController
}
