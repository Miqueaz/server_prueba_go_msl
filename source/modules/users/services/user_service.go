package user_service

import (
	"context"
	base_service "main/pkg/base/service"
	user_model "main/source/modules/users/models"
)

type UserService struct {
	base_service.Service[user_model.UserStruct]
}

func (s *UserService) Read() {
	// Implementation for reading user data
	ctx := context.Background()
	s.Model.Find.Exec(ctx)

}

var Service = base_service.NewService[UserService](*user_model.Model)
