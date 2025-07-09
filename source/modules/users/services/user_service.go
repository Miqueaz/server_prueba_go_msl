package user_service

import (
	base_service "main/pkg/base/service"
	user_model "main/source/modules/users/models"
)

type UserService struct {
	base_service.Service[user_model.UserStruct]
}

func (s *UserService) ReadOne(id int) (*user_model.UserSanitizer, error) {
	// Implementation for reading user data

	data, _ := s.Read(map[string]any{"ID": id})

	return base_service.Sanitizar[*user_model.UserSanitizer](data[0])
}

var Service = base_service.NewService[UserService](*user_model.Model)
