package user_service

import (
	user_model "main/modules/users/model"
	base_service "main/pkg/base/service"
)

type UserService struct {
	base_service.Service[user_model.Struct]
}

// Read: Simula la lectura de datos de un modelo de usuario
func (u *UserService) Read(filter map[string]any, config map[string]int) ([]map[string]any, error) {
	data := []string{"Alice", "Bob"}
	result := make([]map[string]any, len(data))
	for i, name := range data {
		result[i] = map[string]any{"name": name}
	}
	return result, nil
}
