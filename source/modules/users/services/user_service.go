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

func (s *UserService) ReadOne(idStr string) (user_model.UserSanitizer, error) {
	// Implementation for reading user data
	id, err := strconv.Atoi(idStr)
	data, err := s.Model.Find.Where("ID", "=", id).Exec(context.Background())
	return data[0], err
}

func Sanitizar[S any, E any](data E) (S, error) {
	// Inicializar un valor vacío de tipo S
	var sanitized S

	// Convertir los datos a JSON para sanitización (por ejemplo, eliminando valores peligrosos)
	sanitizedData, err := json.Marshal(data)
	if err != nil {
		return sanitized, errors.New("failed to sanitize data: " + err.Error())
	}

	// En este punto, podrías modificar o verificar los datos en sanitizedData
	// Aquí se puede agregar lógica para procesar los datos antes de deserializarlos.

	// Deserializar los datos sanitizados de vuelta al tipo `S`
	err = json.Unmarshal(sanitizedData, &sanitized)
	if err != nil {
		return sanitized, errors.New("failed to unmarshal sanitized data: " + err.Error())
	}

	// Retornar los datos sanitizados y sin errores
	return sanitized, nil
}

var Service = base_service.NewService[UserService](*user_model.Model)
