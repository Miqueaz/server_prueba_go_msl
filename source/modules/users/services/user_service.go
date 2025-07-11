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

	datas, err := s.Read(map[string]any{"ID": id})

	if len(datas) <= 0 {
		return nil, err
	}

	data := datas[0]

	// role, _ := roles.Service.ReadOne(data.Rol)

	sanitizer := user_model.UserSanitizer{
		PrimerNombre:    data.PrimerNombre,
		SegundoNombre:   data.SegundoNombre,
		PrimerApellido:  data.PrimerApellido,
		SegundoApellido: data.SegundoApellido,
		Matricula:       data.Matricula,
		Correo:          data.Correo,
		// Role:            &role.Nombre,
	}

	return &sanitizer, nil
}

var Service = base_service.NewService[UserService](*user_model.Model)
