package user_service

import (
	"context"
	base_service "main/pkg/base/service"
	user_model "main/source/modules/users/models"
	"strconv"
)

type UserService struct {
	base_service.Service[user_model.UserStruct]
}

func (s *UserService) ReadOne(idStr string) (user_model.UserStruct, error) {
	// Implementation for reading user data
	id, err := strconv.Atoi(idStr)
	data, err := s.Model.Find.Where("ID", "=", id).Exec(context.Background())
	return data[0], err
}

var Service = base_service.NewService[UserService](*user_model.Model)
