package user_service

import (
	base_service "main/pkg/base/service"
	user_model "main/source/modules/users/models"
)

type UserService struct {
	base_service.Service[user_model.UserStruct]
}

var Service = base_service.NewService[UserService](*user_model.Model)
