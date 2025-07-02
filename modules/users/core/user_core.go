package user_core

import (
	user_model "main/modules/users/model"
	base_controller "main/pkg/base/controller"
	base_models "main/pkg/base/models"
	base_service "main/pkg/base/service"
)

var Model = base_models.NewModel[user_model.Struct]("user", "users")

//	var Service = base_service.NewServices(*model, &user_service.UserService{
//		Service: base_service.Service[user_model.Struct]{
//			Model: *model,
//		},
//	})
var Service = base_service.NewServices(*Model, nil)
var Controller = base_controller.NewController(*Service, nil)
