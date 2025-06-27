package user_service

import (
	user_model "main/modules/users/model"
	base_controller "main/pkg/base/controller"
)

type UserController struct {
	base_controller.Controller[user_model.Struct]
}

func Init() {
	var UserController = UserController{}
	base_controller.NewController(UserController.Model, UserController.Methods)
}
