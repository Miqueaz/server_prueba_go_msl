package user_service

import (
	user_model "main/modules/users/model"
	base_service "main/pkg/base/service"
)

type UserService struct {
	base_service.Service[user_model.Struct]
}

// Read: Simula la lectura de datos de un modelo de usuario
func (u *UserService) Read(filter map[string]any, config map[string]int) ([]user_model.Struct, error) {
	data := []string{"Alice", "Bob"}
	result := make([]user_model.Struct, len(data))
	for i, name := range data {
		result[i] = user_model.Struct{Username: name}
	}
	return result, nil
}
